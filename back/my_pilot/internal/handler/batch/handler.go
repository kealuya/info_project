package batch

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/bytedance/sonic"
	"github.com/streadway/amqp"
	"golang.org/x/time/rate"
	"my_pilot/common"
	"my_pilot/internal/repository"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
程序建议就是
1，初始化：connection，开channel（channel就是用来作业的基础单位），在mq声明交换机(exchange)，声明路由(routing_key)，然后声明队列(queue)，绑定交换机路由和队列。
2，生产者启动：使用channel，进行消息发送，指定交换机和路由（自然也就存储到指定队列中）。
3，消费者：使用channel，直接消费对应的队列就可以。（一个goroutine 一个channel）
如果初始化已经声明了交换机、路由，并且绑定了路由、交换机，队列，那么消费者其实只需要指定对应名称的队列，就可以消费。
** 每次batch处理，如果中间有错误停掉了程序，记得清空队列，不然计数会发生错误 **

docker run -itd --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management


*/

const (
	EXCHANGE    = "hotel_exchange"
	ROUTING_KEY = "hotel_key"
	QUEUE       = "hotel_consume_queue"

	DLX_EXCHANGE    = "hotel_dlx_exchange"
	DLX_QUEUE       = "hotel_dlx_queue"
	DLX_ROUTING_KEY = "dlx_key"
)

// 初始化RabbitMQ连接
func connectionInit() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}
	// 声明交换机
	err = ch.ExchangeDeclare(
		EXCHANGE,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %v", err)
	}
	// 声明队列
	// 一旦reject，会自动发送死信队列
	args := amqp.Table{
		"x-dead-letter-exchange":    DLX_EXCHANGE,
		"x-dead-letter-routing-key": DLX_ROUTING_KEY,
	}
	q, err := ch.QueueDeclare(
		QUEUE, //队列名
		false, //是否持久化
		false, //是否自动删除
		false, //是否排他
		false, //是否阻塞
		args,  //额外属性
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err)
	}

	//  绑定队列到交换机（指定交换机，绑定队列和路由，消费者只需要绑定队列就可以消费了）
	err = ch.QueueBind(
		q.Name,      // 队列名
		ROUTING_KEY, // 路由键
		EXCHANGE,    // 交换机名
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to register DLQ consumer: %v", err)
	}

	//==============================================================
	// 声明死信交换机
	err = ch.ExchangeDeclare(
		DLX_EXCHANGE,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare dead letter exchange: %v", err)
	}

	// 声明死信队列
	dlq, err := ch.QueueDeclare(
		DLX_QUEUE,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare dead letter queue: %v", err)
	}

	// 绑定死信队列到死信交换机
	err = ch.QueueBind(
		dlq.Name,
		DLX_ROUTING_KEY,
		DLX_EXCHANGE,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to register DLQ consumer: %v", err)
	}

	ch.QueuePurge(QUEUE, false)

	_ = ch.Close()
	return conn, nil
}

func Main() (bizError error) {
	// 查看当前 GOMAXPROCS
	fmt.Printf("Default GOMAXPROCS: %d	", runtime.GOMAXPROCS(0))

	// 设置为CPU核心数
	runtime.GOMAXPROCS(20)

	fmt.Printf("Current GOMAXPROCS: %d	", runtime.GOMAXPROCS(0))
	defer common.RecoverHandler(func(err error) {
		bizError = err
	})

	// 初始化 sonic 编码器
	jsonEncoder := sonic.ConfigDefault
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	conn, err := connectionInit()
	common.ErrorHandler(err)

	// 构造生产者
	p, err := GetMqProductInstance(conn)
	common.ErrorHandler(err)

	// 构造消费者
	for i := 0; i < 10; i++ {

		c, err := GetMqConsumeInstance(conn, p, ctx)
		common.ErrorHandler(err)
		go func() {
			c.Consume(func(bytes []byte) error {

				hi := repository.HotelInfo{}
				err := jsonEncoder.Unmarshal(bytes, &hi)
				if err != nil {
					return fmt.Errorf("json unmarshal error::%v", err)
				}
				err = getHotelStaticInfoFromSZJLAndInsertDb(hi)
				if err != nil {
					logs.Notice(err, hi.HotelId)
					return err
				}
				wg.Done()
				return nil
			})
		}()

	}

	// 构造死信队列
	dead, err := GetMqDeadLetterInstance(conn, ctx)
	common.ErrorHandler(err)
	go func() {
		dead.Consume(func(bytes []byte) {
			// 记录死信消息
			logs.Warning("========Dead letter message received: %s", string(bytes))
			wg.Done()
		})
	}()

	// 发送生产者消息
	var limitSystemInit = rate.NewLimiter(rate.Limit(50), 5)
	var count int32 = 0
	dbHotelArrayChan := getHotelInfoFromDb()
	for dbHotelArray := range dbHotelArrayChan {

		for _, record := range dbHotelArray {
			// 限流
			_ = limitSystemInit.Wait(context.Background())

			b, err := jsonEncoder.Marshal(record)
			if err != nil {
				logs.Error(fmt.Errorf("批量发送错误::%v , hotelInfo::%+v", err, record))
			}
			err = p.Publish(b, 0)
			if err != nil {
				logs.Error(fmt.Errorf("批量发送错误::%v , hotelInfo::%+v", err, record))
			}
			atomic.AddInt32(&count, 1)
			wg.Add(1)
		}
	}

	logs.Info("发送完成：：", count)

	wg.Wait()
	cancel()
	p.Close()

	logs.Info("处理完成：：", count)
	return nil
}
