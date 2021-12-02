package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	clostNum:=nums[0]+nums[1]+nums[2]
	for i:=0;i<len(nums)-2;i++{
		l,r:=i+1,len(nums)-1
		for l<r{
			threeSum:=nums[l]+nums[r]+nums[i]
			if abs(threeSum-target)<abs(clostNum-target){
				clostNum=threeSum
			}
			if threeSum>target{
				r--
			}else if threeSum<target{
				l++
			}else{
				return target
			}
		}
	}
	return clostNum
}
func abs(a int)int{
	if a<0{
		return -a
	}
	return a
}