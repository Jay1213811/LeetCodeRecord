package main

import (
	"fmt"
)

func countKDifference(nums []int, k int) int {
	count:=0

	for i:=0;i<len(nums)-1;i++{
		for j:=i;j<len(nums);j++{
			if abs(nums[j]-nums[i])==k{
				if (nums[i]==k && nums[j]==0) || (nums[j]==0 && nums[i]==k){
					continue
				}
				count++
			}
		}

	}
	return count
}
func abs(a int)int{
	if a>=0{
		return a
	}else{
		return -a
	}
}
func main(){

	nums:=[]int{1,2,2,1}
	fmt.Println(countKDifference(nums,1))
}