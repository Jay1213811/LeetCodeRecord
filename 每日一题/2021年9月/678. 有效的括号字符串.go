package main

import "fmt"

/*
情况1,*为空 情况2,*为) 情况3,*为(
    还有特殊情况 (44个  )45个 *10个
*/


func checkValidString(s string) bool {
	leftCount, rightCount := 0, 0
	for _, c := range []byte(s) {
		if c == ')' {
			rightCount++
		} else {
			leftCount++
		}
		if rightCount > leftCount {
			return false
		}
	}
	leftCount, rightCount = 0, 0
	for i := len(s) -1 ; i >= 0; i-- {
		c := s[i]
		if c == '(' {
			leftCount++
		} else {
			rightCount++
		}
		if leftCount > rightCount {
			return false
		}
	}

	return true
}

func main()  {
	//fmt.Println(checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"))
	//fmt.Println(compareVersion("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"))
	fmt.Println(checkValidString("(*)"))
}