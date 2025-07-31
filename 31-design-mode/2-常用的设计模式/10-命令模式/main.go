/*
 * @Author: WinstonRD
 * @Date: 2025-07-10 10:18:16
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-10 11:00:42
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/10-命令模式/main.go
 * @Description: 命令模式, 通过中间调用者，调度执行者
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

/* --------- 治疗单抽象 -------- */
type Commend interface {
	Treat()
}

/* --------- 治疗单实现 -------- */
type CommandTreatEye struct {
	doctor *Doctor
}

func (cte *CommandTreatEye) Treat() {
	cte.doctor.TreatEye()
}

type CommandTreatNose struct {
	doctor *Doctor
}

func (ctn *CommandTreatNose) Treat() {
	ctn.doctor.TreatNose()
}

type Doctor struct{}

func (d *Doctor) TreatEye() {
	fmt.Println("医生治疗眼睛")
}

func (dl *Doctor) TreatNose() {
	fmt.Println("医生治疗鼻子")
}

type Nurse struct {
	CommandList []Commend // 治疗单列表
}

func (n *Nurse) Notify() {
	if len(n.CommandList) == 0 {
		return
	}

	for _, cmd := range n.CommandList {
		cmd.Treat()
	}
}

func main() {
	doctor := new(Doctor)
	// 填写病单，病单依赖医生
	cmdEye := &CommandTreatEye{doctor: doctor}
	cmdNose := &CommandTreatNose{doctor: doctor}

	// 收集病单
	nurse := new(Nurse)
	nurse.CommandList = append(nurse.CommandList, cmdEye, cmdNose)
	// 下发执行
	nurse.Notify()
}
