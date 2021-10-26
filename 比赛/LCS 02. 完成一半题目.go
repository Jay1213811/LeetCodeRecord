package main

import (
	"fmt"
	"sort"
)

//至少包含多少种知识点类型
func halfQuestions(questions []int) int {
	leetcodePeople:=len(questions)/2
	CountMap:=make(map[int]int,0)
	for i:=0;i<len(questions);i++{
		_,ok:=CountMap[questions[i]]
		if ok{
			CountMap[questions[i]]++
		}else{
			CountMap[questions[i]]=1
		}
	}
	ans:=make([]int,0)
	for v:=range(CountMap){
		ans=append(ans,CountMap[v])
	}
	sort.Ints(ans)//1 2 3
	res:=0
	//fmt.Println(ans)
	//fmt.Println(CountMap)
	for i:=len(ans)-1;i>=0;i--{
		if leetcodePeople>0{
			leetcodePeople-=ans[i]
			res++
		}
	}
	return res
}
func main()  {
	questions :=[]int{1000,1000}
	fmt.Println(halfQuestions(questions))
}