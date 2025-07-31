/*
 * @Author: WinstonRD
 * @Date: 2025-06-26 00:07:15
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-26 15:09:40
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/24-sync-once/main.go
 * @Description: sync.Once 可用于加载配置文件
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
)

func onceBody() {
	fmt.Println("calling onceBody")
}

func closure() func() {
	return func() {
		fmt.Println("decorator here")
		onceBody()
	}
}

func main() {
	var once sync.Once

	for i := 0; i < 5; i++ {
		once.Do(closure()) // 多次调用，只被执行一次
	}
}
