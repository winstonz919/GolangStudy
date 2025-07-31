/*
 * @Author: WinstonRD
 * @Date: 2025-06-18 10:32:49
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-18 11:03:41
 * @FilePath: /GolangStudy/11-interface/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// interface是万能类型，Go语言中所有的类型都实现了interface{}
func myFunc(arg interface{}) {

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
		return
	}
	fmt.Println("type of value is", value)
	fmt.Printf("value type is %T\n", value)
}

type Book struct {
	auth string
}

func main() {
	book := Book{"golang"}
	myFunc(book)
	fmt.Println("==============")
	myFunc("abc")
	fmt.Println("==============")
	myFunc(100)
}
