/*
 * @Author: WinstonRD
 * @Date: 2025-07-07 21:43:52
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-07 22:07:17
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/2-var/revertStr/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
)

var s = "abcdefghijklmn你"

func revertStr(str string) string {
	res := []rune(str)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}

func main() {
	ss := revertStr(s)
	fmt.Println(ss)
}
