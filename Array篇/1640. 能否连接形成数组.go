package main

import "fmt"

func canFormArray(arr []int) map[int]int {
	arrMap:=map[int]int{}
	for i,v := range arr{
		arrMap[v]=i
	}
	return arrMap
}
func main()  {
	var test1 =[]int{1,2,9}
	fmt.Println(canFormArray(test1))
}