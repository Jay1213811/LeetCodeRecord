package main

import "fmt"

func checkZeroOnes(s string) bool {
	oneIndex,zeroIndex:=0,0
	var maxCountZero,maxCountOne int =0,0
	for i:=0;i<len(s);{
		//如果等于0
		if s[i]==48{
			oneIndex=i+1
			i++
			maxCountZero=max3(maxCountZero,i-zeroIndex)
		}else{
			zeroIndex=i+1
			i++
			maxCountOne=max3(maxCountOne,i-oneIndex)
		}

	}
	if maxCountOne>maxCountZero{
		return true
	}else {
		return false
	}

}
func max3(a,b int)int  {
	if a>b{
		return a
	}else {
		return b
	}
}
func main()  {
	fmt.Println(checkZeroOnes("1"))
}