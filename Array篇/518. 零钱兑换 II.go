package main

import (
	"fmt"
)
/*
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
思路：

*/
func change(amount int, coins []int) int {
	// 定义dp数组
	dp := make([]int, amount+1)
	// 初始化
	dp[0]  = 1
	// 遍历顺序
	for i := 0 ;i < len(coins);i++ {

		for j:= coins[i] ; j <= amount ;j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}
func main(){
	a:=[]int{1,2,5}
	fmt.Println(change(5,a))
}