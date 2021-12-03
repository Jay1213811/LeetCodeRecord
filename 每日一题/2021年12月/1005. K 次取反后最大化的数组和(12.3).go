package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	//找到第一个大于等于0的下标.如果Index==0 || Index==-1说明数组权
	//如果数组全为正数 修改最小的k个数
	//1.第一满足的是尽可能的把所有的负数变为正数
	//负数的个数
	number:=len(nums)-getIndex(nums)
	//只有正数
	if number==0{
		if k%2==0{
			return getSum(nums)
		}else{
			nums[0]=nums[0]*-1
		}
	}else if number<=k{//要看是否还多一个要取负的数
		temp:=k-number
		for i:=0;i<number;i++{
			nums[i]=nums[i]*-1
		}
		sort.Ints(nums)
		if temp%2!=0{
			nums[0]=nums[0]*-1
		}

	}else if number>k{     //有负数 负数个数大于k.把负数全变正数
		for i:=0;i<k;i++{
			nums[i]=nums[i]*-1
		}
	}
	return getSum(nums)
}
func getIndex(nums []int)int{
	res:=0
	for _,v:=range nums{
		if v>=0{
			res++
		}
	}
	return res
}
func getSum(nums []int)int{
	res:=0
	for _,v:=range nums{
		res+=v
	}
	return res
}
