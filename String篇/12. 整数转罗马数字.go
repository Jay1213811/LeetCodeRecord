package main

import "fmt"

func intToRoman(num int) string {
	res:=[]byte{}
	LuoMaMap:=map[string]int{}
	LuoMaMap["I"]=1
	LuoMaMap["V"]=5
	LuoMaMap["X"]=10
	LuoMaMap["L"]=50
	LuoMaMap["C"]=100
	LuoMaMap["D"]=500
	LuoMaMap["M"]=1000
	LuoMaMap["IV"]=14
	LuoMaMap["IX"]=9
	LuoMaMap["XL"]=40
	LuoMaMap["XC"]=90
	LuoMaMap["CD"]=400
	LuoMaMap["CM"]=900
	for stringLuoMa,Number:=range LuoMaMap{
		for num>=Number{
			num-=Number
			byteLuoMa:=[]byte(stringLuoMa)
			res=append(res,byteLuoMa...)
		}
		if num==0{
			break
		}

	}
	return string(res)


}
func main()  {
	fmt.Println(intToRoman(3))
}