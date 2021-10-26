package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	sort(nums1)
	sort(nums2)
	var nums3 []int
	var i int =0
	var j int =0
	var k int =0
	if i<len(nums1)&&j<len(nums2){
		if(nums1[i]<nums2[j]){
			i++
		}else if(nums1[i]>nums2[j]){
			j++
		}else if(nums1[i]==nums2[j]){
			nums3[k]=nums1[i]
			k++
			i++
			j++
		}
	}
	return nums3
}
func sort(nums1 []int)  {
	for i:=0;i<len(nums1);i++{
		for j:=1;j<i;j++ {
			if(nums1[i]>nums1[j]){
				var temp int
				temp=nums1[i]
				nums1[i]=nums1[j]
				nums1[j]=temp
			}
		}
	}

}

func main()  {
	var  nums1 = []int {1,6,2,3,4}
	var  nums2 = []int {1,7,3}
	fmt.Println(intersect(nums1,nums2))

}