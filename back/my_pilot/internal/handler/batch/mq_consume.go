package batch

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gohouse/t"
	"github.com/streadway/amqp"
	"sync/atomic"
)

type Consume struct {
	no         int32
	exchange   string
	key        string
	queue      string
	Connection *amqp.Connection
	Channel    *amqp.Channel
	product    *Product
	ctx        context.Context
}

//	func (c Consume) Close() {
//		c.Connection.Close()
//		c.Channel.Close()
//	}
func (c Consume) Consume(process func([]byte) error) {

	tag := "hotel_consume_tag" + t.New(c.no).String()

	// 自动关闭，因为consume是goroutine执行，方法结束后就关闭
	defer func() {
		// 按照正确的顺序关闭资源
		if err := c.Channel.Cancel(tag, false); err != nil {
			logs.Error("Error canceling consumer: %v", err)
		}

		if err := c.Channel.Close(); err != nil {
			logs.Error("Error closing channel: %v", err)
		}

		logs.Info(tag + " Successfully closed all AMQP resources")
	}()

	//   开始消费
	messages, err := c.Channel.Consume(
		c.queue, // 队列名
		tag,     // 消费者标签
		false,   // 自动确认
		false,   // 排他
		false,   // no-local
		false,   // no-wait
		nil,     // 参数
	)
	if err != nil {
		logs.Error("Failed to register a consumer: %v", err)
	}

	// 阻塞~ 消息处理循环
	for {
		select {
		case <-c.ctx.Done():
			logs.Info(tag + " Consumer shutting down...")
			return
		case d, ok := <-messages:
			if !ok {
				logs.Error("Message channel closed", ok)
				return
			}

			// 获取重试次数
			retryCount := 0
			if d.Headers["retry_count"] != nil {
				retryCount = t.New(d.Headers["retry_count"]).Int()
			}
			err := process(d.Body)
			if err != nil {
				if retryCount >= 3 {
					// 超过最大重试次数，发送到死信队列
					logs.Info("Message failed after %d retries, sending to DLQ , %+v", retryCount, string(d.Body))
					d.Reject(false) // 拒绝消息，不重新入队
				} else {
					// 重试逻辑:重试记录+1，重新发回队列等待消费
					err = c.product.Publish(d.Body, retryCount+1)
					if err != nil {
						logs.Error("Failed to publish a message: %v , %+v", err, string(d.Body))
					}
					logs.Info("重新发送 retryCount:%v , body:%v", retryCount+1, string(d.Body))
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

func GetMqConsumeInstance(conn *amqp.Connection, p *Product, ctx context.Context) (*Consume, error) {
	no := atomic.AddInt32(&consumeCount, 1)
	// 因为consume是阻塞，需要go func执行，每一个channel对应一个goroutine
	ch, _ := conn.Channel()
	err := ch.Qos(
		10,    // prefetch count = 1，每次只处理一条消息、预取数量，为了可以让消费者平分
		0,     // prefetch size = 0，不限制消息大小
		false, // global = false，设置应用于每个消费者
	)
	if err != nil {
		return nil, fmt.Errorf("failed to Qos: %v", err)
	}

	return &Consume{
		exchange:   EXCHANGE,
		no:         no,
		queue:      QUEUE,
		Connection: conn,
		Channel:    ch,
		ctx:        ctx,
		product:    p,
	}, nil

}
