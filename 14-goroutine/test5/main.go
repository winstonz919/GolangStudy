/*
 * @Author: WinstonRD
 * @Date: 2025-07-02 22:46:08
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-02 23:46:01
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/14-goroutine/test5/main.go
 * @Description: 测试select的随机case
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	// 如果使用无缓冲的通道，必须保证读写同时准备好，即使用goroutine
	ch1, ch2 := make(chan struct{}, 1), make(chan struct{}, 1) // 读写同在主协程中时，要使用缓冲通道

	ch1 <- struct{}{}
	ch2 <- struct{}{}

	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
		panic("error!")
	}

}
