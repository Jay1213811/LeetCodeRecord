package main

import "sort"
//类似于三数之和
//三数之和通过双指针将O(n^2)降低到O(n)
//三数之和通过固定一个数，四数之和通过固定两个数将O(n^3)降低到O(n^2)
func fourSum(nums []int, target int) [][]int {
	res:=[][]int{}
	if nums==nil || len(nums)<4{
		return res
	}
	sort.Ints(nums)
	for i:=0;i<len(nums)-1;i++{
		//去重
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		for j:=i+1;j<len(nums)-2;j++{
			//去重
			if j > i+1 && nums[j] == nums[j-1]{
				continue
			}
			now:=nums[i]+nums[j]
			left,right:=j+1,len(nums)-1
			for left<right{
				sum:=now+nums[left]+nums[right]
				if sum==target{
					ans:=[]int{nums[i],nums[j],nums[left],nums[right]}
					res=append(res,ans)
					//去重
					for left<right && nums[left]==nums[left+1]{
						left++
					}
					//去重
					for left<right && nums[right]==nums[right-1]{
						right--
					}
					left++
					right--
				}else if sum>target{
					right--
				}else{
					left++
				}
			}
		}
	}
	return res
}



