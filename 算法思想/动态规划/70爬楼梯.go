package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

思路：采用动态规划思想	单串dp

*/
func climbStairs(n int) int {
	dp:=make([]int,n+1)
	dp[0]=1
	dp[1]=1
	for i:=2;i<=n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}
func main()  {
	fmt.Println(climbStairs(6))
}