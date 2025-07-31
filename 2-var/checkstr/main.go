/*
 * @Author: WinstonRD
 * @Date: 2025-07-05 11:49:13
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-05 11:59:48
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/2-var/checkstr/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */

/*
请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允 许使⽤额外的存
储结构】。 给定⼀个string，请返回⼀个bool值,true代表所有字符全都 不同，false代表存在相同的字
符。 保证字符串中的字符为【ASCII字符】。字符串的⻓ 度⼩于等于【3000】
*/
package main

import (
	"fmt"
	"strings"
)

func check(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	for _, v := range s {
		fmt.Println(v)
	}

	return true
}

func main() {
	s := "hello world!你好世界！"
	rt := check(s)
	fmt.Println(rt)
}
