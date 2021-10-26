package main

import "fmt"

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/
//动态规划法

func lengthOfLIS(nums []int) int {
	//
	// dp[i]=max(dp[i],dp[j]+1)
	var res int =0
	dp:=make([]int,len(nums))
	if len(nums)<=1{
		return len(nums)
	}
	for i:=0;i<len(nums);i++{
		dp[i]=1
		for j:=0;j<i;j++{
			if nums[i]>nums[j]{
				dp[i]=max(dp[i],dp[j]+1)
			}
		}
		res=max(res, dp[i])
	}
	return res

}

func main()  {
	var test =[]int{0,1,0,3,2,3}
	fmt.Println(lengthOfLIS(test))
}