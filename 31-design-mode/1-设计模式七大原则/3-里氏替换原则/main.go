/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 10:55:07
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-08 11:17:06
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/31-design-mode/3-里氏替换原则/main.go
 * @Description: 用接口调用逻辑，而不直接依赖具体的类
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// ---------抽象层---------
type ICar interface {
	Run()
}

type Driver interface {
	Drive()
}

// --------实现层----------
type Benz struct{}

func (b *Benz) Run() {
	fmt.Println("Benz is running")
}

type Bmw struct{}

func (b *Bmw) Run() {
	fmt.Println("BMW is running")
}

type Driver1 struct{}

func (d *Driver1) Drive(car ICar) {
	fmt.Println("driver1 drive car")
	car.Run()
}

type Driver2 struct{}

func (d *Driver2) Drive(car ICar) {
	fmt.Println("driver2 drive car")
	car.Run()
}

// ------业务逻辑层---------
func main() {
	var bmw ICar
	bmw = &Bmw{}
	driver := &Driver1{}
	driver.Drive(bmw)

	var benz ICar
	benz = &Benz{}
	driver.Drive(benz)
}
