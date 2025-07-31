/*
 * @Author: WinstonRD
 * @Date: 2025-06-17 20:44:43
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-18 09:12:15
 * @FilePath: /GolangStudy/10-struct/test2/main.go
 * @Description: 继承
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

type Man struct {
	name string
	sex  string
}

func (m *Man) SayHi() {
	fmt.Println("hi, I'm", m.name)
}

type SuperMan struct {
	Man
	level int
}

func main() {
	m := Man{name: "winston", sex: "male"}
	m.SayHi()
	var sm SuperMan
	sm.name = "Liming"
	sm.sex = "male"
	sm.level = 10
	sm.SayHi()
}
