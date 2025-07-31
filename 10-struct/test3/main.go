package main

import "fmt"

type Notify interface {
	SendMessage()
}

type WechatUser struct {
	name string
}

func (u *WechatUser) SendMessage() {
	fmt.Println("send message to", u.name)
}

type AdminUser struct {
	name string
}

func (ad *AdminUser) SendMessage() {
	fmt.Println("send message to", ad.name)
}

func SendMessage(n Notify) {
	n.SendMessage()
}

func main() {
	// u := &WechatUser{"weixin001"}
	// u.SendMessage()

	// ad := &AdminUser{"admin001"}
	// ad.SendMessage()
	user1 := &WechatUser{"weixin001"}

	SendMessage(user1)

	user2 := &AdminUser{"admin001"}
	SendMessage(user2)
}
