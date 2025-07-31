/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 12:26:52
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-08 12:44:23
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/1-工厂模式/main.go
 * @Description: 简单工厂模式实现
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// 简单理解为作为使用方不需要直接同产品打交道，而只需要同中间的工厂交涉

// --------抽象层---------
type Fruit interface {
	Show()
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
type Factory struct{}

// 工厂模式违背了开闭原则
func (factory *Factory) CreateFruit(kind string) (fruit Fruit) {
	switch kind {
	case "apple":
		fruit = &Apple{}
	case "banana":
		fruit = &Banana{}
	case "pear":
		fruit = &Pear{}
	default:
		panic("unknown kind")
	}
	return fruit
}

func main() {
	factory := &Factory{}
	fruit := factory.CreateFruit("apple")
	fruit.Show()
}
