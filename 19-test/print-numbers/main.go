/*
 * @Author: WinstonRD
 * @Date: 2025-06-24 12:48:06
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-24 18:10:27
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/19-test/print-numbers/main.go
 * @Description: 多协程交替打印数字
 * 	1.互斥锁
 *  2.原子操作
 *  3.channel
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
)

var counter int

var ch = make(chan int, 1)

var wg sync.WaitGroup

func f1() {
	fmt.Println("f1:", counter)
	counter = counter + 1
	ch <- counter
	wg.Done()
}

func f2() {
	fmt.Println("f2:", <-ch)
	wg.Done()
}

func main() {
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()
}
