package main

import (
	"fmt"
	"sort"
)

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func permute(nums []int) [][]int {
	sort.Ints(nums)
	res:=[][]int{}
	visited:=make(map[int]bool,0)
	var dfs func(path []int)
	dfs= func(path []int) {
		if len(path)==len(nums){
			temp:=make([]int,len(path))
			copy(temp,path)
			res=append(res,temp)
			return
		}
		for i:=0;i<len(nums);i++ {
			if visited[i]==true{
				continue
			}
			path=append(path,nums[i])
			visited[i]=true
			dfs(path)
			visited[i]=false
			path=path[:len(path)-1]
		}
	}
	dfs([]int{})
	return res
}
func main(){
	nums:=[]int{1,1,2}
	fmt.Println(permute(nums))
}