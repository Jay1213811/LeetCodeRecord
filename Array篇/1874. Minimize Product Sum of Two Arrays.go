package main

import (
	"fmt"
	"sort"
)

func minProductSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var counter int=0
	var length int
	for i:=0;i<len(nums1);i++{
		length=len(nums1)-i-1
		counter+=nums1[i]*nums2[length]

	}
	return counter
}
func main()  {
	a:=[]int{5,3,4,2}
	b:=[]int{4,2,2,5}
	fmt.Println(minProductSum(a,b))
}