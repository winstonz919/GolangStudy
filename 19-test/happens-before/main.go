/*
 * @Author: WinstonRD
 * @Date: 2025-06-24 15:02:31
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-24 15:53:41
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/19-test/happens-before/main.go
 * @Description: 对于有缓冲的队列，写在读之前发生，对于无缓冲的队列，读在写之前发生
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

var c = make(chan int, 5)
var ch = make(chan int)
var a string

func f() {
	a = "hello world"
	<-c
}

func main() {
	go f()
	c <- 0
	fmt.Println(a) // 输出空字符串，因为c是有缓冲通道，那么写操作在读之前执行
}
