package main

import "fmt"

/*
是正数插入第一个比x小的数字前面
是负数插入第一个比x大的数字前面
*/
func maxValue(n string, x int) string {
	//如果是负数找到第一个比x大的
	if n[0]=='-'{
		for i:=0;i<len(n);i++{

			if(n[i]>byte('0'+x)){
				return n[:i] + string('0'+x) + n[i:]
			}
		}
	}else {
		for i:=1;i<len(n);i++{
			//如果是正数找到第一个比x小的
			if(n[i]<byte('0'+x)){
				return n[:i] + string('0'+x) + n[i:]
			}
		}
	}
	return n+string('0'+x)
}

func main()  {
	fmt.Println(maxValue("-132",3))
}