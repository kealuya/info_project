package batch

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/streadway/amqp"
)

type DeadLetterConsume struct {
	exchange   string
	key        string
	queue      string
	Connection *amqp.Connection
	Channel    *amqp.Channel
	ctx        context.Context
}

func (c DeadLetterConsume) Consume(process func([]byte)) {

	defer func() {

		if err := c.Channel.Close(); err != nil {
			logs.Error("Error closing channel: %v", err)
		}

		logs.Info("  DeadLetterConsume Successfully closed all AMQP resources")
	}()

	messages, err := c.Channel.Consume(
		c.queue,
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
			logs.Info("DeadLetterConsume" + " Consumer shutting down...")
			return
		case d, ok := <-messages:
			if !ok {
				logs.Error("DeadLetterConsume"+" Message channel closed", ok)
				return
			}
			process(d.Body)
		}
	}
}
func GetMqDeadLetterInstance(conn *amqp.Connection, ctx context.Context) (*DeadLetterConsume, error) {
	ch, _ := conn.Channel()
	return &DeadLetterConsume{
		exchange:   DLX_EXCHANGE,
		queue:      DLX_QUEUE,
		ctx:        ctx,
		Connection: conn,
		Channel:    ch,
	}, nil

}
