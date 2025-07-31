package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
 * @Author: WinstonRD
 * @Date: 2025-07-02 21:03:47
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-02 21:04:53
 * @FilePath: /27-gin/Users/winston/go/src/GolangStudy/14-goroutine/test4/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */

func main() {
	letter, number := make(chan bool), make(chan bool)

	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)
	number <- true
	wait.Wait()
}
