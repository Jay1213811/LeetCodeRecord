package main
func toLowerCase(s string) string {
	byteS:=[]byte(s)
	res:=make([]byte,0)
	for _,ch:=range byteS{
		if ch>='A' && ch <= 'Z'{
			ch=ch+32
			res=append(res,byte(ch))
		}else{
			res=append(res,byte(ch))
		}
	}
	return string(res)
}