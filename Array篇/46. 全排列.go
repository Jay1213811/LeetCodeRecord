package main

import "fmt"

func lengthOfLastWord(s string) int {
	tmp:=len(s)-1
	for j:=len(s)-1;j>=0;j--{
		if string(j)==" "{
			tmp--
		}else{
			break
		}
	}
	res:=0
	for j:=tmp;j>=0;j--{
		if string(j)!=" "{
			res++
		}else{
			break
		}
	}
	return res
}
func main()  {
	fmt.Println(lengthOfLastWord("hello world"))
}