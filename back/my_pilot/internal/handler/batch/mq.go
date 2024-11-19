package batch

import (
	"fmt"
	"github.com/streadway/amqp"
)

// 初始化RabbitMQ连接
func initConnection() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	return conn, ch, nil
}

type Product struct {
	exchange   string
	key        string
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (p Product) Close() {
	p.Connection.Close()
	p.Channel.Close()
}

func (p Product) Publish(payload []byte) error {

	// 发送消息
	err := p.Channel.Publish(
		p.exchange, // 交换器
		p.key,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
			Headers: amqp.Table{
				"retry_count": 0, // 声明重试次数
			},
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}
	return nil
}

func GetMqProductInstance() (*Product, error) {

	conn, ch, err := initConnection()
	if err != nil {
		return nil, fmt.Errorf("Failed to initConnection: %v", err)
	}
	// 声明交换机
	exchange := "hotel_exchange"
	key := "hotel_key"
	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to declare exchange: %v", err)
	}
	return &Product{
		exchange:   exchange,
		key:        key,
		Connection: conn,
		Channel:    ch,
	}, nil

}
