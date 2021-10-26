package main

import (
	"sort"
)

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方
//组成的新数组，要求也按 非递减顺序 排序。
/*
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func sortedSquares(nums []int) []int {
	result:=make([]int,0)
	for i:=0;i<len(nums);i++{
		result=append(result,nums[i]*nums[i])
	}
	sort.Ints(result)
	return result

}
