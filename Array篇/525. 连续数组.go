package main

import "fmt"

func findMaxLength(nums []int) int {
	cntIdx := map[int]int{
		0: -1,
	}

	cnt, maxLen := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt += -1
		} else {
			cnt += 1
		}
		if v, ok := cntIdx[cnt]; ok {
			maxLen = max(maxLen, i - v)
		} else {
			cntIdx[cnt] = i
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main()  {
	var s =[]int{0,1,0}
	fmt.Println(findMaxLength(s))
}