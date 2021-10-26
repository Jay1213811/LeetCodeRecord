package main

import "fmt"

func numOfPairs(nums []string, target string) int {
	counter:=0
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums);j++{
			if nums[i]+nums[j]==target && i!=j{
				fmt.Println(i,j)
				counter++
			}
		}
	}
	return counter
}
func main()  {
	nums:=[]string{"777","7","77","77"}
	fmt.Println(numOfPairs(nums,"7777"))
}