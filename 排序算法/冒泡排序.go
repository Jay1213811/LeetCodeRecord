package main


import "fmt"

func bubbleSort(nums []int) []int {
	flag:=true
	for i:=0;i<len(nums)-1;i++{
		if flag==false{
			break
		}
		flag=false
		for j:=i;j<len(nums);j++{
			if(nums[i]>nums[j]){
				nums[i],nums[j]=nums[j],nums[i]
				flag=true
			}
		}
	}
	return nums
}

func main() {
	var  test = []int {4,1,5,7,2}
	fmt.Println(bubbleSort(test))
}