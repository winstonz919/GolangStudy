/*
 * @Author: WinstonRD
 * @Date: 2025-06-19 19:37:08
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-20 15:13:01
 * @FilePath: /GolangStudy/15-select/main.go
 * @Description:	select 用于同时监控多个channel的状态
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// 定义一个fibonacii函数，接收两个通道参数ch和quit
func fibonacii(ch, quit chan int) {
	// 初始化x和y为1
	x, y := 1, 1
	// 无限循环
	for {
		// 使用select语句，监听ch和quit通道
		select {
		// 如果ch通道有数据可读，则将x写入ch通道，并将x和y更新为y和x+y
		case ch <- x:
			x, y = y, x+y
			// fmt.Println(x, y)
		// 如果quit通道有数据可读，则退出循环
		case <-quit:
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)

	go func() {
		for range 10 {
			fmt.Println(<-ch)
		}

		quit <- 0
	}()

	fibonacii(ch, quit)
}
