package main

import "fmt"

/*
思路：
第一步将字符串每个数字转换为整数存放到一个数组
通过for循环不断减少k的值，让它为0
不断累加
*/
func getLucky(s string, k int) int {
	count:=make([]int,0)
	res:=0
	//a的ascii为97
	for i:=0;i<len(s);i++{
		count=append(count,int(s[i])-96)
	}
	//第一次转换将数组中每个数字相加
	for i:=0;i<len(count);i++{
			res=res+sumn(count[i])
	}
	for k!=1{
		res=sumn(res)
		k--
	}
	return res
}
func sumn(n int) int {
	var sum int=0
	//输入一个整数，返回所谓位置的数字之和
	for n!=0{
		sum=sum+n%10
		n=n/10
	}
	return sum
}
func main()  {
	fmt.Println(getLucky("leetcode",2))
}