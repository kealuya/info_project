package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	fmt.Println("start")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("11111")
			break
		default:
			fmt.Println("9999")
		}
	}

	fmt.Println("over")

}
