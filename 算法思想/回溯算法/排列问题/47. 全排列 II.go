package main

import (
	"fmt"
	"sort"
)

//组合无顺序,排列有顺序
//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	res:=[][]int{}
	sort.Ints(nums)
	var dfs func(path []int)
	visited:=make(map[int]bool,0)
	dfs= func(path []int) {
		if len(path)==len(nums) {
			temp:=make([]int,len(nums))
			copy(temp,path)
			res=append(res,temp)
			return
		}
		for i:=0;i<len(nums);i++{
			if visited[i]==true{
				continue
			}
			if i>0 && nums[i]==nums[i-1] && visited[i-1]==false{
				continue
			}
			visited[i]=true
			path=append(path,nums[i])
			dfs(path)
			visited[i]=false
			path=path[:len(path)-1]
		}
	}
	dfs([]int{})
	return res
}
func main()  {
	nums:=[]int{1,1,2}
	fmt.Println(permuteUnique(nums))
}