/*
 * @Author: WinstonRD
 * @Date: 2025-07-07 23:44:45
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-07 23:48:08
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/9-map/structMap/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

type A struct {
	Name  string
	Param map[string]any
}

func main() {
	a := A{}
	a.Name = "aaa"
	// a.Param["aaa"] = 1 // map需要初始化
	a.Param = make(map[string]interface{})
	a.Param["aaa"] = 1
	a.Param["bbb"] = "bbbbbb"
	fmt.Println(a.Name, a.Param)
}
