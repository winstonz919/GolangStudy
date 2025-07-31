/*
 * @Author: WinstonRD
 * @Date: 2025-07-02 20:40:50
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-02 22:17:44
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/14-goroutine/test3/main.go
 * @Description: 交替打印数字，两个channel和1个channel的不同算法
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
)

// 用两个通道互相写入数据实现交替打印
func chann2() {
	wg := sync.WaitGroup{}
	ch1, ch2 := make(chan struct{}), make(chan struct{})
	defer func() {
		close(ch1)
		close(ch2)
	}()
	wg.Add(1)
	go func() {
		i := -1
		for {
			select {
			case <-ch1:
				i += 2
				if i > 10 {
					wg.Done()
					break
				}
				fmt.Println("ch1:", i)
				ch2 <- struct{}{}
			}
		}
	}()

	go func() {
		i := 0
		for {
			select {
			case <-ch2:
				i += 2
				fmt.Println("ch2:", i)
				ch1 <- struct{}{}
			}
		}
	}()

	ch1 <- struct{}{}
	wg.Wait()

}

// 两个协程监听同一个通道
func chann1() {
	wg := sync.WaitGroup{}
	i := 1
	ch := make(chan int)
	defer close(ch)
	wg.Add(1)
	go func() {
		for {
			select {
			case i = <-ch:
				fmt.Println("go1:", i)
				i++
				if i >= 10 {
					fmt.Println("done")
					wg.Done()
					return
				}
				ch <- i
			}
		}
	}()

	go func() {
		for {
			select {
			case i = <-ch:
				fmt.Println("go2:", i)
				i++
				ch <- i
			}
		}
	}()

	ch <- i
	wg.Wait()
}

func main() {
	// chann2()
	chann1()
}
