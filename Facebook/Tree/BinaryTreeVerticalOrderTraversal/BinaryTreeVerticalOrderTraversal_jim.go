package main

import(
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func verticalOrder(root *TreeNode) [][]int {
    // layer * 4,layer * 5 的位置
    var result = make([][]int,0)

    result.Append(result,root.Val)
    verticalOrder(root.)
}