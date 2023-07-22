// Convert Sorted Array to Binary Search Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	root := &TreeNode{Val: nums[i]}
	if i != 0 {
		root.Left = sortedArrayToBST(nums[:i])
	}
	if i != len(nums)-1 {
		root.Right = sortedArrayToBST(nums[i+1:])
	}
	return root
}


// Jim
func sortedArrayToBST(nums []int) *TreeNode {
	// 高度平衡，所以是左右數量盡可能平分

	if len(nums) <= 0{
		return nil
	}

	// 因為index是從0開始，所以剛好不用加1
	l := len(nums)/2
	root := &TreeNode{Val:nums[l],Left:nil,Right:nil}

	// 直接切一半
	root.Left = sortedArrayToBST(nums[:l])
	root.Right = sortedArrayToBST(nums[l+1:])
	
	return root
}