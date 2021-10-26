package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var slow int=0
	if(len(nums)==0){
		return 0
	}
	for fast:=1;fast<len(nums);fast++{
		if nums[slow]!=nums[fast]{
			nums[slow+1]=nums[fast]
			slow++
		}
	}
	return slow+1
}
func main()  {
	var  test = []int {1,1,2,3}

	fmt.Println(removeDuplicates(test))

}