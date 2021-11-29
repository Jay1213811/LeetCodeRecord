package main

import "fmt"

//3.整数二分模版
//用于二分查找某个数
//author:https://github.com/Jay1213811
/*
二分的本质不是单调性.
二分和单调性的关系:单调一定能够二分.不单调不一定不能二分
二分的本质是:
如果在区间[l,r] 中.存在区间左半边或者右半边满足某种性质。那么就可以用二分.

⚠️情况1:区间[l, r]被划分成[l, mid - 1]和[mid, r]
重点‼️如果更新l=mid的话,mid=(l+r+1)>>1。为了避免溢出建议写成l+((r-l)>>1)
.
划分区间一分为二:左半边left:=[l,mid-1]右半边[mid,r]
if check(mid)==true那么要找的数在右边‼️ l=mid
否则在左边r=mid-1
⚠️情况2：区间[l, r]被划分成[l, mid]和[mid + 1, r]
如果在区间[l,r]中,mid=l+((r-l+1)>>1).在区间左边满足某种性质
划分区间一分为二:左半边left:=[l,mid]右半边[mid+1,r]
if check(mid)==true那么要找的数在左边‼️ r=mid
否则在左边l=mid+1
⚠️分成两半左边[l,mid-1]右边[mid,r]两个区间是没有交集的
背:如果更新l=mid的话,mid=(l+r+1)>>1。
如果更新r=mid的话mid不需要+1.mid=(l+r)>>1
*/
func check(n int)bool{
	//随便写的check函数
	if n%2==0{
		return true
	}
	return false
}
// 区间[l, r]被划分成[l, mid - 1]和[mid, r]时使用：
func binSearch1(l,r int) int{
	for l<r{
		mid:=(l+r+1)>>1
		if (check(mid))==true{
			l=mid
		}else{
			r=mid-1
		}
	}
	return l
}
// 区间[l, r]被划分成[l, mid]和[mid + 1, r]时使用：
func binSearch2(l,r int)int{
	for l<r{
		mid:=(l+r)>>1
		if (check(mid))==true{
			r=mid
		}else{
			l=mid+1
		}
	}
	return l
}
//一个具体的例子:
/*
Acwing789:数的范围【题目有缩减】
给定一个升序排列长度为n的数组,以及一个要查询的数返回它的启始位置和结束位置
如
1.输入nums:=[]int{1,2,2,3,3,4}	k=3
返回 3 4
2.输入nums:=[]int{1,2,2,3,3,4}	k=4
返回 5 5
3.2.输入nums:=[]int{1,2,2,3,3,4}	k=5
返回 -1 -1
*/
func getNumsRange(nums []int,k int)[]int{
	l,r:=0,len(nums)-1
	ans:=[]int{}
	for l<r{
		//寻找左边第一个出现
		mid:=(l+r)/2
		//1 2 2 3 3 4 mid=2 那么要找的mid一定在右半边
		//我们要找的边界在左半边
		if nums[mid]>=k{
			r=mid
		}else{
			l=mid+1
		}
	}
	fmt.Println(l)
	ans=append(ans,l)
	if nums[l]!=k{
		return []int{-1,-1}
	}else{
		//寻找右边最后一个出现
		l,r=0,len(nums)-1
		for l<r{
			mid:=l+r+1>>1
			//1  2 2 3 3 4 mid=2 那么要找的mid一定在右半边
			//我们要找的边界在左半边
			if nums[mid]<=k{
				l=mid
			}else {
				r=mid-1
			}
		}
		ans=append(ans,l)
	}
	return ans

}
func main()  {
	nums:=[]int{1,2,2,3,3,4}
	fmt.Println(getNumsRange(nums,3))//[3,4]
	//fmt.Println(getNumsRange(nums,4))//[5,5]
	//fmt.Println(getNumsRange(nums,5))//[-1,-1]

}

