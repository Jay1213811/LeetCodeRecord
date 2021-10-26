package main

import "fmt"

func plusOne(digits []int) []int {
	for i:=len(digits)-1;i>0;i--{
		digits[i]++
		if digits[i]<10{
			return digits
		}else {
			digits[i]=0
		}
	}
	//超位了
	digits[0]=1
	digits=append(digits,0)
	return digits
}
func main()  {
	var test1 =[]int{1,2,9}
	var test2=[]int{9,9,9}
	var test3=[]int{1,2,3}
	fmt.Println(plusOne(test1))
	fmt.Println(plusOne(test2))
	fmt.Println(plusOne(test3))
}