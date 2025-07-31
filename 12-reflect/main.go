/*
 * @Author: WinstonRD
 * @Date: 2025-06-18 11:11:41
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-19 11:47:01
 * @FilePath: /GolangStudy/12-reflect/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user *User) Call() {
	fmt.Println("user is called..")
	fmt.Printf("%v\n", user)
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("type of input is", inputType)

	inputValue := reflect.ValueOf(input)
	inputFieldNum := inputType.NumField()
	fmt.Println("value of input is", inputValue)
	fmt.Println("num of input fields is", inputFieldNum)
	for i := 0; i < inputType.NumField(); i++ {
		fieldName := inputType.Field(i)
		fieldValue := inputValue.Field(i).Interface()
		fmt.Println("name:", fieldName, "value:", fieldValue)
	}
}

func main() {
	u := User{1, "recker", 29}
	DoFiledAndMethod(u)
}
