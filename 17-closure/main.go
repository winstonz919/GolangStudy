/*
 * @Author: WinstonRD
 * @Date: 2025-06-20 19:42:05
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-20 20:19:07
 * @FilePath: /GolangStudy/17-closure/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// 闭包 = 函数 + 周围环境
// 定义一个函数Add，接收一个整数a作为参数，返回一个匿名函数
func Add(a int) func() {
	// 返回一个匿名函数，该函数将a加1，并打印出结果
	return func() {
		a = a + 1
		fmt.Println(a)
	}
}

func MakeFibGen() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	add := Add(1)
	add() // 2
	add() // 3

	// 斐波那契数列
	r := MakeFibGen()
	for range 5 {
		fmt.Println(r())
	}
}
