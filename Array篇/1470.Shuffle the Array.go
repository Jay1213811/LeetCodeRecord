package main

import "fmt"

func shuffle(nums []int, n int) []int {
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[n+i])
	}
	return result

}
func main()  {
	var nums=[]int{2,5,1,3,4,7}
	var k int=3
	fmt.Println(shuffle(nums,k))
}