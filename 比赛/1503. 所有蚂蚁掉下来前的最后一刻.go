package main

import (
	"fmt"
	"sort"
)

func getLastMoment(n int, left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	//没有者向右的
	if len(right)==0{
		return left[0]
	}
	//没有者向左的
	if len(right)==0{
		return n-right[0]
	}

	rightFallTime:=n-right[0]
	//左边掉下去花费时间
	leftFallTime:=left[len(left)-1]

	return max(leftFallTime,rightFallTime)



}

func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}
func main()  {
	left:=[]int{0}
	right:=[]int{}
	fmt.Println(getLastMoment(1000,left,right))
}