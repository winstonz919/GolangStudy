/*
 * @Author: WinstonRD
 * @Date: 2025-06-20 21:11:28
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-20 21:49:59
 * @FilePath: /GolangStudy/18-generic/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
)

type MyType[T int | float32 | string] []T

// 泛型函数
func Add[T int | float32 | float64 | string](a T, b T) T {
	return a + b
}

// 抽象为类型集合
type Number interface {
	int | float32 | float64
}

// 用类型集合代替冗长的类型列表
type Slice[T Number] []T

// any为interface{}的别名
func Print(a any) {
	fmt.Println(a)
}

func main() {
	var a MyType[int] = MyType[int]{1, 2, 3}
	fmt.Println(a)
	var b MyType[float32] = MyType[float32]{1.0, 2.0, 3.0}
	fmt.Println(b)
	var c MyType[string] = MyType[string]{"a", "b", "c"}
	fmt.Println(c)

	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.1, 2.1))
	fmt.Println(Add("a", "b"))

	var s Slice[int] = Slice[int]{1, 2, 3}
	fmt.Println(s)

	Print("any是interface{}的别名")

}
