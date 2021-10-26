package main

import "fmt"

// Definition for a binary tree node.
 type TreeNode struct {
    Val int
    Left *TreeNode
  	Right *TreeNode
}

func numColor(root *TreeNode) int {
	vals:=[]int{}
	vals=preorderTraversal(root)
	Rm_duplicate(vals)
	return len(vals)

}
func Rm_duplicate(list []int) []int {
	var x []int = []int{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}
func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
func main()  {
	var node = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	fmt.Println(numColor(node))
}