package batch

import (
	"context"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"strconv"
	"sync"
	"time"
)

func Main() {

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	// 构造生产者
	p, err := GetMqProductInstance()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 构造消费者
	for i := 0; i < 5; i++ {

		c, err := GetMqConsumeInstance(p, ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			c.Consume(func(bytes []byte) error {
				fmt.Println("handle ::" + string(bytes))
				time.Sleep(1 * time.Second)
				// fixme
				if string(bytes) == "3" {
					return errors.New("error~~ ::" + string(bytes))
				}
				wg.Done()
				return nil
			})
		}()

	}

	// 构造死信队列
	dead, err := GetMqDeadLetterInstance(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		dead.Consume(func(bytes []byte) {
			fmt.Println("dead handle ::" + string(bytes))
			// 记录死信消息
			logs.Warning("Dead letter message received: %s", string(bytes))
			wg.Done()
		})
	}()

	// 发送生产者消息
	for i := 0; i < 100; i++ {
		wg.Add(1)
		err := p.Publish([]byte(strconv.Itoa(i)), 0)
		if err != nil {
			fmt.Println("p.Publishs", err)
			//return
		}

	}
	fmt.Println("发送完成", 100)

	wg.Wait()
	cancel()
	fmt.Println("over")
	time.Sleep(2 * time.Second)
	p.Close()
}
