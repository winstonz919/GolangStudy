/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 14:07:37
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-09 15:58:31
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/5-代理模式/main.go
 * @Description: 代理模式
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// ---- 抽象主题 -----
type Subject interface {
	Do() string
}

// ---- 真实主题 -----
type RealSubject struct{}

func (rs *RealSubject) Do() string {
	return "real"
}

// ---- 代理 ----
type Proxy struct {
	real RealSubject // 组合真实主题
}

// 代理逻辑
func (p *Proxy) Do() string {
	pre, sub := "前置逻辑", "后置逻辑"
	// 代理模式用于对真实主题进行增强，比如类型判断、缓存操作等
	return pre + p.real.Do() + sub
}

func main() {
	p := &Proxy{}
	res := p.Do()
	fmt.Println(res)
}
