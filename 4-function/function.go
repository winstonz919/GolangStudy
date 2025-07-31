/*
 * @Author: WinstonRD
 * @Date: 2025-03-25 23:09:29
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-17 10:18:33
 * @FilePath: /GolangStudy/4-function/function.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println(a, b)
	c := 100
	return c
}

func foo2(a, b, c int) int {
	return c
}

func foo3(a, b int) (r1 int, r2 int) {
	return r1, r2
}

func main() {
	r1 := foo1("abc", 123)
	fmt.Println("r1 is", r1)

	r2 := foo2(1, 2, 3)
	fmt.Println("r2 is", r2)

	r3, r4 := foo3(1, 2)
	fmt.Println("r3 is", r3, "r4 is", r4)
}
