package main

import (
	"fmt"
	"math"
)

/*给你一个二进制字符串 s 。如果字符串中由 1 组成的 最长 连续子字符串 严格长于 由 0 组成的 最长 连续子字符串，返回 true ；否则，返回 false 。

例如，s = "110100010" 中，由 1 组成的最长连续子字符串的长度是 2 ，由 0 组成的最长连续子字符串的长度是 3 。
注意，如果字符串中不存在 0 ，此时认为由 0 组成的最长连续子字符串的长度是 0 。字符串中不存在 1 的情况也适用此规则。

输入：s = "1101"
输出：true
解释：
由 1 组成的最长连续子字符串的长度是 2："1101"
由 0 组成的最长连续子字符串的长度是 1："1101"
由 1 组成的子字符串更长，故返回 true 。
*/
func checkZeroOnes(s string) bool {

	var zeroMax,oneMax int=0,0
	lcc:=[]byte(s)
	//从前往后遍历
	for index:=0;index<len(lcc);{
		//第一个数字为0
		if lcc[index]==48{
			lenOf0:=0
			//开始遍历连续的0的数量
			for index<len(lcc) && lcc[index]==48{
				lenOf0++
				index++
			}
			zeroMax= int(math.Max(float64(zeroMax), float64(lenOf0)))
		}else {
			lenOf1:=0
			for index<len(lcc) && lcc[index]==49{
				lenOf1++
				index++
			}
			oneMax=int(math.Max(float64(oneMax), float64(lenOf1)))
		}
	}
	if oneMax>zeroMax{
		return true
	}else {
		return false
	}
}
func main() {
	str := "011000111"

	fmt.Println(checkZeroOnes(str))
}