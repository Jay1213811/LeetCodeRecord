package main

import "fmt"

/*
1.快速排序模版
author:https://github.com/Jay1213811
*/

//1.快速排序
//空间复杂度是Θ(lg𝑛)；平均情况下快速排序的时间复杂度是Θ(𝑛lg𝑛)，最坏情况是𝑛2，
//快排的主要思想是分治
/*
主要步骤
先设数组为nums []int;l,r:=0,len(nums)-1
Step1:确定分界点:分界点一般取nums[l]或nums[(l+r)/2]【注意边界划分为Left,j	j+1 right一定不能取
base为nums[r]不然会无限递归下去】
Step2:根据分界点.重新调整区间.使得左边都小于等于分界点.右边都大于等于分界点
Step3:递归.不断给左右两个区间排序.等左右两边都排序完了.那么整个数组就排序完了
*/
func quicklySort(nums []int,left,right int){
	//边界问题
	if left>=right{
		return
	}
	//定义左右开始遍历的位置,和基准
	l,r,base:=left-1,right+1,nums[(left+right) >> 1]
	//遍历完成一轮
	for l<r{
		//进来前先各往前走一步,再开始对比
		l++
		r--
		//左边都小于基准
		for nums[l]<base{
			l++
		}
		//右边都大于基准
		for nums[r]>base {
			r--
		}
		if l<r{
			nums[l],nums[r]=nums[r],nums[l]
		}
	}
	//不断递归调用,完成左右两个区间的排序即完成
	quicklySort(nums,left,r)
	quicklySort(nums,r+1,right)
}


func main(){
	nums:=[]int{1,3,3,4,6,2}

	quicklySort(nums,0,len(nums)-1)	//调用快排
	fmt.Println(nums)


}