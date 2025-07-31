/*
 * @Author: WinstonRD
 * @Date: 2025-06-25 00:24:09
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-25 15:53:16
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/21-atomic/main.go
 * @Description: 原子
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func add() {
	var counter int32

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&counter, 1) // 原子操作
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func cas() {
	var i int64 = 6789
	oldValue := atomic.SwapInt64(&i, 5567)
	fmt.Println(oldValue, i)
	swapped := atomic.CompareAndSwapInt64(&i, oldValue, 9089) // 交换失败
	fmt.Println(swapped, i)
}

// 自旋锁
type spin int64

func (l *spin) lock() bool {
	for {
		if atomic.CompareAndSwapInt64((*int64)(l), 0, 1) {
			return true
		}
		continue
	}
}

func (l *spin) unlock() bool {
	for {
		if atomic.CompareAndSwapInt64((*int64)(l), 1, 0) {
			return true
		}
		continue
	}
}

func spinLock() {
	s := new(spin)
	for i := 0; i < 5; i++ {
		s.lock()
		go func(i int) {
			fmt.Println(i)
			s.unlock()
		}(i)
	}
}

func main() {
	// add()
	// cas()
	spinLock()
	for {

	}
}
