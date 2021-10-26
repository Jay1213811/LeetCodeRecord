package main

import (
	"fmt"
	"strings"
)

func maskPII(s string) string {
	//情况1 为邮箱
	needs:=strings.ToLower(s)	//把大写转换为小写
	temp:=[]byte{}
	end:=""
	if judgeEmail(s)==true{
		for i:=0;i<len(s);i++{
			if s[i]!='@'{
				temp=append(temp,needs[i])

			}else{
				end=needs[i:]
				break
			}
		}
	//存储了邮箱前的名字
	return string(temp[0])+"*****"+string(temp[len(temp)-1])+end
	}
	count:=getNum(s)
	//情况2 为电话号码
	switch count {
	case 13:return "+***-***-***-" + getendS(s,count)
	case 12:return "+**-***-***-" + getendS(s,count)
	case 11:return "+*-***-***-" + getendS(s,count)
	default: // 默认为10位
		return "***-***-" + getendS(s,count)
	}






}
//判断是否为本地号码 是的话返回true
func judgeLocal(s string)bool  {
	count:=0
	bytes:=[]byte(s)
	for i:=0;i<len(bytes);i++{
		if bytes[i]>='0' && bytes[i]<='9'{
			count++
		}
	}
	if count==10{
		return true
	}
	return false
}
//返回后四位数字
func getendS(s string,number int)string{
	count:=0
	bytes:=[]byte(s)
	temp:=[]byte{}
	for i:=0;i<len(s);i++ {
		for i := 0; i < len(bytes); i++ {
			if bytes[i] >= 48 && bytes[i] <= 57 {
				count++
				if count>6 && number==10{
					temp=append(temp,bytes[i])
					if len(temp)==4{
						return string(temp)
					}

				}else if count>7 && number==11{
					temp=append(temp,bytes[i])
					if len(temp)==4{
						return string(temp)
					}
				}else if count>8 && number==12{
					temp=append(temp,bytes[i])
					if len(temp)==4{
						return string(temp)
					}
				}else if count>9 && number==13{
					temp=append(temp,bytes[i])
					if len(temp)==4{
						return string(temp)
					}
				}
			}
		}
	}
	return string(temp)
}

func getNum(s string)int  {
	count:=0
	bytes:=[]byte(s)
	for i:=0;i<len(bytes);i++{
		if bytes[i]>='0' && bytes[i]<='9'{
			count++
		}
	}
	return count
}
func judgeEmail(s string)bool  {
	for i:=0;i<len(s);i++{
		if s[i]=='@'{
			return true
		}
	}
	return false
}
func main()  {
	fmt.Print(maskPII("+86(88)1513-7-74"))
}