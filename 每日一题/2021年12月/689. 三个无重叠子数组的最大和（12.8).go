package main

import "fmt"
//source:https://leetcode-cn.com/problems/maximum-sum-of-3-non-overlapping-subarrays/solution/san-ge-wu-zhong-die-zi-shu-zu-de-zui-da-4a8lb/
//如何得到一个长度为k的子数组最大值
//滑动窗口法:每次控制长度为k-1个数之和。如果新加的数大于删除第一个，那么第一个数不要了
func maxSumOfOneSubarray(nums []int, k int) (ans []int) {
		sum:=0
		maxSum:=0
		var index int
		for i:=0;i<len(nums);i++{
			sum+=nums[i]
			if i>=k-1{
				if sum>maxSum{
					maxSum=sum
					index=i-k+1
				}
				sum-=nums[i-k+1]
			}
		}
		ans=append(ans,index)
		return ans
}
//如何得到两个长度为k的子数组最大值
//
func maxSumOfTwoSubarrays(nums []int, k int) (ans []int) {
	sum1,maxSum1,maxSum1Idx:=0,0,0
	sum2,maxSum2:=0,0
	for i:=k;i<len(nums);i++{
		sum1+=nums[i-k]
		sum2+=nums[i]
		if i>=2*k-1{
			if sum1>maxSum1{
				maxSum1=sum1
				maxSum1Idx=i-2*k+1
			}
			if maxSum1+sum2>maxSum2{
				maxSum2=maxSum1+sum2
				ans=[]int{maxSum1Idx,i-k+1}
			}
			sum1-=nums[i-2*k+1]
			sum2-=nums[i-k+1]
		}
	}
	return ans
}

func main(){
	nums:=[]int{1,2,1,2,6,7,5,1}
	k:=2
	fmt.Println(maxSumOfOneSubarray(nums,k))
}