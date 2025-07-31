/*
 * @Author: WinstonRD
 * @Date: 2025-07-01 15:37:20
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-01 16:11:34
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/29-make-new/main.go
 * @Description: make和new的区别
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// make 的作用是初始化内置的数据结构，也就是我们在前面提到的切片、哈希表和 Channel；
// new 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针；
func main() {

	var i *int
	// *i = 1 // invalid memory address or nil pointer dereference
	fmt.Println(i) // 此时i为nil，无法赋值
	i = new(int)   // new为类型分配了内存空间
	fmt.Println(i) // 此时i为0x14000102028
	*i = 1         // 可以赋值
	fmt.Println(*i)

	// map slice channel属于引用类型，分配内存空间后无法直接赋值
	m := new(map[string]int)
	// (*m)["a"] = 1 // assignment to entry in nil map
	fmt.Println(m)
	*m = make(map[string]int) // 需要使用make初始化内存
	(*m)["a"] = 1
	fmt.Println(m)

	fmt.Println("=================")

	type Food struct {
		Name string
		Type *int
	}

	f := new(Food)
	fmt.Println(f)
	*f.Type = 1 // 此时地址为nil不可写
	fmt.Println(f)
}
