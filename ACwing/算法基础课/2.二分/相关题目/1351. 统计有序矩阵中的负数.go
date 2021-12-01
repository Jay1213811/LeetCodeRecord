package main
/*
@source:https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/
给你一个 m * n 的矩阵 grid，矩阵中的元素无论是按行还是按列，都以非递增顺序排列。

请你统计并返回 grid 中 负数 的数目。
输入：grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
输出：8
解释：矩阵中共有 8 个负数。
*/

import "fmt"

func countNegatives(grid [][]int) int {
	ans:=0
	//逐行二分
	for i:=0;i<len(grid);i++{
		ans+=countNum(grid[i])
	}
	return ans
}
func countNum(nums []int)int{
	l,r:=0,len(nums)-1
	//做一个边界处理
	if nums[len(nums)-1]>=0{
		return 0
	}
	for l<r{
		mid:=(l+(r-l)>>1)
		if nums[mid]>=0{
			l=mid+1
		}else{
			r=mid
		}
	}
	return len(nums)-l
}
func main(){
	nums:=[]int{3,2}
	fmt.Println(countNum(nums))
}
