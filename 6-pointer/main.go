/*
 * @Author: WinstonRD
 * @Date: 2025-03-26 00:50:08
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-03-26 01:09:25
 * @FilePath: /GolangStudy/6-pointer/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

func copyValue(x int) {
	fmt.Println("x is", x)
	x = 10
	fmt.Println("x is", x)
}

func changeValue(x *int) {
	*x = 10
	fmt.Println("x is", *x)
}

func swap(x, y *int) {
	*x, *y = *y, *x
	fmt.Println("x is", *x, "y is", *y)
}

func main() {
	x, y := 1, 2
	copyValue(x)
	changeValue(&x)
	fmt.Println("y is", y)
	swap(&x, &y)

	var a int = 100
	var p *int
	fmt.Println("p is", p)
	// *p = 100
	p = &a
	fmt.Println("p is", p, "value of p is", *p)

	// 二级指针，指向指针的指针
	var pp **int = &p
	fmt.Println("pp is", pp, "value of pp is", **pp)
}
