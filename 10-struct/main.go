/*
 * @Author: WinstonRD
 * @Date: 2025-06-17 13:30:17
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-17 13:36:11
 * @FilePath: /GolangStudy/10-struct/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

type Myint int

type Book struct {
	title string
	auth  string
}

func changeBook(b Book) {
	b.title = "Python"
}

func changeBook2(b *Book) {
	b.title = "Python"
}

func main() {
	// var a Myint = 10
	// fmt.Printf("a is %T", a)
	var b Book
	b.title = "Golang"
	b.auth = "WinstonRD"
	fmt.Println(b)
	changeBook(b)
	fmt.Println(b)
	changeBook2(&b)
	fmt.Println(b)
}
