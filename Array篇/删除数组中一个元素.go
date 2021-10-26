package main

import "fmt"

func remove(slice []int, i int) []int {
	a:=slice[i:]
	b:=slice[i+1:]

	copy(a, b)

	fmt.Println("a=",a)
	fmt.Println("b=",b)

	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
}