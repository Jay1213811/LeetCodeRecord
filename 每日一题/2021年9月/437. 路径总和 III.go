package main


type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	res:=0
	if root==nil{
		return 0
	}
	res+=pathSum(root,targetSum)
	res+=pathSum(root.Left,targetSum)
	res+=pathSum(root.Right,targetSum)
	return res

}

func rootSum(root *TreeNode,targetSum int)int  {
	res:=0
	if root==nil{
		return 0
	}
	val:=root.Val
	if val==targetSum{
		res++
	}
	res+=rootSum(root.Left,targetSum-val)
	res+=rootSum(root.Right,targetSum-val)
	return res
}
