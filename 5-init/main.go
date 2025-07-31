/*
 * @Author: WinstonRD
 * @Date: 2025-03-25 23:45:26
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-03-26 00:41:47
 * @FilePath: /GolangStudy/5-init/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	_ "GolangStudy/5-init/lib1"      // 匿名导入包，只执行包中的init函数，无法使用包中的函数
	mylib2 "GolangStudy/5-init/lib2" // 导入包，并使用包名调用包中的函数
	// . "GolangStudy/5-init/lib2" // 导入包，直接调用包中的函数，不推荐使用
)

// init函数会在main函数之前执行
// init函数不能被手动调用
// 一个包中可以包含多个init函数
// init函数没有参数也没有返回值
// init函数的执行顺序是按照代码的顺序，而不是按照文件名的顺序
// init函数在包被引用时执行，而不是在包被定义时执行
// init函数在包被引用时只执行一次
func main() {
	// lib1.Lib1Test()
	mylib2.Lib2Test()
}
