/*
 * @Author: WinstonRD
 * @Date: 2025-07-01 22:31:08
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-02 00:53:42
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/15-select/test2/main.go
 * @Description: 多路复用
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var ch1 = make(chan int)
var ch2 = make(chan int)
var quit = make(chan struct{})
var wg sync.WaitGroup

func sendMessage(ch chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func selectMessage() {
	for {
		select {
		case i, ok := <-ch1:
			fmt.Println("ch1:", i, ok)
			time.Sleep(500 * time.Millisecond)
		case i, ok := <-ch2:
			fmt.Println("ch2", i, ok)
			time.Sleep(500 * time.Millisecond)
		case i, ok := <-quit:
			fmt.Println("all done:", i, ok)
			return
		default:
			fmt.Println("default")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	go selectMessage() // 监听多个通道
	wg.Add(1)
	go sendMessage(ch1)
	wg.Add(1)
	go sendMessage(ch2)

	wg.Wait()

	close(ch1)
	close(ch2)
	close(quit) // 关闭通道

}
