package main

import (
	"fmt"
	"sort"
)

func reductionOperations(nums []int) int {
	sort.Ints(nums)
	var res,cnt int=0,0 //res代表总操作数 cnt代表每个元素的操作数
	for i:=1;i<len(nums);i++{
		if nums[i]!=nums[i-1]{
			cnt++
		}
		res+=cnt
	}
	return res
}


func main()  {
	a:=[]int{1,1,2,3,4}
	fmt.Println(reductionOperations(a))
}