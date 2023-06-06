package main

import (
	"fmt"
)

// BST node
type TreeNode struct{
	value int
	left *TreeNode
	right *TreeNode
}


// BST => Binary Search Tree
func main(){

	fmt.Println(1)
}

// 做成BST
func genBST(sArr []string) TreeNode{
	var node TreeNode

	for i := 0 ; i < len(sArr); i++{
		node.value = sArr[i]
	}

	return node
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var result int

	if root == nil{
		return 0
	}

	

    return result
}