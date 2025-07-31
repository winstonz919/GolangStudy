/*
 * @Author: WinstonRD
 * @Date: 2025-03-26 01:21:06
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-17 11:20:00
 * @FilePath: /GolangStudy/8-slice/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
)

func changeArray(arr [10]int) {
	arr[0] = 100
}

func changeSlice(arr []int) {
	arr[0] = 100
}

func main() {
	var arr [10]float64
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for k, v := range arr {
		fmt.Println(k, v)
	}

	// array
	var arr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr1)
	changeArray(arr1) // array是值传递
	fmt.Println(arr1)

	// slice
	// 动态长度，底层是数组
	var arr2 = []int{1, 2, 3, 4, 5}
	arr3 := append(arr2, 1)
	fmt.Println("arr3 is", arr3)
	fmt.Printf("type of arr2 is %T\n", arr2)
	fmt.Printf("type of arr3 is %T\n", arr3)
	changeSlice(arr2) // slice是引用传递
	fmt.Println(arr2)

	arr4 := make([]int, 3) // make函数创建slice
	fmt.Println(arr4)
	if arr4 == nil {
		fmt.Println("arr4 is nil")
	} else {
		fmt.Println("arr4 is not nil")
	}

	var arr5 []int
	if arr5 == nil {
		fmt.Println("arr5 is nil")
	} else {
		fmt.Println("arr5 is not nil")
	}

	// slice的cap和len
	// len是当前slice的长度, cap是当前slice的容量, cap是底层数组的长度, cap >= len
	var arr6 = make([]int, 3)
	fmt.Printf("arr6 is %v, len is %d, cap is %d\n", arr6, len(arr6), cap(arr6))
	// cap不够用, 会自动扩容, 扩容后cap是原来的2倍,如果cap还是不够用, 会继续扩容, 直到cap >= len
	arr6 = append(arr6, 1)
	fmt.Printf("arr6 is %v, len is %d, cap is %d\n", arr6, len(arr6), cap(arr6))

	fmt.Println("------")

	// slice的截取
	arr7 := make([]int, 3, 4)
	// fmt.Printf("arr7 id:%p\n", &arr7)
	fmt.Printf("arr7 is %v, len is %d, cap is %d\n", arr7, len(arr7), cap(arr7))
	arr8 := arr7[0:2]
	fmt.Printf("arr8 is %v, len is %d, cap is %d\n", arr8, len(arr8), cap(arr8))
	fmt.Println("------change arr7[0]--------")
	arr7[0] = 1
	fmt.Println(arr7, arr8)
	arr7 = append(arr7, 1)
	arr7[1] = 2
	fmt.Println(arr7, arr8)
	arr7 = append(arr7, 1)
	arr7[0] = 3
	fmt.Println(arr7, arr8)
	// slice的底层结构体是数组, 所以slice的cap是固定的, 当slice扩容后, 原来的slice会指向新的数组
	// 所以arr7和arr8指向的是不同的数组
	// type Slice struct {
	// 	array unsafe.Pointer // 指向底层数组的指针
	// 	len   int            // 当前slice的长度
	// 	cap   int            // 当前slice的容量
	// }
	// fmt.Printf("arr7 id:%p\n", &arr7)
	// fmt.Printf("arr7 is %v, len is %d, cap is %d\n", arr7, len(arr7), cap(arr7))
	// // arr8的值也会变, 因为arr8是arr7的引用
	// fmt.Printf("arr8 is %v, len is %d, cap is %d\n", arr8, len(arr8), cap(arr8))
	// fmt.Println("------append 1 to arr7------------")
	// arr7 = append(arr7, 1)
	// fmt.Printf("arr7 id:%p\n", &arr7)
	// fmt.Printf("arr7 is %v, len is %d, cap is %d\n", arr7, len(arr7), cap(arr7))
	// fmt.Printf("arr8 is %v, len is %d, cap is %d\n", arr8, len(arr8), cap(arr8))
	// fmt.Println("------append 2 to arr7------------")
	// arr7 = append(arr7, 2)
	// fmt.Printf("arr7 id:%p\n", &arr7)
	// fmt.Printf("arr7 is %v, len is %d, cap is %d\n", arr7, len(arr7), cap(arr7))
	// // arr8的cap还是4, 因为arr8是arr7的引用, arr7扩容后, arr8的cap不会变
	// fmt.Printf("arr8 is %v, len is %d, cap is %d\n", arr8, len(arr8), cap(arr8))
	// fmt.Println("------change arr7[1]--------")
	// arr7[1] = 2
	// fmt.Printf("arr7 is %v, len is %d, cap is %d\n", arr7, len(arr7), cap(arr7))
	// fmt.Printf("arr8 is %v, len is %d, cap is %d\n", arr8, len(arr8), cap(arr8))

}
