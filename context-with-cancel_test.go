package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return // stop goroutine and close 'destination' channel
			default:
				destination <- counter // send counter to destination
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // default (goroutine = 2)

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	cancel() // send cancel signal to context, then context will done and stopped

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}
