package main

import (
	"fmt"

)

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

*/
func maxSubArray(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	res:=nums[0]

	for i:=1;i<len(nums);i++{
		nums[i]=max(nums[i],nums[i]+nums[i-1])
		res=max(nums[i],res)
	}

	return res
}

func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}
func main(){
	var a=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(a))
}