/*
 * @Author: WinstonRD
 * @Date: 2025-07-10 14:03:32
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-10 14:28:26
 * @FilePath: /GolangStudy/31-design-mode/2-常用的设计模式/11-策略模式/main.go
 * @Description: 策略模式
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
)

type WeaponStrategy interface {
	UseWeapon()
}

type KnifeStrategy struct{}

func NewKnifeStrategy() WeaponStrategy {
	return &KnifeStrategy{}
}

func (k *KnifeStrategy) UseWeapon() {
	fmt.Println("用匕首攻击")
}

type AK47Strategy struct{}

func NewAK47Strategy() WeaponStrategy {
	return &AK47Strategy{}
}

func (a *AK47Strategy) UseWeapon() {
	fmt.Println("用AK47攻击")
}

type Hero struct {
	weaponStrategy WeaponStrategy
}

func NewHero() *Hero {
	return &Hero{}
}

func (h *Hero) SetWeaponStrategy(weaponStrategy WeaponStrategy) {
	h.weaponStrategy = weaponStrategy
}

func (h *Hero) Fight() {
	if h.weaponStrategy == nil {
		fmt.Println("请选择武器")
		return
	}
	h.weaponStrategy.UseWeapon()
}

func main() {
	hero := NewHero()
	hero.Fight()                              // 未选择策略，此时weaponstrategy只有类型没有具体的值，为nil pointer
	hero.SetWeaponStrategy(NewAK47Strategy()) // 设置策略
	hero.Fight()

	hero.SetWeaponStrategy(NewKnifeStrategy()) // 更换策略
	hero.Fight()

}
