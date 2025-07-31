/*
 * @Author: WinstonRD
 * @Date: 2025-06-19 13:18:22
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-19 14:57:23
 * @FilePath: /GolangStudy/12-reflect/test1/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"reflect"
)

type Myint int

func main() {
	var myint Myint = 1

	v := reflect.ValueOf(myint)
	// Type是静态类型，是实际使用中定义的类型，Kind是底层最基本的类型
	fmt.Println(v.Type(), v.Kind()) // 输出：main.Myint int
}
