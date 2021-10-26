package main

import "fmt"

func buildArray(nums []int) []int {
	ans:=make([]int,0)
	for i:=0;i<len(nums);i++{
		ans=append(ans,nums[nums[i]])
	}
	return ans
}
func main()  {
	nums:=[]int{0,2,1,5,3,4}
	fmt.Println(buildArray(nums))
}