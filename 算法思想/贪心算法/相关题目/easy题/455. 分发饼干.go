package main

import "sort"

func findContentChildren(greed, size []int) (ans int) {
	sort.Ints(greed)
	sort.Ints(size)
	n, m := len(greed), len(size)
	for i, j := 0, 0; i < n && j < m; i++ {
		for j < m && greed[i] > size[j] {
			j++
		}
		if j < m {
			ans++
			j++
		}
	}
	return
}
