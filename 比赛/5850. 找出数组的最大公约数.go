package main

import "sort"

func findGCD(nums []int) int {
	sort.Ints(nums)
	return gcdx(nums[0],nums[len(nums)-1])
}

func gcdx(x, y int) int {
	var tmp int
	for {
		tmp = (x % y)
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}
