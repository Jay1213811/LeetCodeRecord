package main

import (
	"fmt"
	"sort"
)

func rearrangeArray(nums []int) []int {
	//从小到大排序,小大小依次轮流放入新数组
	sort.Ints(nums)
	res:=make([]int,0)
	start,end:=0,len(nums)-1
	for start<end{
		res=append(res,nums[start])
		res=append(res,nums[end])
		start++
		end--
	}
	if start==end{
		res=append(res,nums[start])
	}
	return res

}
func main()  {
	tmp:=[]int{3,4,0,2,5}
	fmt.Println(rearrangeArray(tmp))
}