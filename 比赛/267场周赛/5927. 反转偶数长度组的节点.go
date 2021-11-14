package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}


func reverseEvenLengthGroups(head *ListNode)*ListNode {
	headArray:=[]int{}
	for head!=nil{
		headArray=append(headArray,head.Val)
		head=head.Next
	}
	fmt.Println(headArray)
	i,group:=0,0
	for i<len(headArray){
		//组别
		group=group+1
		//偶数组肯定是偶数长度
		//如果到了末尾要判断最后一组长度为奇数还是偶数
		if (i+group>len(headArray)-1){
			if ((len(headArray)-i)%2==0){
				headArray=reverseArray(headArray,i,len(headArray)-1)
				break
			}else{
				break
			}
		}
		if group%2==0{
			headArray=reverseArray(headArray,i,i+group-1)
			i=i+group
		}else{
			i=i+group
		}

	}
	return makeListNode(headArray)
}
func reverseArray(nums []int,left,right int)[]int{
	for left<right{
		nums[left],nums[right]=nums[right],nums[left]
		left++
		right--
	}
	return nums
}
//数组转链表
func makeListNode(elements []int) *ListNode {
	if len(elements) == 0 {
		return nil
	}
	res := &ListNode{
		Val: elements[0],
	}
	// temp对res复制，二者没有关系了
	temp := res
	for i := 1; i < len(elements); i++ {
		// 对temp的子属性赋值
		temp.Next = &ListNode{Val: elements[i]}
		// 再将temp指向当前子属性 作为下一轮循环的当前值
		temp = temp.Next
	}
	return res
}
