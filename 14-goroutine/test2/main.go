/*
 * @Author: WinstonRD
 * @Date: 2025-06-19 19:18:26
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-19 19:32:35
 * @FilePath: /GolangStudy/14-goroutine/test2/main.go
 * @Description: 解决因通道读写问题造成的死锁
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("input:", i)
		}
		close(ch) // 如果不关闭通道将造成死锁
	}()

	// for {
	// 	if i, ok := <-ch; ok {
	// 		fmt.Println("output:", i)
	// 	} else {
	// 		break
	// 	}
	// }
	// 从channel中取数据，以上可用range简写
	for i := range ch {
		fmt.Println("output:", i)
	}

	fmt.Println("all done...")
}
