package main

import "fmt"

//选择排序的思想是：双重循环遍历数组，每经过一轮比较，找到最小元素的下标，将其交换至首位。
func selectionSort(nums []int)  []int {
	var minIndex int
	for i:=0;i<len(nums)-1;i++{
		minIndex=i
		for j:=i+1;j<len(nums);j++{
			if(nums[minIndex]>nums[j]){
				minIndex=j
			}

		}
		nums[i],nums[minIndex]=nums[minIndex],nums[i]
	}

	return nums

}
func main() {
	var  test = []int {4,1,5,7,2}
	fmt.Println(selectionSort(test))
}