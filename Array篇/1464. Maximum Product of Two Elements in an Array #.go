package main
func maxProduct(nums []int) int {
	for i:=0;i<len(nums);i++{
		for j:=1;j<len(nums);j++{
			if(nums[i]<nums[j]){
				nums[i],nums[j]=nums[j],nums[i]
			}
		}
	}
	return (nums[0]-1)*(nums[1]-1)
}