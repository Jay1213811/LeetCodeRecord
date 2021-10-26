package main

import (
	"fmt"
	"sort"
)

/*
思路:排序后返回
*/

func smallestK(arr []int, k int) []int {
	result:=make([]int, k)
	sort.Ints(arr)
	for i:=0;i<k;i++{
		result=append(result,arr[i])
	}
	return  result
}
func main()  {
	fmt.Println()
}

