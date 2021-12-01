package main

import "fmt"
//author:https://github.com/Jay1213811

/*前缀和是一种重要的预处理，能大大降低查询的时间复杂度。*/
//1.一维前缀和模版【主要作用在于快速求数组某一区间范围的和】
//S[i] = a[1] + a[2] + ... a[i]
//a[l] + ... + a[r] = S[r] - S[l - 1]

//2.二维前缀和【主要作用在于快速求某一子矩阵的和】
//S[i, j] = 第i行j列格子左上部分所有元素的和
//以(x1, y1)为左上角，(x2, y2)为右下角的子矩阵的和为：
//a[i,j]=S[x2, y2] - S[x1 - 1, y2] - S[x2, y1 - 1] + S[x1 - 1, y1 - 1]
arr=
[1 5 6 8]
[9 6 7 3]
[5 3 2 4]
sum=
	[1  6  12 20]
	[10 21 34 45]
	[15 29 44 59]
//sum矩阵的求法
//边界情况
1.i==0 && j==0
sum[0][0]=g[0][0]
2.i==0 类似于一维前缀和
sum[0][i]=sum[0][i-1]+g[0][i]
3.j==0
sum[i][0]=sum[i-1][0]+g[i][0]
//通项
sum{x1,0}{x2,0}=sum[x2][0]-sum[x1-1][0]

二维前缀和的例题
求{1,1} 到{2,2}小矩阵的和即6+7+3+2=18
=sum{2,2}(44)-sum{1,2}(12)-sum(2,1)(15)+sum(1,1)1
=44-12-15+1
=18	//✅

//二维前缀和板子
func getPreSum(g [][]int)[][]int{
	sum:=make([][]int,len(g))
	//边界情况【如果下标从1开始就不需要处理边界情况即第一行和第一列全0的情况】
	sum[0][0]=g[0][0]
	for i:=1;i<len(sum);i++{
		sum[i][0]=sum[i-1][0]+g[i][0] //第一列
	}
	for j:=1;j<len(sum[0]);j++{
		sum[0][j]=sum[0][j-1]+g[0][j]//第一行
	}
	//通式
	for i:=1;i<len(sum);i++{
		for j:=1;j<len(sum[0]);j++{
			//下面的小块=整个矩形-左边一大块-上面一大块+左上角多减了一次的块
			//g[i][j]=sum[i][j]-sum[i-1][j]-sum[i][j-1]+sum[i-1][j-1]
			sum[i][j]=g[i][j]+sum[i][j-1]+sum[i-1][j]-sum[i-1][j-1]
		}
	}
	return sum
}




//原:a1 a2 a3 ... an
//前缀和S=a1+a2+...an
//Q1:如何求S前缀和
//for 循环求一下即可
//s[i]=s[i-1]+a[i]
//Q2:前缀和有什么作用
//A2:可以方便求某一段的数的和[最大的作用]
//如[l,r]这一段的数之和=s(r)-s(l-1).
//其中
//s(r)=a1+a2+...+a[l-1]+...+a[r]
//s(l-1)=a1+a2+...+a[l-1]

/*
一维前缀和模版题 LC 560 和为 K 的子数组
例题.给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。
输入：nums = [1,1,1], k = 2
输出：2
*/

//前缀和：nums 的第 0 项到 当前项 的和。
//定义 prefixSum 数组，prefixSum[x]：第 0 项到 第 x 项 的和。
//prefixSum[x] = nums[0] + nums[1] +…+nums[x]
//prefixSum[x]=nums[0]+nums[1]+…+nums[x]

//nums 的某项 = 两个相邻前缀和的差：
//nums[x] = prefixSum[x] - prefixSum[x - 1]
//nums[x]=prefixSum[x]−prefixSum[x−1]
//
//nums 的 第 i 到 j 项 的和，有：
//nums[i] +…+nums[j]=prefixSum[j] - prefixSum[i - 1]
//nums[i]+…+nums[j]=prefixSum[j]−prefixSum[i−1]
//
//当 i 为 0，此时 i-1 为 -1，我们故意让 prefixSum[-1] 为 0，使得通式在i=0时也成立：
//nums[0] +…+nums[j]=prefixSum[j]
//nums[0]+…+nums[j]=prefixSum[j]
/*前缀和+哈希优化*/
func subarraySum2(nums []int, k int) int {
		count, pre := 0, 0
		m := map[int]int{}
		m[0] = 1
		for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre - k]; ok {
		count += m[pre - k]
	}
		m[pre] += 1
	}
		return count
	}

func subarraySum(nums []int, k int) int {
		ans:=0
		prefix:=make([]int,len(nums))
		prefix[0]=nums[0]
		for i:=1;i<len(nums);i++{
		prefix[i]+=prefix[i-1]+nums[i]
	}
		// fmt.Println(prefix)
		for i:=0;i<len(prefix);i++{
		if prefix[i]==k{
		ans++
	}
		for j:=i+1;j<len(prefix);j++{
		if prefix[j]-prefix[i]==k{
		ans++
	}
	}
	}
		return ans
}
	//tips 为了让i=0的时候也成立可以这样
func sumOddLengthSubarrays(arr []int) int {
	preSum:=make([]int,len(arr)+1)
	preSum[0]=arr[0]
	for i:=1;i<=len(arr);i++{
		preSum[i]=preSum[i-1]+arr[i-1]
		fmt.Println(preSum[i])
	}
	//输入arr=1 2 3
	//前缀和S为 1 2 4 7
	return 0
}