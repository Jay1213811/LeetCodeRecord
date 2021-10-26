package main
/*
建一个哈希表m，遍历数组，将数组中的值和出现的次数存到哈希表m中
对数组进行排序
遍历数组，将存在双倍的数num添加到返回数组ans中
当遍历到的数num为0时，判断0出现的个数是否为偶数，是则将其添加到返回数组ans中，添加次数为其出现次数的一半
当遍历到的数num在哈希表中存在 2 * num，且num的出现次数小于等于2 * num的出现次数时，将num添加到返回数组ans中
如果返回数组ans的长度正好等于原数组changed的一半时，返回ans，否则返回空数组
*/
import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	var ans []int
	n := len(changed)
	//如果数组changed的长度为奇数，返回空数组
	if n % 2 == 1{
		return ans
	}
	//将数组changed中的值和值出现的次数存入哈希表m中
	m := map[int]int{}
	for _, v := range changed{
		if cnt, ok := m[v]; ok{
			m[v] = cnt + 1
		}else{
			m[v] = 1
		}
	}
	//对数组changed排序，保证后面遍历添加元素时不会出问题
	sort.Ints(changed)
	for i := 0; i < n; i++{
		//如果该元素已经被访问过了，就跳过
		if i > 0 && i < n && changed[i] == changed[i-1]{
			continue
		}
		k := changed[i]
		//如果该元素出现次数为0，跳过
		if m[k] == 0{
			continue
		}
		//如果该元素是0
		if k == 0{
			if m[k] % 2 == 0{
				count := m[k] / 2
				for count > 0{
					count--
					ans = append(ans, k)
				}
			}
			continue
		}
		//如果该元素num在哈希表m中存在 2 * num
		if cnt, ok := m[k*2]; ok{
			count := m[k]
			if count > cnt{
				break
			}else{
				m[k] = 0
				m[k*2] = cnt - count
				for count > 0{
					count--
					ans = append(ans, k)
				}
			}
		}
	}
	if len(ans) != n / 2{
		return []int{}
	}else{
		return ans
	}
}

func main()  {
	changed :=[]int{1,3,4,2,6,8}
	fmt.Println(findOriginalArray(changed))
}