package main

import "fmt"

func generateParenthesis(n int) []string {
	res := new([]string)
	back(n,n,"",res)
	return *res
}
func back(left,right int,tmp string,res *[]string)  {
	if right==0{
		*res=append(*res,tmp)
		return
	}
	if left>0{
		back(left-1,right,tmp+"(",res)
	}
	if right>left{
		back(left,right-1,tmp+")",res)
	}
}
func main()  {
	fmt.Println(generateParenthesis(3))
}