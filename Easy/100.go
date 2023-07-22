func isSameTree(p *TreeNode, q *TreeNode) bool {

    // 新思維，如果都對，那可以比到最後一個是null
    if p == nil && q == nil {
        return true
    }

    if (p == nil) != (q == nil){
        return false
    }

    if p.Val != q.Val {
        return false
    }

    // 比左邊，比右邊
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}