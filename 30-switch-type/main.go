/*
 * @Author: WinstonRD
 * @Date: 2025-07-04 13:55:13
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-05 11:35:27
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/30-switch-type/main.go
 * @Description: 类型判断
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"reflect"
)

type MyStr string

func (m *MyStr) P() {
	fmt.Println("MyStr")
}

type Car struct {
	maxSpeed int
	price    float64
}

func (c *Car) Buy() {
	fmt.Println("to buy a car")
}

func printType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("type is string")
	case int:
		fmt.Println("type is int")
	case float64:
		fmt.Println("type is float64")
	default:
		fmt.Println("unknown type")
	}
}

func reflectType(input interface{}) {
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)
	fmt.Println(input, t.Kind(), t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
		fmt.Println(v.Field(i))
	}
}

func main() {
	var i int
	printType(i)
	var s string
	printType(s)
	s = "123"
	fmt.Println(s)
	s = "456"
	fmt.Println(s)

	c := Car{}
	reflectType(c)
}
