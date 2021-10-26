package main

import (
	"fmt"
	"sort"
)

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	res:=make([]int,0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	sort.Ints(nums3)
	nums1=del(nums1)
	nums2=del(nums2)
	nums3=del(nums3)
	for i:=0;i<len(nums1);i++{
		for j:=0;j<len(nums2);j++{
			for k:=0;k<len(nums3);k++{
				if nums1[i]==nums2[j] {
					res=append(res,nums1[i])
				}else if nums2[j]==nums3[k]{
					res=append(res,nums2[j])
				}else if nums1[i]==nums3[k] {
					res=append(res,nums1[i])
				}
			}
		}
	}
	res=del(res)
	return res
}
func del(list []int) []int {
	var x []int = []int{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}
func main()  {
	nums1:=[]int{1,1,3,2}
	nums2:=[]int{2,3}
	nums3:=[]int{3}
	fmt.Println(twoOutOfThree(nums1,nums2,nums3))
}