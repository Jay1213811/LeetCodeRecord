package 第_269_场周赛

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	res:=[]int{}
	for i:=0;i<len(nums);i++{
		if nums[i]<target {
			continue
		}else if nums[i]==target{
			res=append(res,i)
		}else{
			return res
		}
	}
	return res
}