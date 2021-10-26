package main

import (
	"fmt"
	"strings"
)

/*
Time:2021年8月8日周赛
第一题:

给你一个字符串 s 和一个字符串数组 words ，请你判断 s 是否为 words 的 前缀字符串 。

字符串 s 要成为 words 的 前缀字符串 ，需要满足：s 可以由 words 中的前 k（k 为 正数 ）个字符串按顺序相连得到，且 k 不超过 words.length 。

如果 s 是 words 的 前缀字符串 ，返回 true ；否则，返回 false 。
思路:
*/
func isPrefixString(s string, words []string) bool {
	var str2 []string = []string{}
	for i:=0;i<len(words);i++{
		str2=append(str2,words[i])
		s3:=strings.Join(str2,"")
		if s3==s{
			return true
		}
	}
	return false
}
func main()  {
	s := "iloveleetcode"
	var str2 []string = []string{"apples","i","love","leetcode"}
	fmt.Println(isPrefixString(s,str2))
}