package main
/*
题目：给定一个字符串 s ，找出 至多 包含 k 个不同字符的最长子串 T。
输入: s = "eceba", k = 2
输出: 3
解释: 则 T 为 "ece"，所以长度为 3。
*/
/*
思路：滑动窗口法
定义左右指针。默认都为0
用一个map[byte]int存储出现的每个数个数及种类，key是这个单词，value是出现的次数。如果出现的次数大于k那么就会删除一个元素，左指针移动
*/
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var left,right,maxLen int=0,0,0
	m:=make(map[byte]int,0)
	//遍历
	for left<len(s) && right<len(s){
		//出现字符种类小于k，右指针不断移动，并添加元素进m
		if len(m)<k+1{
			addElementToMap(s[right],m)
		}else{
			//出现字符种类不止k了，删除元素
			deleteElementFromMap(s[left],m)
			left++
		}
		right++
		maxLen=max(maxLen,right-left)
	}
	return maxLen
}
//添加元素
func addElementToMap(element byte,m map[byte]int){
	//先看下元素是否存在
	count,flag:=m[element]
	if flag{
		m[element]=count+1
	}else {
		m[element]=1
	}
}
//删除元素
func deleteElementFromMap(element byte,m map[byte]int)  {
	count,flag:=m[element]
	if flag{
		if count==1{
			delete(m,element)
		}else {
			m[element]=count-1
		}
	}
}

