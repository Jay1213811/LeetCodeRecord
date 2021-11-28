package main

import "fmt"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	//根据时间排序
	for i:=0;i<len(meetings);i++{
		for j:=i+1;j<len(meetings);j++{
			if meetings[i][2]>meetings[j][2]{
				meetings[i],meetings[j]=meetings[j],meetings[i]
			}
		}
	}
	res:=[]int{}
	res=append(res,0)
	res=append(res,firstPerson)
	for i:=0;i<len(meetings);i++{
		if judge(res,meetings[i][0])==true && judge(res,meetings[i][1])!=true{
			res=append(res,meetings[i][1])
		}else if judge(res,meetings[i][0])!=true && judge(res,meetings[i][1])==true{
			res=append(res,meetings[i][0])
		}else{
			continue
		}
	}
	return res
}
func judge(nums []int,target int)bool{
	for _,v:=range nums{
		if v==target{
			return true
		}
	}
	return false
}
func main()  {
	n:=4
	//meetings:=[][]int{{1,2,5},{2,3,8},{1,5,10}}
	meetings:=[][]int{{3,1,3},{1,2,2},{0,3,3}}
	fmt.Println(findAllPeople(n,meetings,3))
}