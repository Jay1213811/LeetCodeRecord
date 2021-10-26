package main

import "fmt"

/*
滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。
一旦出现了重复字符，就需要缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。
以此类推，每次移动需要计算当前长度，并判断是否需要更新最大长度，最终最大的值就是题目中的所求。*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, 0

	for left < len(s) {
		if right < len(s) && freq[s[right]-'a'] == 0 {
			freq[s[right]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left)
	}
	return result
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main()  {
	fmt.Println(lengthOfLongestSubstring("abcad"))
}
