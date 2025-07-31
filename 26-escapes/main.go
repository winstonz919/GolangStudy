/*
 * @Author: WinstonRD
 * @Date: 2025-06-27 00:34:59
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-27 01:43:30
 * @FilePath: /scripts/Users/winston/go/src/GolangStudy/26-escapes/main.go
 * @Description: 内存逃逸常见的几种情况
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"unsafe"
)

var ch = make(chan *int)
var c = make(chan int)

// 返回指针
func f1() *int {
	var i int
	return &i
}

// 向channel中写入指针
func f2() {
	var i int
	ch <- &i
}

// 闭包
func f3() func() int {
	var i int
	return func() int {
		return i
	}
}

type animal interface {
	run()
}

type dog struct {
}

func (d *dog) run() {

}

// 接口持有对象，导致类型不确定
func f4() {
	var a animal
	a = &dog{}
	a.run()
	fmt.Println(a)
}

func f5() {
	m := make(map[string]*int)
	i := 1
	m["a"] = &i
}

func f6() {
	d := &dog{}
	d.run()
}

// println函数接收一个空接口作为参数，也会发生逃逸
// 在生产环境中应去掉打印
func f7() {
	var i int
	fmt.Println(i)
}

// 栈空间不足，当切片占用内存超过一定大小或者无法确定长度时，会分配在堆上
func f8() {
	nums := make([]int, 8193) // 超过64KB，发生逃逸
	fmt.Println(unsafe.Sizeof(nums), len(nums), cap(nums))
}

func main() {
	f8()
}

/*
go程序会在2个地方为变量分配内存空间，一个是全局的堆（heap）用来动态分配内存，另一个是每个goroutine的栈（stack）空间。
go语言实现垃圾回收（GC）机制，go语言的内存管理是自动的，通常开发者不需要关心内存分配在栈上还是堆上，但是从性能角度出发，在栈上分配内存和在堆上分配内存性能差异是很大的。
在函数中申请一个对象，如果分配在栈中那么函数结束后自动回收，如果分配在堆中，则在函数结束后某个时间点进行垃圾回收。
在栈上分配内存和回收内存的开销很低，只需要2个cpu指令，PUSH和POP，消耗的仅仅是将数据拷贝到内存的时间，内存IO通常能达到30GB/s，效率非常。
在堆上分配内存，一个很大的开销是垃圾回收，go语言使用的是标记清除算法，并在此基础上使用了三色标级法和写屏障技术，提高了效率。
*/

// 一般情况下，对于需要修改原对象值，或占用内存比较大的结构体，选择传指针。对于只读的占用内存较小的结构体，直接传值能够获得更好的性能。
