package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int)[][]int{
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<=intervals[j][0]
	})
	res:=[][]int{}
	start,end:=intervals[0][0],intervals[0][1]
	for i:=1;i<len(intervals);i++{
		if intervals[i][0]<=end{
			end=max(intervals[i][1],end)
		}else{
			res=append(res,[]int{start,end})
			start=intervals[i][0]
			end=intervals[i][1]
		}
	}
	return append(res,[]int{start,end})
}
func max(a,b int)int  {
	if a>b{
		return a
	}
	return b
}
func main(){
	intrevals:=[][]int{{1,3},{8,10},{15,18},{2,6}}
	fmt.Println(merge(intrevals))
}