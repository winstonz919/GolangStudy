/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 09:55:03
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-08 10:02:48
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/31-design-mode/1-单一职责/main.go
 * @Description: 单一指责原则
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// 对外只提供一种功能，而引起变化的原因都应该只有一个
type ClothesShop struct{}

func (c *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

type ClothesWork struct{}

func (c *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}

func main() {
	cs := ClothesShop{}
	cs.Style()

	cw := ClothesWork{}
	cw.Style()
}
