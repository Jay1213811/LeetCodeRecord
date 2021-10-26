package main

import "fmt"

func minimumMoves(s string) int {
	index:=0
	res:=0
	for index<len(s){
		if s[index]=='X'{
			index+=3
			res+=1
		}else {
			index++
		}
	}

	return res
}
func main()  {
	fmt.Println(minimumMoves("OOOO"))
}