package main

import(
	"fmt"
)


// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func verticalOrder(root *TreeNode) [][]int {
   if root == nil {
       return [][]int{}
   }

   var lw, rw int
   getWidth(root, &lw, &rw, 0)

   if lw < 0 {
       lw *= -1
   }

   columns := make([][]int, lw + rw + 1)

   for i := 0; i < len(columns); i++ {
       columns[i] = []int{}
   }

   nq := []*TreeNode{root}
   cq := []int{lw + 1}

   for len(nq) > 0 {
       n := nq[0]
       nq = nq[1:]
       c := cq[0]
       cq = cq[1:]

       if n == nil {
           continue
       }

       columns[c - 1] = append(columns[c - 1], n.Val)
       nq = append(nq, n.Left, n.Right)
       cq = append(cq, c - 1, c + 1)
   }

   return columns
}

func getWidth(node *TreeNode, left *int, right *int, width int) {
   if node == nil {
       return
   }

   if width < *left {
       *left = width
   }

   if width > *right {
       *right = width
   }

   getWidth(node.Left, left, right, width - 1)
   getWidth(node.Right, left, right, width + 1)
}

