/*
 * @Author: WinstonRD
 * @Date: 2025-07-08 10:03:46
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-08 10:43:39
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/31-design-mode/main.go
 * @Description: 开闭原则，对扩展开放，对修改关闭
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

// 用接口抽象，在对象中具体实现，避免增加业务需求导致修改源代码
type Notice interface {
	SendNotice()
}

// 具体实现
type MessageNotice struct{}

func (m *MessageNotice) SendNotice() {
	fmt.Println("发送短信")
}

// 具体实现
type EmailNotice struct{}

func (e *EmailNotice) SendNotice() {
	fmt.Println("发送邮件")
}

func sendNotice(n Notice) {
	n.SendNotice()
}

func main() {
	// 用接口接收对象，避免使用具体对象
	var n Notice
	n = &MessageNotice{}
	sendNotice(n) // 复用接口方法
	n = &EmailNotice{}
	sendNotice(n)
}
