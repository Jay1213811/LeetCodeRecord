package main

import "sort"

//题目:给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
/*
输入: [2,2,3,4]
输出: 3
解释:
有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3
*/
//思路:
//1.使用DFS遍历得到所有的三元组,再统计满足三角形成立的三元组计算数量.超时
//2.采用整数二分.
//第一步:排序.第二步确定两条边.计算和.第三步通过二分查找满足条件的第三条边
/*

2   2   3   3   4
拿了2和2后.
i,j:=
*/

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res:=0
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			sum:=nums[i]+nums[j]
			//找到数组[j+1:len(nums)]最后一个满足小于sum的值下标为k
			//3 3 4 返回1 .满足条件的个数为k-j
			//
			//[j+1:j+k+1] j=1 k=0
			k:=getIndex(sum,nums[j+1:len(nums)])
			if k==-1{
				res+=0
			}else{

				res+=(k+1)
			}
		}
	}
	return res
}
//3 3 4
func getIndex(sum int,nums []int)int{
	left,right:=0,len(nums)-1
	for left<right{
		mid := (left+(right-left+1)>>1)
		if nums[mid]<sum{
			left=mid
		}else{
			right=mid-1
		}
	}
	if left<len(nums) && nums[left]<sum{
		return left
	}
	return -1
}