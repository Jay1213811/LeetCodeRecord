package main

import "math"

func threeSumClosest(nums []int, target int) int {
	result, difference := 0, math.MaxInt8

	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			for k:=j+1;k<len(nums);k++{
				if(int(math.Abs(float64(nums[i]+nums[j]+nums[k]-target)))<difference){
					difference= int(math.Abs(float64(nums[i] + nums[j] + nums[k] - target)))
					result=nums[i]+nums[j]+nums[k]
				}
			}
		}
	}
	return result
}
