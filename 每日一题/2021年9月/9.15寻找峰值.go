package main
func findPeakElement(nums []int) int {
	if len(nums)==1{
		return 0
	}else if len(nums)==2{
		if nums[0]>nums[1]{
			return 0
		}else{
			return 1
		}
	}
	slow,fast:=0,1
	for fast<len(nums){
		if nums[slow]<nums[fast]{
			slow++
			fast++
		}else{
			return slow
		}
	}
	return fast-1
}