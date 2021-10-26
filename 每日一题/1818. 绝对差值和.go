package main

import "sort"

/*
排序 + 二分查找法
*/
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	res, max := 0, 0
	maxIndex := make(map[int]int, 0)
	for i := 0; i < len(nums1); i++ {
		temp := abs(nums1[i] - nums2[i])
		res += temp % 1000000007
		if temp >= max {
			if temp > max {
				maxIndex = make(map[int]int, 0)
			}
			maxIndex[i] = 1
			max = temp
		}
	}
	if res == 0 {
		return res
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	min, cha := max, 0
	for key, _ := range maxIndex {
		target, result := nums2[key], 0
		left, right := 0, len(nums1)
		if target <= nums1[left] {
			result = 0
		} else if target >= nums1[right - 1] {
			result = right - 1
		} else {
			for left < right {
				mid := (left + right + 1) / 2
				if target < nums1[mid] {
					right = mid - 1
				} else {
					left = mid
				}
			}
			result = left
		}
		if abs(nums1[result] - nums2[key]) < min {
			min = abs(nums1[result] - nums2[key])
			cha = max - abs(nums1[result] - nums2[key])
		}
		if result + 1 < len(nums1) && abs(nums1[result + 1] - nums2[key]) < min {  //left下标的左边相邻数也有可能是最优解...
			min = abs(nums1[result + 1] - nums2[key])
			cha = max - abs(nums1[result + 1] - nums2[key])
		}
	}
	return (res - cha) % 1000000007
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

