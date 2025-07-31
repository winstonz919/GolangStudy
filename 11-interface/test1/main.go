/*
 * @Author: WinstonRD
 * @Date: 2025-07-02 13:19:17
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-02 19:46:54
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/11-interface/test1/main.go
 * @Description: 抽象依赖
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Eat() {
	fmt.Println("eat")
}

func (a *Animal) Drink() {
	fmt.Println("drink")
}

type Cat struct {
}

type ICat interface {
	Eat()
	Drink()
}

func main() {
	var cat ICat
}
