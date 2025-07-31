/*
 * @Author: WinstonRD
 * @Date: 2025-07-09 20:21:43
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-09 21:17:13
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/7-适配器模式/main.go
 * @Description: 适配器模式, 适配器模式用于转换一种接口适配另一种接口
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// --- 适配的目标 ----
type V5 interface {
	Use5V()
}

// ---- 被适配的角色 ----
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("当前电压为220V")
}

func NewV220() *V220 {
	return &V220{}
}

// ---- 适配器 ----
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器")

	// 调用被适配的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

// ---- 业务 ----
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

func (p *Phone) Charge() {
	fmt.Println("手机开始充电")
	p.v.Use5V()
}

func main() {
	phone := NewPhone(NewAdapter(NewV220())) // 通过适配器来匹配接口，间接调用被适配的方法
	phone.Charge()
}
