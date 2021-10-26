package main

import (
	"fmt"
	"sort"
)
func minOperations(grid [][]int, x int) (ans int) {
	n := len(grid) * len(grid[0])
	a := make([]int, 0, n)
	for _, row := range grid {
		for _, v := range row {
			if (v-grid[0][0])%x != 0 { // 以其中一元素为基准，若所有元素与它的差均为 x 的倍数，则任意两元素之差为 x 的倍数
				return -1
			}
		}
		a = append(a, row...)
	}
	sort.Ints(a) // 也可以用求第 k 大算法来找中位数
	for _, v := range a {
		ans += abs2(v-a[n/2]) / x
	}
	return
}

func abs2(x int) int {
	if x < 0 {
		return -x
	}
	return x
}




func main()  {
	grid:=[][]int{{931,128},{639,712}}
	x:=73
	//grid:=[][]int{{2,4},{6,8}}
	//x:=2
	//grid:=[][]int{{1,5},{2,3}}
	//x:=1
	fmt.Println(minOperations(grid,x))
}