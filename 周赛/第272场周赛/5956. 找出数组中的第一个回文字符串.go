package 第272场周赛
func firstPalindrome(words []string) string {
	for _,v:=range words{
		if judge2(v)==true{
			return v
		}
	}
	return ""
}
func judge2(word string)bool{
	left,right:=0,len(word)-1
	for left<right{
		if word[left]==word[right]{
			left++
			right--
		}else{
			return false
		}
	}
	return true
}