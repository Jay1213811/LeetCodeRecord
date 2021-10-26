package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	var sum int =0
	var length int=len(arr)
	for i:=length/20;i<length-(length/20);i++{
		sum+=arr[i]
	}
	return float64(sum)/float64((length-(length/10)))
}