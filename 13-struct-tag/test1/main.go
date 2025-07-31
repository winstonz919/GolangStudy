/*
 * @Author: WinstonRD
 * @Date: 2025-06-19 15:27:15
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-19 15:43:29
 * @FilePath: /GolangStudy/13-struct-tag/test1/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func main() {
	u := User{1, "recker", "male"}
	j, err := json.Marshal(u)
	if err != nil {
		fmt.Println("json err")
	} else {
		fmt.Printf("%s\n", j)
	}

	mu := &User{}
	err = json.Unmarshal(j, mu)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mu)

}
