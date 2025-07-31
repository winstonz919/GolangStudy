/*
 * @Author: WinstonRD
 * @Date: 2025-03-25 22:44:59
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-03-25 22:55:26
 * @FilePath: /GolangStudy/var/test1_var.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

var A int = 100
var B = 100

func main() {
	var a int
	fmt.Println(a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 100
	fmt.Println(b)
	fmt.Printf("type of b = %T\n", b)

	var c = 100
	fmt.Println(c)
	fmt.Printf("type of c = %T\n", c)

	d := 100
	fmt.Println(d)
	fmt.Printf("type of d = %T\n", d)

	// 仅在函数内使用，全局变量不能使用 := 赋值
	e := 3.14
	fmt.Println(e)
	fmt.Printf("type of e = %T\n", e)

	fmt.Println(A, B)

	// 多重赋值
	var xx, yy = 100, "abc"
	fmt.Println(xx, yy)
	fmt.Printf("type of xx = %T, type of yy = %T\n", xx, yy)
	var (
		xxx = 100
		yyy = "abc"
	)
	fmt.Println(xxx, yyy)
}
