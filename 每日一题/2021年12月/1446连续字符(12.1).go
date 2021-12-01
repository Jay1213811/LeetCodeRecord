package main

import "fmt"

/*
给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。

请你返回字符串的能量。
*/
//标准双指针题
//left,right:=0,0
//移动right节点,如果s[left]==s[right] 移动right即可
//如果不等于左边 left++ right++ maxLen=max(maxLen,right-left+1)
//
func maxPower(s string) int {
	left,right:=0,0
	maxLen:=0
	for right<len(s){
		if s[left]==s[right]{
			right++

		}else {
			//maxLen=max(maxLen,right-left)
			left=right
			right++
		}
	}
	//这里有一个track,如果返回maxLen会WA.原因是如果right>=len(s)后不会走else判断
	return max(maxLen,right-left)
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
func main()  {
	fmt.Println(maxPower("aa"))
}