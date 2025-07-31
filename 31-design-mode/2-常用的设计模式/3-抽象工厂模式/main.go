/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 13:09:38
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-08 13:40:55
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/3-抽象工厂模式/main.go
 * @Description: 抽象工厂模式
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// ----抽象层-----
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

type ChinaFactory struct {
	AbstractFactory
}

func (cf *ChinaFactory) CreateApple() (fruit AbstractApple) {
	fruit = &ChinaApple{}
	return
}

func (cf *ChinaFactory) CreateBanana() (fruit AbstractBanana) {
	fruit = &ChinaBanana{}
	return
}

func (cf *ChinaFactory) CreatePear() (fruit AbstractPear) {
	fruit = &ChinaPear{}
	return
}

type ChinaApple struct {
	AbstractApple
}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("I'm a chinese apple")
}

type ChinaBanana struct {
	AbstractBanana
}

func (ca *ChinaBanana) ShowBanana() {
	fmt.Println("I'm a chinese banana")
}

type ChinaPear struct {
	AbstractPear
}

func (ca *ChinaPear) ShowPear() {
	fmt.Println("I'm a chinese pear")
}

func main() {
	// 创建一个工厂
	var factory AbstractFactory
	factory = &ChinaFactory{}

	// 创建一个苹果
	var apple AbstractApple
	apple = factory.CreateApple()
	apple.ShowApple()
}
