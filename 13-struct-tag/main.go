/*
 * @Author: WinstonRD
 * @Date: 2025-06-19 15:15:39
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-19 15:22:01
 * @FilePath: /GolangStudy/13-struct-tag/main.go
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
	Id   int    `info:"user_id" doc:"用户ID"`
	Name string `info:"user_name" doc:"用户名称"`
}

func findTag(ori interface{}) {
	t := reflect.TypeOf(ori)
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println(tagInfo, tagDoc)
	}
}

func main() {
	user := User{1, "recker"}
	findTag(user)
}
