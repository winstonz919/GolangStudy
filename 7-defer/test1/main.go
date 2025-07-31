/*
 * @Author: WinstonRD
 * @Date: 2025-06-30 23:11:05
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-30 23:39:14
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/7-defer/test1/main.go
 * @Description: 测试匿名和有名返回值情况下，defer的执行顺序
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// 在函数return时先执行return后的语句，将i赋值给临时变量
// 此时defer中的操作是针对i变量，所以不影响返回值
func t1() int {
	i := 0
	defer func() {
		i++
	}()
	return i // return 0
}

func t2() (i int) {
	defer func() {
		i++
	}()
	return // return 1
}

func main() {

	go func() {
		if err := http.ListenAndServe("127.0.0.1:6060", nil); err != nil {
			panic(err)
		}
	}()

	r := t1()
	fmt.Println("r1:", r)
	r = t2()
	fmt.Println("r2:", r)
}
