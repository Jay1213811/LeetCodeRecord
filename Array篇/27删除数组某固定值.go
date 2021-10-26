package main
func removeElement(nums []int, val int) int {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}