/*
 * @Author: WinstonRD
 * @Date: 2025-03-25 23:00:05
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-03-25 23:07:57
 * @FilePath: /GolangStudy/const/const.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

const (
	BEIJING  = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

// iota 是一个常量生成器，在 const 关键字出现时将被重置为 0
// iota 可理解为 const 语句块中的行索引
// iota 可以被用作枚举值
// iota 仅能出现在const中，在const中，每新增一行常量声明将使 iota 计数一次
const (
	a, b = iota + 1, iota + 2
	c, d

	e, f = iota * 2, iota * 3
	g, h
)

func main() {
	fmt.Println("BEIJING:", BEIJING)
	fmt.Println("SHANGHAI:", SHANGHAI)
	fmt.Println("SHENZHEN:", SHENZHEN)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("g:", g)
	fmt.Println("h:", h)
}
