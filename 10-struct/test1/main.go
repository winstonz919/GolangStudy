/*
 * @Author: WinstonRD
 * @Date: 2025-06-17 13:44:23
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-17 20:32:54
 * @FilePath: /GolangStudy/10-struct/test1/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) SetName(name string) {
	h.Name = name
}

func (h *Hero) Show() {
	fmt.Printf("Name: %s, Ad: %d, Level: %d\n", h.Name, h.Ad, h.Level)
}

func main() {
	hero := Hero{
		Name:  "winston",
		Ad:    100,
		Level: 1,
	}
	fmt.Println(hero.GetName())

	hero.SetName("recker")
	fmt.Println(hero.GetName())
	hero.Show()
}
