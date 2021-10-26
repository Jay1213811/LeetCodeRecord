package main

import (
	"fmt"
	"math"

	"sort"
)

/*
给你一个整数数组 piles ，数组 下标从 0 开始 ，其中 piles[i] 表示第 i 堆石子中的石子数量。另给你一个整数 k ，请你执行下述操作 恰好 k 次：

选出任一石子堆 piles[i] ，并从中 移除 floor(piles[i] / 2) 颗石子。
注意：你可以对 同一堆 石子多次执行此操作。

返回执行 k 次操作后，剩下石子的 最小 总数。

floor(x) 为 小于 或 等于 x 的 最大 整数。（即，对 x 向下取整）。
*/
func minStoneSum(piles []int, k int) int {
	for i:=0;i<k;i++{
		//要进行k次操作 每次操作前排序
		sort.Ints(piles)
		piles[len(piles)-1]=int(math.Ceil(float64(piles[len(piles)-1])/2.0))
	}
	sum:=0
	for i := 0; i < len(piles); i ++ {
		sum += piles[i]
	}
	return sum

}

func main()  {
	test:=[]int{4,3,6,7}
	test2:=[]int{5,4,9}
	fmt.Println(minStoneSum(test,3))
	fmt.Println(minStoneSum(test2,2))
}