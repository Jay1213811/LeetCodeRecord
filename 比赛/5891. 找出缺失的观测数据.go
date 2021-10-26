package main

import "fmt"

/*
首先计算rolls的总和
1+5+6=12	12/3=4=mean
(12+X)/n+len(rolls)=4
(12+X)/7=3
12+X=21
X=9	X是4个数的和
也就是1-9相加遍历n次然后得到总和为9
*/
func missingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	for _, v := range rolls {
		sum += v
	}
	miss := (n+len(rolls))*mean - sum
	if miss < n || miss > n*6 {
		return nil
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	miss -= n
	for i := range ans {
		if miss < 6 {
			ans[i] += miss
			break
		}
		ans[i] = 6
		miss -= 5
	}
	return ans
}


func main(){
	//rolls:=[]int{4,5,6,2,3,6,5,4,6,4,5,1,6,3,1,4,5,5,3,2,3,5,3,2,1,5,4,3,5,1,5}
	rolls2:=[]int{1,5,6}
	fmt.Println(missingRolls(rolls2,3,4))
	//fmt.Println(missingRolls(rolls,4,5))
}