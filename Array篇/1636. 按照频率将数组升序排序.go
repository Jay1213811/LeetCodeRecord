package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) map[int]int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[j] < nums[i]
		}
		return freq[nums[i]] < freq[nums[j]]
	})


	return freq
}
func main()  {
	var nums=[]int{-1,1,-6,4,5,-6,1,4,1}
	fmt.Println(frequencySort(nums))

}