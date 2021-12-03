package main
/*
题目要求O(1)空间复杂度,那么必须在原数组基础上修改
slow,fast快慢指针
fast不断往前移动.slow在没有出现目标值的时候移动
*/

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
