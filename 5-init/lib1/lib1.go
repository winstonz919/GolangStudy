/*
 * @Author: WinstonRD
 * @Date: 2025-03-25 23:43:26
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-03-25 23:44:44
 * @FilePath: /GolangStudy/5-init/lib1/lib1.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package lib1

import "fmt"

func Lib1Test() {
	fmt.Println("lib1 test")
}

func init() {
	fmt.Println("lib1 init")
}
