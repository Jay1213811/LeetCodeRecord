package main

import (
	"fmt"
)

func specialArray(nums []int) int {
	res := -1
	for i := 1; i <= len(nums); i++ {
		var cnt int
		for j := range nums {
			if nums[j] >= i {
				cnt += 1
			}
		}
		if cnt == i {
			res = i
			break
		}
	}
	return res
}
func main() {
	nums:=[]int{3,6,7,7,0}
	fmt.Println(specialArray(nums))
}