package main
func moveZeroes(nums []int)  {
	 var i int=0
	 var temp int
	 for j:=0;j<len(nums);j++{
	 	if(nums[j]!=0){
	 		temp=nums[j]
	 		nums[j]=nums[i]
	 		nums[i]=temp
	 		i++
		}
	 }
}