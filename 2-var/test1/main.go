/*
 * @Author: WinstonRD
 * @Date: 2025-06-25 09:21:44
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-25 12:34:45
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/2-var/test1/main.go
 * @Description: 变量的内存空间
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var a int
	at := reflect.TypeOf(a)
	fmt.Println(runtime.GOARCH, unsafe.Sizeof(a), at.Kind()) // int在64位系统中占64bits，即8个字节

	var b [10]int
	fmt.Println(unsafe.Sizeof(b)) // 数组依数据类型和长度决定占用空间

	var c *int
	fmt.Println(unsafe.Sizeof(c)) // 指针固定占8个字节

	var d string
	d = "你好"
	// 字符串固定占用16字节，字符串实际是一个结构体，包含实际数据的指针uintptr和长度int，各占8个字节
	// 字符串采用utf-8编码，1一个英文字母占1个字节，中文占3个字节，len()方法返回的是占内存空间的大小而不是字符数量
	fmt.Println(unsafe.Sizeof(d), len(d))
	fmt.Println(utf8.RuneCountInString(d))

	for _, v := range d {
		fmt.Println(v) // 你 20320，好 22909，遍历字符串实际输出的是rune，每个字符为int32
	}

	var e []byte
	fmt.Println(unsafe.Sizeof(e))

	type f struct{}

	fmt.Println(unsafe.Sizeof(f{}))

	const g int = 1
	const h = 1.0 + g // 此时操作符左边的1.0为未命名常量可以进行int转换，即隐式转换
	fmt.Println(h)
}
