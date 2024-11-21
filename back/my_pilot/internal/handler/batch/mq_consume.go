package batch

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gohouse/t"
	"github.com/streadway/amqp"
	"log"
	"sync/atomic"
)

type Consume struct {
	no         int32
	exchange   string
	key        string
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
	product    *Product
	ctx        context.Context
}

//	func (c Consume) Close() {
//		c.Connection.Close()
//		c.Channel.Close()
//	}
func (c Consume) Consume(process func([]byte) error) {

	tag := "hotel_consume_tag" + t.New(c.no).String()

	defer func() {
		// 按照正确的顺序关闭资源
		//if err := c.Channel.Cancel(tag, false); err != nil {
		//	log.Printf("Error canceling consumer: %v", err)
		//}

		if err := c.Channel.Close(); err != nil {
			log.Printf("Error closing channel: %v", err)
		}

		if err := c.Connection.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}

		log.Println(tag + " Successfully closed all AMQP resources")
	}()

	//   开始消费
	messages, err := c.Channel.Consume(
		c.Queue.Name, // 队列名
		tag,          // 消费者标签
		false,        // 自动确认
		false,        // 排他
		false,        // no-local
		false,        // no-wait
		nil,          // 参数
	)
	if err != nil {
		logs.Error("Failed to register a consumer: %v", err)
	}

	// 消息处理循环
	for {
		select {
		case <-c.ctx.Done():
			log.Println(tag + " Consumer shutting down...")
			return
		case d, ok := <-messages:
			if !ok {
				log.Println("Message channel closed", ok)
				return
			}

			// 获取重试次数
			retryCount := 0
			if d.Headers["retry_count"] != nil {
				retryCount = t.New(d.Headers["retry_count"]).Int()
			}
			// 处理消息（这里模拟处理过程）
			log.Printf("%s :: Received a message: %s", tag, d.Body)
			err := process(d.Body)
			if err != nil {
				if retryCount >= 3 {
					// 超过最大重试次数，发送到死信队列
					logs.Warning("Message failed after %d retries, sending to DLQ", retryCount)
					d.Reject(false) // 拒绝消息，不重新入队
				} else {
					// 重试逻辑:重试记录+1，重新发回队列等待消费
					err = c.product.Publish(d.Body, retryCount+1)
					fmt.Println(fmt.Sprintf("重新发送 retryCount:%v , body:%v", retryCount+1, string(d.Body)))
					if err != nil {
						logs.Error("Failed to publish a message: %v", err)
					}
					d.Ack(false) // 确认原消息
				}
			} else {
				// 处理成功
				d.Ack(false)
			}
		}
	}
}

var consumeCount int32 = 0

func GetMqConsumeInstance(p *Product, ctx context.Context) (*Consume, error) {
	no := atomic.AddInt32(&consumeCount, 1)
	conn, ch, err := initConnection()
	if err != nil {
		return nil, fmt.Errorf("Failed to initConnection: %v", err)
	}
	// 声明交换机
	exchange := "hotel_exchange"
	key := "hotel_key"
	// 0. 声明队列 （带死信配置）
	// 一旦reject，会自动发送死信队列
	args := amqp.Table{
		"x-dead-letter-exchange":    "hotel_dlx_exchange",
		"x-dead-letter-routing-key": "dlx_key",
	}
	q, err := ch.QueueDeclare(
		"hotel_consume_queue", //队列名
		false,                 //是否持久化
		false,                 //是否自动删除
		false,                 //是否排他
		false,                 //是否阻塞
		args,                  //额外属性
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	// 1. 设置QoS
	err = ch.Qos(
		1,     // prefetch count = 1，每次只处理一条消息、预取数量，为了可以让消费者平分
		0,     // prefetch size = 0，不限制消息大小
		false, // global = false，设置应用于每个消费者
	)
	// 2. 绑定队列到交换机
	err = ch.QueueBind(
		q.Name,   // 队列名
		key,      // 路由键
		exchange, // 交换机名
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	return &Consume{
		exchange:   exchange,
		no:         no,
		Queue:      q,
		key:        key,
		Connection: conn,
		Channel:    ch,
		ctx:        ctx,
		product:    p,
	}, nil

}

//func init() {
//	conn, ch := initConnection()
//	defer conn.Close()
//	defer ch.Close()
//	// 0. 声明队列 （带死信配置）
//	// 一旦reject，会自动发送死信队列
//	args := amqp.Table{
//		"x-dead-letter-exchange":    "hotel_dlx_exchange",
//		"x-dead-letter-routing-key": "dlx_key",
//	}
//	q, err := ch.QueueDeclare(
//		"hotel_queue_1", //队列名
//		false,           //是否持久化
//		false,           //是否自动删除
//		false,           //是否排他
//		false,           //是否阻塞
//		args,            //额外属性
//	)
//	if err != nil {
//		log.Fatalf("Failed to declare a queue: %v", err)
//	}
//	// 1. 设置QoS
//	err = ch.Qos(
//		1,     // prefetch count = 1，每次只处理一条消息、预取数量，为了可以让消费者平分
//		0,     // prefetch size = 0，不限制消息大小
//		false, // global = false，设置应用于每个消费者
//	)
//	// 2. 绑定队列到交换机
//	err = ch.QueueBind(
//		q.Name,           // 队列名
//		"hotel_key",      // 路由键
//		"hotel_exchange", // 交换机名
//		false,
//		nil,
//	)
//	if err != nil {
//		log.Fatalf("Failed to bind queue: %v", err)
//	}
//	// 3. 开始消费
//	msgs, err := ch.Consume(
//		q.Name,            // 队列名
//		"hotel_consume_2", // 消费者标签
//		false,             // 自动确认
//		false,             // 排他
//		false,             // no-local
//		false,             // no-wait
//		nil,               // 参数
//	)
//	if err != nil {
//		log.Fatalf("Failed to register a consumer: %v", err)
//	}
//
//	forever := make(chan struct{})
//
//	go func() {
//		for d := range msgs {
//			// 获取重试次数
//			retryCount := 0
//			if d.Headers["retry_count"] != nil {
//				retryCount = t.New(d.Headers["retry_count"]).Int()
//			}
//			// 处理消息（这里模拟处理过程）
//			log.Printf("Received a message: %s", d.Body)
//			err := processMessage(d.Body)
//			if err != nil {
//				if retryCount >= 3 {
//					// 超过最大重试次数，发送到死信队列
//					log.Printf("Message failed after %d retries, sending to DLQ", retryCount)
//					d.Reject(false) // 拒绝消息，不重新入队
//				} else {
//					// 重试逻辑:重试记录+1，重新发回队列等待消费
//					headers := amqp.Table{
//						"retry_count": retryCount + 1,
//					}
//					err = ch.Publish(
//						q.Name,
//						"hotel_key1",
//						false,
//						false,
//						amqp.Publishing{
//							ContentType: "text/plain",
//							Body:        d.Body,
//							Headers:     headers,
//						})
//					if err != nil {
//						log.Fatalf("Failed to publish a message: %v", err)
//					}
//					d.Ack(false) // 确认原消息
//				}
//			} else {
//				// 处理成功
//				d.Ack(false)
//			}
//		}
//	}()
//
//	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
//	<-forever
//}
