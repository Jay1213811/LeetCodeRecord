package main

import (
	"fmt"
	"math"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var wordMap map[rune]int /*创建集合 */
	wordMap = make(map[rune]int)
	k:=0
	for i := 'a'; i <= 'j'; i++ {
		wordMap[i]=k
		k++
	}
	result1:=[]int{}
	result2:=[]int{}
	result3:=[]int{}
	var oneSum,twoSum,threeSum int = 0,0,0
	for i:=0;i<len(firstWord);i++{
		result1=append(result1,int(wordMap[rune(firstWord[i])]))
	}
	for i:=0;i<len(secondWord);i++{
		result2=append(result2,int(wordMap[rune(secondWord[i])]))

	}
	for j:=0;j<len(targetWord);j++{
		result3=append(result3,int(wordMap[rune(targetWord[j])]))
	}
	length1:=len(result1)
	length2:=len(result2)
	length3:=len(result3)
	for i:=0;i<len(result1);i++{

		oneSum+=result1[i]*int(math.Pow10(length1))
		fmt.Println(oneSum)
		length1--
	}

	for i:=0;i<len(result2);i++{
		twoSum+=result2[i]*int(math.Pow10(length2))
		length2--
	}
	for i:=0;i<len(result3);i++{

		threeSum+=result3[i]*int(math.Pow10(length3))
		length3--
	}
	//fmt.Println(oneSum,twoSum,threeSum)
	if oneSum+twoSum==threeSum{
		return true
	}else {
		return false
	}
}
func main()  {
	fmt.Println(isSumEqual("acb","cba","cdb"))
	//fmt.Println(isSumEqual("i","j","bi"))
}
