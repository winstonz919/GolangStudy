/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 13:45:36
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-08 14:06:40
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/4-单例模式/main.go
 * @Description: 保证一个类永远只有一个对象，并且这个对象还能被系统的其它模块使用
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var once sync.Once

// 这个类应为私有类
type singlton struct{}

// // 只有一个对象且为私有
// var instance = new(singlton) // 饿汉模式，只要包被import，就会初始化对象

// // 对外提供调用获取对象的访问权限
// func GetInstance() *singlton {
// 	return instance
// }

var instance *singlton // 懒汉模式

// 懒汉模式，首次调用GetInstance时才会创建对象
func GetInstance() *singlton {
	// 懒汉模式需要并发安全保障
	once.Do(func() {
		instance = new(singlton)
	})
	return instance
}

func (s *singlton) Show() {
	fmt.Println("I'm singlton")
}

func main() {
	i := GetInstance() // 对外提供调用接口
	i.Show()
}
