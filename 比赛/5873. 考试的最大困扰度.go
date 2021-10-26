package main

import "fmt"

func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	find := func(target byte) int {
		res := 0
		cnt := 0

		for i, j := 0, 0; j < n; j++ {
			if answerKey[j] == target { cnt++ }
			for cnt > k {
				if answerKey[i] == target { cnt-- }
				i++
			}
			res = max(res, j-i+1)
		}
		return res
	}
	return max(find('T'), find('F'))
}

func max(i, j int) int {
	if i > j { return i }
	return j
}


func main()  {
	fmt.Println(maxConsecutiveAnswers("TTFF",2))
}
