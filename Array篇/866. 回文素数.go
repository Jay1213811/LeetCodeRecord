package main

import (
	"fmt"
	"math"
	"strconv"
)

func primePalindrome(n int) int {
	for true{
		if isHuiWen(n)==true && isPrime(n)==true{
			return n
		}else {
			n++
		}
	}
	return n

}
//判断素数
func isPrime(n int)bool{
	if n<=1{
		return false
	}
	sqrt:=int(math.Sqrt(float64(n)))
	for i:=2;i<=sqrt;i++{
		if n%i==0{
			return false
		}
	}
	return true
}
//判断是否为回文数
func isHuiWen(x int)  bool{
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		} }
	return true

}
func main()  {
	fmt.Println(primePalindrome(1))
}
