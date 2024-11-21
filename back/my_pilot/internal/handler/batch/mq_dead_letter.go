package batch

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/streadway/amqp"
	"log"
)

type DeadLetterConsume struct {
	exchange   string
	key        string
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
	ctx        context.Context
}

func (c DeadLetterConsume) Consume(process func([]byte)) {

	defer func() {

		if err := c.Channel.Close(); err != nil {
			log.Printf("Error closing channel: %v", err)
		}

		if err := c.Connection.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}

		log.Println("  DeadLetterConsume Successfully closed all AMQP resources")
	}()

	messages, err := c.Channel.Consume(
		c.Queue.Name,
		"hotel_dlx_consume",
		true, // 自动确认
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logs.Error("Failed to register a consumer: %v", err)
	}
	// 消息处理循环
	for {
		select {
		case <-c.ctx.Done():
			log.Println("DeadLetterConsume" + " Consumer shutting down...")
			return
		case d, ok := <-messages:
			if !ok {
				log.Println("DeadLetterConsume"+" Message channel closed", ok)
				return
			}
			process(d.Body)
		}
	}
}
func GetMqDeadLetterInstance(ctx context.Context) (*DeadLetterConsume, error) {
	conn, ch, err := initConnection()
	if err != nil {
		return nil, fmt.Errorf("Failed to initConnection: %v", err)
	}
	// 声明死信交换机
	exchange := "hotel_dlx_exchange"
	key := "dlx_key"
	// 声明死信交换机
	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logs.Error("Failed to declare dead letter exchange: %v", err)
	}

	// 声明死信队列
	dlq, err := ch.QueueDeclare(
		"hotel_dlx_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logs.Error("Failed to declare dead letter queue: %v", err)
	}

	// 绑定死信队列到死信交换机
	err = ch.QueueBind(
		dlq.Name,
		key,
		exchange,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to register DLQ consumer: %v", err)
	}

	return &DeadLetterConsume{
		exchange:   exchange,
		Queue:      dlq,
		key:        key,
		ctx:        ctx,
		Connection: conn,
		Channel:    ch,
	}, nil
	//msgs, err := ch.Consume(
	//	"hotel_dlx_exchange",
	//	"hotel_dlx_consume",
	//	true, // 自动确认
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//
	//forever := make(chan bool)
	//
	//go func() {
	//	for d := range msgs {
	//		// 记录死信消息
	//		log.Printf("Dead letter message received: %s", string(d.Body))
	//	}
	//}()
	//
	//log.Printf(" [*] Waiting for dead letter messages. To exit press CTRL+C")
	//<-forever

}
