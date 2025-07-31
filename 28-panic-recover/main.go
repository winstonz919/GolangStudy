/*
 * @Author: WinstonRD
 * @Date: 2025-07-01 14:10:26
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-01 14:53:47
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/28-panic-recover/main.go
 * @Description: 异常捕获
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

func err2() {
	panic("err2")
}

func main() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	err2()
	panic("err1")
}
