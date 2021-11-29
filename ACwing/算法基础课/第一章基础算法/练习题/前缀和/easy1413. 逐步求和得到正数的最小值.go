package main

import "sort"

//给你一个整数数组 nums。你可以选定任意的正数 startValue 作为初始值。
//你需要从左到右遍历 nums数组，并将 startValue 依次累加上nums数组中的值。
//请你在确保累加和始终大于等于 1 的前提下，选出一个最小的正数作为 startValue 。
//
/*
输入：nums = [-3,2,-3,4,2]
输出：5
解释：如果你选择 startValue = 4，在第三次累加时，和小于 1 。
*/

//思路:
//首先通过前缀和计算出前缀和为
//prefixSum:=0 -3 -1 -4 0 2
//对前缀和排序找到最小的值
//如果前缀和最小的都大于1返回1即可
//否则返回1-prefisSum[0]
func minStartValue(nums []int) int {
	prefixSum:=[]int{}
	prefixSum[0]=0
	for i:=0;i<len(nums)-1;i++{
		prefixSum[i+1]=prefixSum[i]+nums[i]
	}
	sort.Ints(prefixSum)
	if prefixSum[0]>=1{
		return 1
	}
	return 1-prefixSum[0]
}