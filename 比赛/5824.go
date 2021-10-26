package main

import (
	"fmt"
)
/*
因为3并没有替换
所以在3的时候程序已经结束了，输出结果。
*/
func maximumNumber(num string, change []int) string {
	s := []byte(num)
	i, n := 0, len(num)
	// 从高到低，找到第一个映射后比现在大的位置 开始替换
	for ; i < n; i++ {
		old := int(s[i] - '0')
		new := change[old]
		if old < new {
			break
		}
	}
	// 从高到底，开始映射，直到映射后比现在小
	for ; i < n; i++ {
		old := int(s[i] - '0')
		new := change[old]
		if old > new {
			//不替换了
			break
		}
		s[i] = '0' + byte(new)

	}
	return string(s)
}


func main()  {
	change:=[]int{9,8,5,0,3,6,4,2,6,8}
	//fmt.Println(sum(change))
	fmt.Println(maximumNumber("132",change))
}