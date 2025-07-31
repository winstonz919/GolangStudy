/*
 * @Author: WinstonRD
 * @Date: 2025-06-20 14:56:56
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-20 15:16:23
 * @FilePath: /GolangStudy/15-select/test1/main.go
 * @Description: select case的随机选取
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := range 5 {
			ch1 <- i
		}
	}()

	go func() {
		for i := range 5 {
			ch2 <- i
		}
	}()

	time.Sleep(2 * time.Second)

loop:
	for {
		select {
		case <-ch1:
			fmt.Println("case1")
		case <-ch2:
			fmt.Println("case2")
		default:
			break loop
		}
	}
}
