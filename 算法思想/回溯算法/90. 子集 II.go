package main

import (
	"fmt"
	"sort"
)

/*
难就难在去重
1   2   2

*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	sl := make([]int, len(nums))
	dfs(nums, &ret, sl, 0, 0)
	return ret
}

func dfs(nums []int, ret *[][]int, sl []int, start int, layer int) {
	dst := make([]int, layer)
	copy(dst, sl)
	*ret = append(*ret, dst)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		sl[layer] = nums[i]
		dfs(nums, ret, sl, i+1, layer+1)
	}
}

func main(){
	nums:=[]int{1,2,2}
	fmt.Println(subsetsWithDup(nums))
}


