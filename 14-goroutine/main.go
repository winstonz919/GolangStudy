/*
 * @Author: WinstonRD
 * @Date: 2025-06-19 18:08:17
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-19 18:10:07
 * @FilePath: /GolangStudy/14-goroutine/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"
)

func newTask() {
	var i int
	for {
		i++
		fmt.Println("new goroutine: i=", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()

	var i int
	for {
		i++
		fmt.Println("main goroutine: i=", i)
		time.Sleep(1 * time.Second)
	}
}
