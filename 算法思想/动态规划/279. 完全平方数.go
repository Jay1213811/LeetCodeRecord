package main

import "fmt"

/*
给定正整数n，找到若干个完全平方数（比如1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
*/
func numSquares(n int) int {
	dp:=make([]int,n+1)
	for i:=1;i<=n;i++{
		dp[i]=i //最坏的情况就是每次+1
		for j:=1;i-j*j>=0;j++{
			dp[i]=min(dp[i],dp[i-j*j]+1) //转移方程
		}
	}
	return dp[n]

}
func min(a int,b int) int{
	if a<b{
		return a
	}else {
		return b
	}
}
func main()  {
	fmt.Println(numSquares(12))
}