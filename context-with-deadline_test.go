package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter3(ctx context.Context) chan int {
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
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // default (goroutine = 2)

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter3(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // goroutine leak (goroutine = 3) even loop is stoped (break)
}
