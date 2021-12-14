package main

import (
	"math"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	m := map[int]int{}
	for i := 0; i < len(licensePlate); i++ {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			m[int(licensePlate[i] - 'a')]++
		}
		if licensePlate[i] >= 'A' && licensePlate[i] <= 'Z' {
			m[int(licensePlate[i] - 'A')]++
		}
	}
	res := ""
	minL := math.MaxInt32
	for i := 0; i < len(words); i++ {
		if check(m, words[i]) {
			if minL > len(words[i]) {
				minL = len(words[i])
				res = words[i]
			}
		}
	}
	return res
}

func check(m map[int]int, s string) bool {
	sm := map[int]int{}
	for i := 0; i < len(s); i++ {
		sm[int(s[i] - 'a')]++
	}
	for k, _ := range m {
		if v, ok := sm[k]; ok {
			if v < m[k] {
				return false
			}
		}   else {
			return false
		}
	}
	return true
}


