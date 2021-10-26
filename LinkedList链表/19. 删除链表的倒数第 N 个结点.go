package main

type ListNode struct {
     Val int
     Next *ListNode
}
//方法1：快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy:=&ListNode{Next: head}
	low,fast:=dummy,dummy
	//fast先走n+1步，然后一起前进，当fast为nil时，low就来到了倒数第n+1个节点
	//(因为需要是倒数第n+1个节点，才方便删除倒数第n个节点)
	for i:=0;i<n+1;i++{
		fast=fast.Next
	}
	//两个人一起走
	for fast!=nil{	
		fast=fast.Next
		low=low.Next
	}
	low.Next=low.Next.Next
	return dummy.Next
}