package main

import (
	"fmt"

)

func combinationSum(candidates []int, target int) [][]int {
	res:=[][]int{}

	var dfs func(path []int,startIndex int,need int)
	dfs= func(path []int, startIndex int,need int) {
		if need<=0{
			if need==0{
				temp:=make([]int,len(path))
				copy(temp,path)
				res=append(res,temp)
				return
			}
			return
		}
		for i:=startIndex;i<len(candidates);i++{
			need=need-candidates[i]
			if need<0{
				continue
			}
			path=append(path,candidates[i])
			dfs(path,i,need)
			need=need+candidates[i]
			path=path[:len(path)-1]
		}
	}
	dfs([]int{},0,target)
	return res

}
func main()  {
	nums:=[]int{2,3,5}
	fmt.Println(combinationSum(nums,8))
}