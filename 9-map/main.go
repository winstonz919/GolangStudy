/*
 * @Author: WinstonRD
 * @Date: 2025-03-27 03:07:09
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-17 13:22:42
 * @FilePath: /GolangStudy/9-map/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	// map需要make初始化后才能使用
	var map1 = make(map[string]int, 2) // 2表示预估容量，避免频繁扩容
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	fmt.Println(map1)
}
