package main

import (
	"fmt"
	"strconv"
)

/*
当复习快排了
第一步写一个快排.将score从高到低排序
score=[10,3,8,9,4]
===>排序后[3,4,8,9,10]
通过hashMap.key为score的分数.value是其在排序后数组中的相对位置
如
10——len(score)-4
3-len(score)-0
排名前三的改成具体名称。其他的返回编号
*/
func findRelativeRanks(score []int) []string {
	hashMap:=make(map[int]int,0)
	rank:=make([]int,len(score))
	copy(rank,score)
	rank=QuicklySort(rank,0,len(rank)-1)
	for i:=0;i<len(rank);i++{
		hashMap[rank[i]]=len(score)-i
	}
	res:=[]string{}
	for i:=0;i<len(score);i++{
		if hashMap[score[i]]==1{
			res=append(res,"Gold Medal")
		}else if hashMap[score[i]]==2{
			res=append(res,"Silver Medal")
		}else if hashMap[score[i]]==3{
			res=append(res,"Bronze Medal")
		}else{
			res=append(res,strconv.Itoa(hashMap[score[i]]))
		}
	}
	return res
}
func QuicklySort(nums []int,left,right int)[]int{
	l,r,base:=left-1,right+1,nums[left+(right-left)>>1]
	if left>=right{
		return nums
	}
	for l<r{
		l++
		r--
		for nums[l]<base{
			l++
		}
		for nums[r]>base{
			r--
		}
		if l<r{
			nums[l],nums[r]=nums[r],nums[l]
		}
	}
	QuicklySort(nums,left,r)
	QuicklySort(nums,r+1,right)
	return nums
}
func main(){
	//score:=[]int{5,4,3,2,1}
	score2:=[]int{10,3,8,9,4}
	//fmt.Println(findRelativeRanks(score))
	fmt.Println(findRelativeRanks(score2))
}