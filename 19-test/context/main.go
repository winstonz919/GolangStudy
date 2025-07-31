/*
 * @Author: WinstonRD
 * @Date: 2025-06-24 17:39:17
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-24 17:42:42
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/19-test/context/main.go
 * @Description: 上下文
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go SlowOperation(ctx)
	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("goroutine:", runtime.NumGoroutine())
		}
	}()
	time.Sleep(4 * time.Second)

}

func SlowOperation(ctx context.Context) {
	done := make(chan int, 1)
	go func() { // 模拟慢操作
		dur := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(dur)
		done <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("SlowOperation timeout:", ctx.Err())
	case <-done:
		fmt.Println("Complete work")
	}
}
