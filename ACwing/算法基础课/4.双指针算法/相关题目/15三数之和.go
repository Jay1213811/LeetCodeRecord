package main

import "sort"


// 排序后固定一个数nums[i]



//相较于16题 难点在去重这一步
// 1. 固定一个数nums[i]若是>0 那么和必然不等于0 跳过
// 2. 若是nums[i] == nums[i+1], 则去重, 跳过
// 2. 若是nums[i] < 0,则向后生成两个指针left和right
// 若是nums[left] == nums[left+1] 则left++
// 若是nums[right] == nums[right-1] 则right--
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	if nums == nil || len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	// 此层for循环控制nums[i]的定位一直循环到nums[i] == 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res := []int{nums[i], nums[left], nums[right]}
				ans = append(ans, res)
				// 需要将left < right 写在 && 的左边 先行判断, 因为&&有短路求值的特性
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right  && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}else if sum > 0 {
				right--
			}else if sum < 0 {
				left++
			}
		}
	}
	return ans
}

