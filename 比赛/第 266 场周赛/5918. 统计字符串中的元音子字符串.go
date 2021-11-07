package main

import (
	"fmt"
)

func countVowelSubstrings(word string) int {
	ans:=0

	for i:=0;i<len(word);i++{
		for j:=i;j<len(word);j++{
			temp:=[]byte(word[i:j+1])
			//fmt.Println(string(temp))
			if judge(temp)==true{
				ans++
			}
		}
	}
	return ans

}
func judge(wordByte []byte)bool{
	a,e,i,o,u:=0,0,0,0,0
	wrong:=0
	for j:=0;j<len(wordByte);j++{
		if wordByte[j]=='a'{
			a++
		}else if wordByte[j]=='e'{
			e++
		}else if wordByte[j]=='i'{
			i++
		}else if wordByte[j]=='o'{
			o++
		}else if wordByte[j]=='u'{
			u++
		}else{
			wrong++
		}
	}
	if wrong>0{
		return false
	}
	if a>0 && e>0 && i>0 && o>0 && u>0{
		return true
	}
	return  false
}
func main()  {
	fmt.Println(countVowelSubstrings("aeiouu"))
}