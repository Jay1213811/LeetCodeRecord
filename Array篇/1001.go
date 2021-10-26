package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var count int=0
	for n!=1{
		if n%2==0{
			n=n/2
		}else {
			n=(3*n+1)/2
		}
		count++
	}
	fmt.Printf("%d",count)
}
