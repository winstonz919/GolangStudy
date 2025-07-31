/*
 * @Author: WinstonRD
 * @Date: 2025-07-09 22:40:54
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-09 23:42:21
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/8-外观模式/main.go
 * @Description: 外观模式，又叫门面模式，通过为多个复杂的子系统提供一个统一调用的外观角色
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

/* --------- 模块A接口 -------- */
type AModuleApi interface {
	EchoA()
}

/* --------- 模块A实现 -------- */
type AModuleImpl struct{}

func NewAModuleImpl() AModuleApi {
	return &AModuleImpl{}
}

func (ami *AModuleImpl) EchoA() {
	fmt.Println("Echo testing: A")
}

/* --------- 模块B接口 -------- */
type BModuleApi interface {
	EchoB()
}

/* --------- 模块B实现 -------- */
type BModuleImpl struct{}

func NewBModuleImpl() BModuleApi {
	return &BModuleImpl{}
}

func (bmi *BModuleImpl) EchoB() {
	fmt.Println("Echo testing: B")
}

/* --------- 统一接口 --------- */
type Api interface {
	Echo()
	EchoA()
}

/* -------- 统一接口实现 -------- */
type ApiImpl struct {
	a AModuleApi
	b BModuleApi
}

func NewApi() Api {
	return &ApiImpl{
		a: NewAModuleImpl(),
		b: NewBModuleImpl()}
}

// 提供不同方法来实现调用
func (a *ApiImpl) Echo() {
	a.a.EchoA()
	a.b.EchoB()
}

func (a *ApiImpl) EchoA() {
	a.a.EchoA()
}

// 业务调用
func main() {
	var api Api = NewApi()
	api.Echo()
	// output:
	// 	Echo testing: A
	// 	Echo testing: B
	api.EchoA() // output: Echo testing: A
}
