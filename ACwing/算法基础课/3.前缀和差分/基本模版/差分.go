package main

import "fmt"

//差分分为一维差分和二维差分
/*模版:*/
//给区间[l, r]中的每个数加上c：B[l] += c, B[r + 1] -= c
/*
差分其实是前缀和的逆运算.差分的含义如下
a1 a2 a3 ......an
构造b1 b2 .... bn
使得ai=b1+b2+...+bi【即使得ai是b[1,i]的前缀和】
构造方法很简单
b1=a1
b2=a2-a1
b3=a3-a2
........
bn=an-an-1
构造方法有很多,假想一个b数组可以满足ai=b1+b2+...+bi.那么可以称b数组是a数组的差分,a数组是b数组的前缀和
‼️差分的作用:构造出b数组后，我们可以通过O(n)的复杂度得到a数组
*/
//例题.
//输入一个数组，接下来输入m个操作。
//每个操作包含三个整数l r c，表示将序列中[l,r]之间的每个数加上c.输出进行完操作后的序列
/*'
输入 nums=[]int{1 2 2 1 2 1} l=1 r=3 c=1
输出 1 3 3 2 2 1
*/
func differ(nums []int,l,r,c int)[]int {
	//为了避免溢出，差分数组大一位
	difference:=make([]int,len(nums)+1)
	difference[l]+=c
	difference[r+1]-=c
	//求differce差分数组前缀和
	for i:=1;i<len(nums);i++{
		difference[i]+=difference[i-1]
	}
	//还原得出结果
	for i:=0;i<len(nums);i++{
		nums[i]+=difference[i]
	}
	return nums
}
func main()  {
	arr:=[]int{1,3,7,5,2}
	arr=differ(arr,2,4,5)
	arr=differ(arr,1,3,2)
	arr=differ(arr,0,2,-3)
	fmt.Println(arr)
}
