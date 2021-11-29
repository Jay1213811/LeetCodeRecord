package main

import (
	"fmt"
)

/*
有一个自行车手打算进行一场公路骑行，这条路线总共由n + 1个不同海拔的点组成。自行车手从海拔为 0的点0开始骑行。
给你一个长度为 n的整数数组gain，其中 gain[i]点 i和点 i + 1的 净海拔高度差（0 <= i < n）。
请你返回 最高点的海拔 。
*/
//输入：gain = [-5,1,5,0,-7]
//输出：1
//解释：海拔高度依次为 [0,-5,-4,1,1,-6] 。最高海拔为 1 。
/*
思路:
典型前缀和问题:
首先对数组排序
然后求得前缀和,返回prefixSum的最大值
*/
func largestAltitude(gain []int) int {
	prefixSum:=make([]int,len(gain)+1)
	prefixSum[0]=0
	for i:=0;i<len(gain);i++{
		prefixSum[i+1]=prefixSum[i]+gain[i]
	}

	QuicklySort(prefixSum,0,len(prefixSum)-1)
	return prefixSum[len(prefixSum)-1]
}
func QuicklySort(nums []int,left,right int){
	if left>=right{
		return
	}
	l,r,base:=left-1,right+1,nums[left+(right-left)>>1]
	for l<r{
		l++
		r--
		for nums[l]<base{
			l++
		}
		for nums[r]>base{
			r--
		}
		if l<r{
			nums[l],nums[r]=nums[r],nums[l]
		}
	}
	QuicklySort(nums,left,r)
	QuicklySort(nums,r+1,right)

}


func main()  {
	fmt.Println(largestAltitude([]int{-4,-3,-2,-1,4,3,2}))
}