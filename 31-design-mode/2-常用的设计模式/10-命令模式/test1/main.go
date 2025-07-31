/*
 * @Author: WinstonRD
 * @Date: 2025-07-10 11:05:11
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-10 13:47:25
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/10-命令模式/test1/main.go
 * @Description: 命令模式练习
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

/* --------- 厨师接口 --------- */
type IChef interface {
	CookChicken()
	CookSheep()
}

// 厨师实现
type Chef struct{}

func NewChef() IChef {
	return &Chef{}
}

func (c *Chef) CookSheep() {
	fmt.Println("厨师烤羊肉")
}

func (c *Chef) CookChicken() {
	fmt.Println("厨师烤鸡肉")
}

// 订单抽象
type Order interface {
	Cook()
}

// 订单实现
type SheepOrderImpl struct {
	chef IChef
}

func NewSheepOrder(chef IChef) Order {
	return &SheepOrderImpl{chef: chef}
}

func (o *SheepOrderImpl) Cook() {
	o.chef.CookSheep()
}

type ChickenOrderImpl struct {
	chef IChef
}

func NewChickenOrder(chef IChef) Order {
	return &ChickenOrderImpl{chef: chef}
}

func (o *ChickenOrderImpl) Cook() {
	o.chef.CookChicken()
}

// 服务员持有订单列表，并负责下发订单给厨师
type IWaiter interface {
	Notify()
	AddOrder(orders ...Order)
}

type Waiter struct {
	IWaiter
	OrderList []Order // 订单列表
}

func NewWaiter() IWaiter {
	return &Waiter{}
}

func (w *Waiter) AddOrder(orders ...Order) {
	for _, order := range orders {
		w.OrderList = append(w.OrderList, order)
	}
}

func (w *Waiter) Notify() {
	if len(w.OrderList) == 0 {
		return
	}

	for _, order := range w.OrderList {
		order.Cook()
	}
}

func main() {
	// 创建厨师、服务员、订单
	var chef IChef = NewChef()
	var waiter IWaiter = NewWaiter()
	waiter.AddOrder(NewChickenOrder(chef), NewSheepOrder(chef))
	waiter.Notify()
}
