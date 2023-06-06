func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
			return 0
	}
	var sum = 0
	if root.Val >= low && root.Val <= high {
			sum += root.Val
	}

	l := rangeSumBST(root.Left, low, high)

	r := rangeSumBST(root.Right, low, high)
	return sum + l + r
}