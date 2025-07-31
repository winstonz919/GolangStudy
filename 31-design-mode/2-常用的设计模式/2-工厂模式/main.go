/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 12:46:27
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-08 12:58:07
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/2-工厂模式/main.go
 * @Description: 工厂模式
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// ----抽象层-----
type Fruit interface {
	Show()
}

type FruitFactory interface {
	CreateFruit() Fruit
}

// --------实现层---------
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("I'm apple")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("I'm banana")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("I'm pear")
}

// --------工厂模块--------
type AppleFactory struct {
	FruitFactory
}

func (af *AppleFactory) CreateFruit() (fruit Fruit) {
	fruit = &Apple{}
	return
}

type BananaFactory struct {
	FruitFactory
}

func (bf *BananaFactory) CreateFruit() (fruit Fruit) {
	fruit = &Banana{}
	return
}

type PearFactory struct {
	FruitFactory
}

func (pf *PearFactory) CreateFruit() (fruit Fruit) {
	fruit = &Pear{}
	return
}

// 业务逻辑只调用抽象层
func main() {
	var appleFac FruitFactory
	appleFac = &AppleFactory{}
	var fruit Fruit
	fruit = appleFac.CreateFruit()
	fruit.Show()
}
