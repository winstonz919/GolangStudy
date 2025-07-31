/*
 * @Author: WinstonRD
 * @Date: 2025-06-26 19:36:37
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-26 23:51:22
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/25-sync-pool/main.go
 * @Description: 缓存池
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
)

type A struct {
	Name string
}

func (a *A) Reset() {
	a.Name = ""
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(A)
	},
}

// 调用测试
func callTimes() {
	i := 0
	p := sync.Pool{
		New: func() interface{} {
			i++
			return i
		},
	}

	fmt.Println(p.Get()) // 未获取到对象，调用New
	fmt.Println(p.Get()) // 未获取到对象，调用New
	p.Put(0)             // 归还
	fmt.Println(p.Get()) // 获取到对象
}

func main() {
	// a := pool.Get().(*A) // 获取一个对象，如果没有，则调用New方法创建一个
	// a.Reset()
	// defer pool.Put(a) // 放入缓存池
	// a.Name = "recker"
	// fmt.Println(a)
	callTimes()
}
