package main

import "fmt"

func isHappy(n int) bool {
	var slow int=n
	var fast int=next(n)
	for fast != 1 && fast != slow {
		slow = next(slow)
		fast = next(next(fast))
	}

	// fmt.Println(fast)
	if fast == 1 {
		return true
	}
	return false

}
func next(n int) int{
	var sum int=0
	for n!=0{
		temp:=n%10
		sum+=temp*temp
		n=n/10
	}
	return sum
}
func main()  {
	fmt.Println(isHappy(19))
}