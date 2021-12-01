package main

import "fmt"

//选择排序算法也是可以优化的，既然每轮遍历时找出了最小值，何不把最大值也顺便找出来呢？这就是二元选择排序的思想。
//
//使用二元选择排序，每轮选择时记录最小值和最大值，可以把数组需要遍历的范围缩小一倍。
//
func selectionSort2(nums []int)[]int {
		var minIndex int
		var maxIndex int
		for i:=0;i<len(nums)/2;i++{
			minIndex=i
			maxIndex=i
			for j:=i+1;j<len(nums)-i;j++{
				if(nums[minIndex]>nums[j]){
					minIndex=j
				}else if nums[maxIndex]<nums[j]  {
					maxIndex=j
				}
			}
			if(minIndex==maxIndex){
				break
			}
			//最小元素换到首位
			nums[i],nums[minIndex]=nums[minIndex],nums[i]
			//如果最大值小标刚好为i，由于nums[i]和nums[minIndex]以及交换了。所以更新maxIndex值
			if(maxIndex==i){
				maxIndex=minIndex
			}
			//最大元素换到末尾
			var lastIndex int=len(nums)-1-i
			nums[lastIndex],nums[maxIndex]=nums[maxIndex],nums[lastIndex]
		}
		return nums
}
func main() {
	var  test = []int {4,1,5,7,2}
	fmt.Println(selectionSort2(test))
}