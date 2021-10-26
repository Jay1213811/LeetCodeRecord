package main

import "fmt"



func lengthOfLongestSubstringTwoDistinct(s string) int {
    if len(s) < 3 {
        return len(s)
    }
    m := map[byte]int{}

    i, j := 0, 0
    maxLen := 1
    for i < len(s) && j < len(s) {
        setMap(s[j], m)
        if len(m) < 3 {
            maxLen = max2(j-i+1, maxLen)
        } else {
            for len(m) >= 3 {
                removeMapKey(s[i], m)
                i++
            }
        }
        j++
    }
    return maxLen
}

func setMap(element byte, m map[byte]int) {
    count, flag := m[element]
    if flag {
        m[element] = count + 1
    } else {
        m[element] = 1
    }
}

func max2(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func removeMapKey(element byte, m map[byte]int) {
    count, flag := m[element]
    if flag {
        if count == 1 {
            delete(m, element)
        } else {
            m[element] = count - 1
        }
    }
}


func main()  {
    fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
}