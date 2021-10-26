package main

import "sort"

func maxProductDifference(nums []int) int {
	/*
	   找出(a*b)-(c*d)的最大值
	   那么就必须让a*b最大 c*d最小
	   先排序，取最大两个值相乘减最小两个数
	*/
	sort.Ints(nums)
	return  (nums[0]*nums[1])-(nums[len(nums)-1]*nums[len(nums)-2])

}