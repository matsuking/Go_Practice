package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func timeout_context() {
	// セレクトを使うことで複数のチャネルを受け取ることができる。
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Microsecond)
		ch1 <- "A"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(800 * time.Microsecond)
		ch1 <- "B"
	}()
loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			break loop
		case v := <-ch1:
			fmt.Println(v)
			ch1 = nil
		case v := <-ch2:
			fmt.Println(v)
			ch2 = nil
		}
	}

	wg.Wait()
	fmt.Println("finish")
}
