package main

import (
	"strings"
)

/*
给你一个字符串数组 patterns 和一个字符串 word
统计 patterns 中有多少个字符串是 word 的子字符串。返回字符串数目。
子字符串 是字符串中的一个连续字符序列。

输入：patterns = ["a","abc","bc","d"], word = "abc"
输出：3
解释：
- "a" 是 "abc" 的子字符串。
- "abc" 是 "abc" 的子字符串。
- "bc" 是 "abc" 的子字符串。
- "d" 不是 "abc" 的子字符串。
patterns 中有 3 个字符串作为子字符串出现在 word 中。
*/

func numOfStrings(patterns []string, word string) int {
	var counter int=0
	for i:=0;i<len(patterns);i++{
		if strings.Contains(word,patterns[i]){
			counter++
		}
	}
	return counter
}