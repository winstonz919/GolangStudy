/*
 * @Author: WinstonRD
 * @Date: 2025-07-31 10:03:13
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-31 10:54:10
 * @FilePath: /35-channel/channel_no_buffer/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(i)
}

// 无缓存的通道如果在写后关闭不影响读
func main() {
	var wg sync.WaitGroup

	chw := make(chan int)

	wg.Add(1)
	go func() {
		for k := 0; k < 100; k++ {
			chw <- k
		}
		close(chw) // 发送者应在发送完数据后关闭通道，否则接收者无法正常退出，导致wait永远无法结束
		wg.Done()
	}()

	// 开启10个协程读取
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := range chw {
				process(i)
			}
			wg.Done()
		}()
	}

	wg.Wait()

}
