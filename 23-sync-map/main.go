/*
 * @Author: WinstonRD
 * @Date: 2025-06-25 23:47:48
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-26 00:02:37
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/23-sync-map/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {
	m.Store("a", 1)
	m.Store("b", "111")

	m.Range(func(key any, value any) bool {
		fmt.Println(key, value)
		return true
	})

	fmt.Println(m.Load("aaa"))
}
