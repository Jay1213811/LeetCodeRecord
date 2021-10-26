package main

import "fmt"

/*
返回出现超过一次的长度为10的字串

用一个mapping映射每一个长度为10的长度的字符串出现的次数
*/

func findRepeatedDnaSequences(s string) []string {
	CountSequenceMap:=make(map[string]int,0)
	str:=""
	res:=[]string{}
	for i:=0;i<len(s)-10;i++{
		str+=s[i:i+10]
		_,ok:=CountSequenceMap[str]
			if ok{
				CountSequenceMap[str]++
			}else{
				CountSequenceMap[str]=1
			}
		str=""	//还原
	}
	//fmt.Println(CountSequenceMap)
	for v:=range CountSequenceMap {
		if CountSequenceMap[v]>1{
			//fmt.Println(CountSequenceMap[v])
			res=append(res,v)
		}
	}
	return res
}
func main()  {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}