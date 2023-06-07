package main

import(
	"fmt"
)

type Node struct{
	val int
	left *Node
	right *Node
}

func Insert(node *Node,val int) *Node{
	if node == nil{
		// if val == nil{
		// 	return &Node{val:nil,left:nil,right:nil}
		// }
		return &Node{val:val,left:nil,right:nil}
	}

	if node.left == nil{
		node.left = Insert(node.left,val)
	}else{
		node.right = Insert(node.right,val)
	}

	return node
}

func InOrder(root *Node) {
    if root != nil {
        InOrder(root.left)
        fmt.Printf("%d ", root.val)
        InOrder(root.right)
    }
}

func main(){
	var node *Node

	// 問題不能輸入nil
	// valid
	// var nodeArr = []int{3,5,1,6,2,0,8,7,4}
	// valid
	// var nodeArr1 []int

	node = Insert(node,3)
	Insert(node,5)
	Insert(node,1)
	Insert(node,6)
	Insert(node,2)
	Insert(node,0)
	// Insert(node,nil)
	// Insert(node,nil)
	Insert(node,7)
	Insert(node,4)
	
	// Printf 應該是字串?
	fmt.Println(node)

	InOrder(node)
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // Ancestor 祖先
 // descendant 後裔
 // lowest common ancestor 最低共同祖先
 // root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 // 注意parameter方式 : root p q 都是TreeNode
 // func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  
// }