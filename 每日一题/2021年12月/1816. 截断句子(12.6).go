package main
func truncateSentence(s string, k int) string {
	counter:=0
	for i:=0;i<len(s);i++{
		if s[i]==' '{
			counter++
			if counter==k{
				return s[:i]
			}
		}
	}
	return s
}