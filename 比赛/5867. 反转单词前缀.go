package main

import (
	"fmt"

)

func reversePrefix(word string, ch byte) string {
	wordbyte:=[]byte(word)
	for i:=0;i<len(wordbyte);i++{
		if wordbyte[i]==ch{
			Reverse(wordbyte,0,i)
			break
		}
	}

	str:=string(wordbyte)
	return str
}
func Reverse(a []byte,left,right int){

	for left<right{
		a[left],a[right]=a[right],a[left]
		left++
		right--
	}

}
func main()  {
	word:="abcdefg"

	fmt.Println(reversePrefix(word,'d'))
}