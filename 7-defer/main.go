/*
 * @Author: WinstonRD
 * @Date: 2025-03-26 01:09:39
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-03-26 01:20:31
 * @FilePath: /GolangStudy/7-defer/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

func deferFunc() {
	fmt.Println("deferFunc")
}

func returnFunc() int {
	fmt.Println("returnFunc")
	return 1
}

// defer 语句会将函数推迟到外层函数返回之后执行
func deferAndReturn() int {
	// 延迟执行deferFunc函数
	defer deferFunc()
	// 返回returnFunc函数的返回值
	return returnFunc()
}

func main() {
	// defer 语句会将函数推迟到外层函数返回之后执行
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
	// defer 语句通常用于简化函数返回时的资源清理工作，例如关闭文件、释放锁等
	// defer 语句可以用于任何函数，包括匿名函数和闭包
	// defer fmt.Println("defer 1")
	// defer fmt.Println("defer 2")
	deferAndReturn()
}
