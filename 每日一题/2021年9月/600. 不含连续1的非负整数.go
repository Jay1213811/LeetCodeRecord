package main

import (
	"fmt"
	"strconv"
)
//思路。首先找出小于或等于 n 的非负整数，转换为二进制存储在数组中，最后遍历
func DecToBin2(n int) string{
	result := ""

	if n == 0 {
		return "0"
	}

	for ;n > 0;n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb)+result
	}

	return result
}
func BinToDec(bin string)  {
	for _,c:=range bin{
		stac
	}
}

func getSum2(a int, b int) int {
	return DecToBin2(a+b)
}
func main()  {
	fmt.Println(getSum2(1,2))
}

