package batch

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/streadway/amqp"
	"sync"
)

type Product struct {
	exchange   string
	key        string
	Connection *amqp.Connection
	Channel    *amqp.Channel
	once       sync.Once
}

func (p *Product) obtainChannel() *amqp.Channel {
	p.once.Do(func() {
		channel, _ := p.Connection.Channel()
		p.Channel = channel
	})
	return p.Channel
}

func (p *Product) Close() {
	if p.Channel != nil {
		p.Channel.Close()
	}
	logs.Info("Product closed")
}

func (p *Product) Publish(payload []byte, retryCount int) error {

	// 发送消息
	err := p.obtainChannel().Publish(
		p.exchange, // 交换器
		p.key,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload,
			Headers: amqp.Table{
				"retry_count": retryCount, // 声明重试次数
			},
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}
	return nil
}

func GetMqProductInstance(conn *amqp.Connection) (*Product, error) {

	return &Product{
		exchange:   EXCHANGE,
		key:        ROUTING_KEY,
		Connection: conn,
	}, nil

}
