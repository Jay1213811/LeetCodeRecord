package main

import "fmt"

func maximumDifference(nums []int) int {
	max:=0
	flag:=false
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[j]-nums[i]>max{
				max=nums[j]-nums[i]
				flag=true
			}
		}
	}
	if flag==false{
		return -1
	}
	return max
}
func main()  {
	nums:=[]int{7,1,5,4}
	fmt.Println(maximumDifference(nums))
}