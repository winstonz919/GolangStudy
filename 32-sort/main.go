/*
 * @Author: WinstonRD
 * @Date: 2025-07-14 13:01:01
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-14 22:47:32
 * @FilePath: /GolangStudy/32-sort/1-bubble/main.go
 * @Description: 冒泡排序
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sort"
)

type NumList []int

// 冒泡排序
func BubbleSort(nums NumList) NumList {
	for k, n := range nums {
		for kk, m := range nums {
			if n < m {
				nums[k], nums[kk] = nums[kk], nums[k]
			}
		}
	}
	return nums
}

// 选择排序
func ChooseSort(nums NumList) NumList { return nums }

// 插入排序
func InsertSort(nums NumList) NumList { return nums }

// 希尔排序
func HillSort(nums NumList) NumList { return nums }

// 快速排序
func FastSort(nums NumList) NumList { return nums }

// 归并排序
func MergeSort(nums NumList) NumList { return nums }

// 堆排序
// 本质是完全二叉树，根节点为最大值或者最小值
func HeapSort(nums NumList) NumList { return nums }

// 桶排序
// 就是把数值按照范围进行划分，把数值依次放入一个个划分的范围内，称之为 桶，然后在桶内进行排序，然后依次输出每个桶的值
func BucketSort(nums NumList) NumList {
	// 确认最大和最小值
	min, max := 10000, 0
	for _, v := range nums {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}
	// 初始化桶
	bucketNum := (max-min)/len(nums) + 1
	var bucketList [][]int
	for i := 0; i < bucketNum; i++ {
		bucketItem := make([]int, 0)
		bucketList = append(bucketList, bucketItem)
	}
	// 把元素装入桶
	for _, v := range nums {
		number := (v - min) / len(nums)
		bucketList[number] = append(bucketList[number], v)
	}
	// 桶内排序
	for _, v := range bucketList {
		sort.Ints(v)
	}
	// 依次输出
	for _, v := range bucketList {
		for _, value := range v {
			fmt.Println(value)
		}
	}
	return nums
}

func main() {
	array := []int{5, 16, 95, 1, 45, 66, 75, 20, 62, 38, 76, 94, 43}
	BucketSort(array)
}
