package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	X:=0
	for i:=0;i<len(operations);i++{
		if operations[i]=="++X" || operations[i]=="X++" {
			X=X+1
		}else if operations[i]=="X--"|| operations[i]=="X--"{
			X=X-1
		}
	}
	return X


}
func main()  {
	s:=[]string{"--X","X++","X++"}
	fmt.Println(finalValueAfterOperations(s))
}