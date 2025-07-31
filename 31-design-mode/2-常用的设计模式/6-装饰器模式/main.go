/*
 * @Author: WinstonRD
 * @Date: 2025-07-09 16:01:03
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-09 20:17:07
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/6-装饰器模式/main.go
 * @Description: 装饰器模式，用于在不改变原有对象的情况下，为对象增加功能，调用不同的装饰器可以增加不同的功能
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// ---- 抽象接口 ----
type ICalcer interface {
	Calc() int
}

// ---- 实现 ----
type Calcer struct{}

func (c Calcer) Calc() int {
	return 0
}

// ---- 装饰器1实现 ----
type MulCalcer struct {
	ICalcer
	Num int
}

func (mc *MulCalcer) Calc() int {
	return mc.ICalcer.Calc() * mc.Num
}

func WarpMuCalcer(i ICalcer, num int) ICalcer {
	return &MulCalcer{
		ICalcer: i,
		Num:     num,
	}
}

// ---- 装饰器2实现 ----
type AddCalcer struct {
	ICalcer
	Num int
}

func (ac *AddCalcer) Calc() int {
	return ac.ICalcer.Calc() + ac.Num
}

func WarpAddCalcer(i ICalcer, num int) ICalcer {
	return &AddCalcer{
		ICalcer: i,
		Num:     num,
	}
}

func main() {
	var i ICalcer = Calcer{}
	i = WarpAddCalcer(i, 5) // 调用装饰器为原对象增加功能
	fmt.Println(i.Calc())
	i = WarpMuCalcer(i, 10) // 调用不同的装饰器实现不同的功能
	fmt.Println(i.Calc())
}
