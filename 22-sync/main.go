/*
 * @Author: WinstonRD
 * @Date: 2025-06-25 15:54:54
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-25 23:24:40
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/22-sync/main.go
 * @Description: 并发安全
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 两个协程同时写，有时会引发concurrent map writes
func readMap() {
	m := make(map[string]int)
	wg.Add(2)
	go func() {
		m["a"] = 0
		wg.Done()
	}()

	go func() {
		m["a"] = 1
		wg.Done()
	}()
}

// 为保证线程安全，为map加入读写锁
type RWMap struct {
	sync.RWMutex
	m map[string]int
}

// map在使用前需要初始化
func NewRWMap() *RWMap {
	m := &RWMap{}
	m.m = make(map[string]int)
	return m
}

func (m *RWMap) Get(key string) (int, bool) {
	m.RLock()         // 读锁，读锁可以允许多个协程读，但是不可写
	defer m.RUnlock() // 延迟释放锁
	v, exists := m.m[key]
	return v, exists
}

func (m *RWMap) Set(key string, value int) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value
}

func readMapWithRWMutex(m *RWMap) {
	wg.Add(2)
	go func() {
		// time.Sleep(100 * time.Millisecond)
		m.Set("a", 1)
		wg.Done()
	}()

	go func() {
		// time.Sleep(100 * time.Millisecond) // 加负载测试一下延迟写入
		m.Set("a", 2)
		wg.Done()
	}()
}

func main() {
	// readMap()
	m := NewRWMap()
	readMapWithRWMutex(m)
	wg.Wait()
	fmt.Println(m.Get("a"))
}
