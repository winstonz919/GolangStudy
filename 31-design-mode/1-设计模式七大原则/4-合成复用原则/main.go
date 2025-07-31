/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 11:22:45
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-08 11:34:09
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/31-design-mode/4-合成复用原则/main.go
 * @Description: 合成复用原则
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 继承
type CatB struct {
	Cat
}

func (c *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// 组合，优先使用，降低模块间的依赖关系，使用组合可以选择性的继承父类的方法
type CatC struct {
	c Cat
}

func (c *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

func (c *CatC) Eat() {
	c.c.Eat()
}

func main() {
	cat := new(CatB)
	cat.Eat()
	cat.Sleep()

	cat1 := new(CatC)
	cat1.Eat()
	cat1.Sleep()
}
