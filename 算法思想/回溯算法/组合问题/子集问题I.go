package main

import "fmt"

//输入一个长度大于2的数组,找出长度为2的子数组
//如输入 nums=[1 2 3 4]
//输出[[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]

func subset(nums []int) [][]int {
	res:=[][]int{}
	var backTracking func(path []int,startIndex int)
	backTracking= func(path []int,startIndex int) {
		if len(path)==2{
			temp:=make([]int,len(path))
			copy(temp,path)
			res=append(res,temp)
			return
		}
		for i:=startIndex;i<len(nums);i++{
			path=append(path,nums[i])
			backTracking(path,i+1)
			path=path[:len(path)-1]
		}
	}
	backTracking([]int{},0)
	return res
}
func main(){
	nums:=[]int{1,2,3,4}
	fmt.Println(subset(nums))
}