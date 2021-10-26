package main

import "fmt"

/*

*/
func totalHammingDistance(nums []int) int {
	result:=0
	for i:=0;i<31;i++{
		count:=[2]int{}
		for t:=0;t<len(nums);t++{
			count[nums[t]&1]++
			nums[t]>>=1
		}
		result+=count[0]*count[1]
	}
	return result
}
func main()  {
	a:=[]int{4,14,2}
	fmt.Println(totalHammingDistance(a))
}