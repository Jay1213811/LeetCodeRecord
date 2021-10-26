package main

import "fmt"
/*
	从前往后遍历。首先定义一个index值非常小，保证如果第一个值就为1的间距肯定大于K。
	两个1之间的间距公式为distance:=i-index-1
	如果间距小于k则直接返回false。否则继续遍历找到1。同时将index定义为最新出现的i的值
*/
func kLengthApart(nums []int, k int) bool {
	var index int=-10000
	for i:=0;i<len(nums);{
		if nums[i]==1{
			distance:=i-index-1
			if(distance<k){
				return false
			}
			index=i
		}
		i++
	}
	return true

}
func main()  {
	var nums=[]int{1,0,0,0,1,0,0,1}
	var k int=2
	fmt.Println(kLengthApart(nums,k))
}