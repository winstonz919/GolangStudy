/*
 * @Author: WinstonRD
 * @Date: 2025-06-24 17:39:17
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-24 23:22:51
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/20-context/main.go
 * @Description: 上下文
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"context"
	"fmt"
	"time"
)

// func doSth(ctx context.Context) {
// 	fmt.Println(ctx.Value("name"))
// }

func fCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("ctx is done")
	}
}

func fTimeOut() {
	quit := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	go func() {
		time.Sleep(200 * time.Millisecond)
		quit <- 0
	}()
	select {
	case <-ctx.Done():
		fmt.Println("ctx done:", ctx.Err())
	case <-quit:
		fmt.Println("everything is done")
	}
}

func main() {
	// fCancel()
	fTimeOut()
}

// func SlowOperation(ctx context.Context) {
// 	done := make(chan int, 1)
// 	go func() { // 模拟慢操作
// 		dur := time.Duration(rand.Intn(5)+1) * time.Second
// 		time.Sleep(dur)
// 		done <- 1
// 	}()

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("SlowOperation timeout:", ctx.Err())
// 	case <-done:
// 		fmt.Println("Complete work")
// 	}
// }
