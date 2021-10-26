package main
func containsDuplicate(nums []int) bool {
	var temp int
	for i:=0;i<len(nums);i++{
		for j:=i;j<len(nums);j++{
			if(nums[i]>nums[j]){
				temp=nums[i]
				nums[i]=nums[j]
				nums[j]=temp
			}
		}
	}
	for i:=0;i<len(nums)-1;i++{
		if nums[i]==nums[i+1]{
			return true
		}
	}
	return false

}