/*
 * @Author: WinstonRD
 * @Date: 2025-06-19 18:17:55
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-19 19:09:08
 * @FilePath: /GolangStudy/14-goroutine/test1/main.go
 * @Description: 协程内使用原子操作atomic
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"sync"
	"sync/atomic"
)

// 包变量
var count int32 = 1

// 打印奇数
func printOne() {
	for {
		// 死循环无限取
		num := atomic.LoadInt32(&count)
		if num > 100 {
			return
		}
		// 取的不是想要就不理，继续无限取
		if num%2 == 1 {
			println("奇数：", num)
			atomic.AddInt32(&count, 1)
		}
	}
}

// 打印偶数
func printTwo() {
	for {
		// 死循环无限取
		num := atomic.LoadInt32(&count)
		if num > 100 {
			return
		}
		// 取的不是想要就不理，继续无限取
		if num%2 == 0 {
			println("偶数：", num)
			atomic.AddInt32(&count, 1)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		printOne()
		wg.Done()
	}()
	go func() {
		printTwo()
		wg.Done()
	}()
	wg.Wait()
}
