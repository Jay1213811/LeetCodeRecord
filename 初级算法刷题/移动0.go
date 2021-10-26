package main

import "fmt"

func moveZeroe(nums []int) []int {
	 j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[j] = nums[i]
		if i != j {
			nums[i] = 0
		}
		j++
	}
	return nums
}

func main() {
	var  test = []int {0,1,0,2,0,4}
	fmt.Println(moveZeroe(test))
}