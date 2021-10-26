package main

import "fmt"

func longestPalindrome(s string) string {
	maxLen:=0
	start,end:=0,0
	for i:=0;i<len(s);i++{
		for j:=0;j<len(s);j++{
			if judge(s,i,j)==true{
				if j-i+1>maxLen{
					maxLen=j-i+1
					start=i
					end=j+1
				}


			}
		}
	}
	//fmt.Println(maxLen)
	return s[start:end]
}
func judge(s string,start,end int)bool{
	i,j:=start,end
	for i<j{
		if s[i]!=s[j]{
			return false
		}else{
			i++
			j--
		}
	}
	return true
}

func main(){
	fmt.Println(longestPalindrome("babad"))
}