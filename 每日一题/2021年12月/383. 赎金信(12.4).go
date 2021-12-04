package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	hashMap:=make(map[byte]int,0)
	for i:=0;i<len(magazine);i++{
		_,ok:=hashMap[magazine[i]]
		if ok{
			hashMap[magazine[i]]+=1
		}else{
			hashMap[magazine[i]]=1
		}
	}
	for i:=0;i<len(ransomNote);i++{
		_,ok:=hashMap[ransomNote[i]]
		if ok{
			hashMap[ransomNote[i]]-=1
			if hashMap[ransomNote[i]]==0{
				delete(hashMap,ransomNote[i])
			}
		}else{
			return false
		}
	}
	return true
}
func main()  {
	fmt.Println(canConstruct("a","b"))
}