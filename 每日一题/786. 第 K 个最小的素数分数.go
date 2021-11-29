package main

import (
	"fmt"
	"sort"
	"strconv"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {

	res:=[]float64{}
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			num1, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", float64(arr[i])/float64(arr[j])), 64) // 保留2位小数
			 // 0.43
			res=append(res,num1)

		}
	}
	sort.Float64s(res)
	need:=res[k-1]
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			num1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(arr[i])/float64(arr[j])), 64) // 保留2位小数
			if need==num1{
				return []int{arr[i],arr[j]}
			}
		}
	}

	return nil
}
func main()  {
	arr:=[]int{1,29,46}
	k:=1
	fmt.Println(kthSmallestPrimeFraction(arr,k))
}