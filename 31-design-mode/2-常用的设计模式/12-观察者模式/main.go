/*
 * @Author: WinstonRD
 * @Date: 2025-07-10 14:43:14
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-10 16:15:56
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/12-观察者模式/main.go
 * @Description: 观察者模式
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

type Event struct {
	Noti    Notifier
	One     Listener
	Another Listener
	Msg     string
}

func NewEvent(Noti Notifier, One Listener, Another Listener, Msg string) *Event {
	return &Event{Noti: Noti, One: One, Another: Another, Msg: Msg}
}

type Listener interface {
	OnFriendBeFight(*Event)
	GetTitle() string
	GetName() string
	GetParty() string
	Fight(Listener, Notifier)
}

type Notifier interface {
	AddListener(listener ...Listener)
	RemoveListener(listener Listener)
	Notify(event *Event)
}

type NotifierImpl struct {
	heroList []Listener
}

func NewNotifier() Notifier {
	return &NotifierImpl{}
}

func (n *NotifierImpl) AddListener(listeners ...Listener) {
	n.heroList = append(n.heroList, listeners...)
}

func (n *NotifierImpl) RemoveListener(listener Listener) {
	for index, l := range n.heroList {
		if l == listener {
			n.heroList = append(n.heroList[:index], n.heroList[index+1:]...)
		}
	}
}

func (n *NotifierImpl) Notify(event *Event) {
	fmt.Println("[世界频道] 百晓生广播消息：", event.Msg)

	// 通知全部listener
	for _, listener := range n.heroList {
		listener.OnFriendBeFight(event)
	}
}

type Hero struct {
	Name  string
	Party string
}

func NewHero(name, party string) Listener {
	return &Hero{Name: name, Party: party}
}

func (h *Hero) OnFriendBeFight(event *Event) {
	// 判断是否为当事人，忽略消息
	if h.Name == event.One.GetName() {
		return
	}
	// 本帮派的同伴挨揍了，报仇
	if h.Party == event.Another.GetParty() {
		fmt.Println(h.GetTitle(), "得知了消息，发起报仇反击！")
		// h.Fight(event.One, event.Noti)
		return
	}
	// 非本帮派的人挨揍了，拍手
	if h.Party == event.One.GetParty() {
		fmt.Println(h.GetTitle(), "得知了消息，拍手叫好！")
	}
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) GetParty() string {
	return h.Party
}

func (h *Hero) GetTitle() string {
	return h.Party + "的" + h.Name
}

func (h *Hero) Fight(another Listener, baixiao Notifier) {
	// 生成一个武林事件
	msg := fmt.Sprintf("%s 将 %s 揍了", h.GetTitle(), another.GetTitle())
	event := NewEvent(baixiao, h, another, msg)
	// 发消息给百晓生
	baixiao.Notify(event)
}

// 丐帮：黄蓉、洪七公、乔峰
// 明教：张无忌、灭绝师太、金毛狮王
// 广播者：百晓生
func main() {
	huangRong := NewHero("黄蓉", "丐帮")
	hongQiGong := NewHero("洪七公", "丐帮")
	qiaoFeng := NewHero("乔峰", "丐帮")
	zhangWuJi := NewHero("张无忌", "明教")
	mieJue := NewHero("灭绝师太", "明教")
	xieXun := NewHero("金毛狮王", "明教")

	baiXiao := NewNotifier()
	baiXiao.AddListener(huangRong, hongQiGong, qiaoFeng, zhangWuJi, mieJue, xieXun)

	huangRong.Fight(zhangWuJi, baiXiao)
}
