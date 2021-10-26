package main

import (
	"fmt"
	"math"
)

func closestDivisors(num int) []int {
	ans := []int{0, int(1e9)}
	for i := num+1; i < num+3; i++ {
		cur :=  divide(i)
		if abs(cur[0]-cur[1])<abs(ans[0]-ans[1]) {
			ans = cur
		}
	}
	return ans
}

func divide(n int) []int {
	for i := int(math.Sqrt(float64(n))); i >= 0; i-- {
		if n % i == 0 {
			return []int{i, n/i}
		}
	}
	return []int{}
}


func abs(a int)int  {
	if a>0{
		return a
	}else{
		return -a
	}
}

func main()  {
	fmt.Println(closestDivisors(123))
}