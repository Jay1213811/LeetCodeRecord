package main

import "fmt"

func leastMinutes(n int) int {
	k := 1
	ans := 0
	for k < n {
		k*=2
		ans++

	}
	return ans+1
}
func main(){
	fmt.Println(leastMinutes(7))	//预期4
}