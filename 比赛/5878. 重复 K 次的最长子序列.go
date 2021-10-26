package main

import (
	"fmt"
	"sort"
)

func longestSubsequenceRepeatedK(s string, k int) int {
	m:=make(map[byte]int)
	for i:=0;i<len(s);i++{
		v,ok:=m[s[i]]
		if(ok){
			m[s[i]]=v+1
		}else{
			m[s[i]]=1
		}
	}
	bytes:=[]byte(s)
	sort.Ints(s)

	fmt.Println(m)
	return 0
}
func main()  {
	fmt.Println(longestSubsequenceRepeatedK("sfasdfgvds",2))
}
