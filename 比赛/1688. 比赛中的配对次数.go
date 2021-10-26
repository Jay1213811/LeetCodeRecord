package main

import "fmt"

func numberOfMatches(n int) int {
	res:=0
	for n>1{
		if n%2==0 {
			n=n/2

			res+=n

		}else if n%2!=0 {
			res+=((n - 1) / 2)
			n=n-((n - 1) / 2)

		}
	}
	return res
}
func main()  {
	fmt.Println(numberOfMatches(2))
	fmt.Println(numberOfMatches(7))
	fmt.Println(numberOfMatches(14))
}