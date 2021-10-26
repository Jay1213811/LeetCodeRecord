package main

import "fmt"

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
func letterCombinations(digits string) []string {

	res:=[]string{}//存储最终的结果
	//特殊情况
	if len(digits)==0{
		return []string{}
	}
	var dfs func(pathS []byte,index int)
	dfs=func(pathS []byte,index int){
		if index==len(digits){
			//终止
			tempS:=string(pathS)
			res=append(res, tempS)
			return
		}
		var combine string
		temp:=digits[index]	//digits中的每一个数字2 3
		combine=phoneMap[string(temp)]    //每个数字可能的字母
		byteCombine:=[]byte(combine)	//换成byte类型，遍历添加
		for i:=0;i<len(byteCombine);i++{
			//遍历叶子节点
			pathS=append(pathS,byteCombine[i])
			dfs(pathS,index+1)
			pathS=pathS[:len(pathS)-1]

		}


	}
	dfs([]byte{},0)
	return res

}
func main()  {
	fmt.Println(letterCombinations("23"))
}