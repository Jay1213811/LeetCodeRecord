package main

import "fmt"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	f := make([]int, n)
	for i := range f {
		f[i] = inf
	}
	f[src] = 0
	ans := inf
	for t := 1; t <= k+1; t++ {
		g := make([]int, n)
		for i := range g {
			g[i] = inf
		}
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			g[i] = min(g[i], f[j]+cost)
		}
		f = g
		ans = min(ans, f[dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main(){
	n := 3
	edges :=[][]int{{0,1,100},{1,2,100},{0,2,500}}
	src,dst,k:=0,2,1
	fmt.Println(findCheapestPrice(n,edges,src,dst,k))
}
