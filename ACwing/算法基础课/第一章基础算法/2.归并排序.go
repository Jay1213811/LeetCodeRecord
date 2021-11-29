package main

import "fmt"

/*
2.归并排序模版
author:https://github.com/Jay1213811
|算法名称|最好时间复杂度|最坏时间复杂度｜平均时间复杂度｜空间复杂度｜是否稳定｜
|归并排序|O(nlogn)	｜O(nlogn)	 ｜  O(nlogn) |  O(n)	｜  稳定	 |
｜---------------------------------------------------------------｜
*/

//归并排序 	主要思想也是分治思想	时间复杂度O(logN)
//tips【归并排序是稳定的】
//稳定是指原数列中两个相等的数在排序完后相对位置不变就叫稳定】
//分治方法和快排有所不同:归并分界点是【下标的中间值】。快排是随机一个数
/*
归并排序的步骤
Step1:确定分界点【下标的中间值】 mid=(l+r)/2	⚠️区别：快排是随机一个数
把数组分为两半如1 3 4 9 2 8
1.确定分界点【下标的中间值】 mid=(l+r)>>1.把长度为 n 的数组分成两个长度为 n/2 的子数组；
2.对这两个子序列分别采用归并排序；
3.归并——合二为一[重难点]将两个排序好的子数组合并成一个最终的排序数组。
*/

func mergeSort(arr []int) []int {
	n:=len(arr)
	// 最后切割只剩下一个元素，可以认为是有序的
	if n<=1{
		return arr
	}
	//Step1 确定分界点
	mid:=n/2
	//Ste2 对这两个子序列分别采用排序；
	left,right:=mergeSort(arr[:mid]),mergeSort(arr[mid:])
	//Step3 归并——合二为一
	var merge func(left,right []int)[]int
	merge=func(left,right []int)[]int{
		result:=make([]int,0)
		i,j:=0,0
		//两个子数组都没遍历完
		for i<len(left) && j<len(right){
			if left[i]<=right[j]{
				result=append(result,left[i])
				i++
			}else{
				result=append(result,right[j])
				j++
			}
		}
		//左数组还没遍历完，右数组遍历完了，直接加上左数组
		if i<len(left){
			result=append(result,left[i:]...)
		}
		if j<len(right){
			result=append(result,right[j:]...)
		}
		return result
	}
	//归并——合二为一
   return 	merge(left,right)
}
func main(){
	nums:=[]int{1,2,3,6,4,4,9,-1}
	nums=mergeSort(nums)
	fmt.Println(nums)

}



