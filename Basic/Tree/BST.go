package main

import "fmt"

// 定义树节点结构体
type Node struct {
    Key   int
    Left  *Node
    Right *Node
}

// 插入节点
func Insert(root *Node, key int) *Node {
    if root == nil {
        fmt.Println("----indert")
        return &Node{Key: key, Left: nil, Right: nil}
    }

    if key < root.Key {
        fmt.Println("----left")
        root.Left = Insert(root.Left, key)
    } else if key > root.Key {
        fmt.Println("----right")
        root.Right = Insert(root.Right, key)
    }

    return root
}

// 删除节点
func Delete(root *Node, key int) *Node {
    if root == nil {
        return root
    }

    if key < root.Key {
        root.Left = Delete(root.Left, key)
    } else if key > root.Key {
        root.Right = Delete(root.Right, key)
    } else {
        // 找到待删除节点
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        }

        // 有两个子节点的情况
        minKey := findMinKey(root.Right)
        root.Key = minKey
        root.Right = Delete(root.Right, minKey)
    }

    return root
}

// 查找最小键值
func findMinKey(root *Node) int {
    for root.Left != nil {
        root = root.Left
    }
    return root.Key
}

// 查找节点
func Search(root *Node, key int) *Node {
    if root == nil || root.Key == key {
        return root
    }

    if key < root.Key {
        return Search(root.Left, key)
    }

    return Search(root.Right, key)
}

// 中序遍历
func InOrder(root *Node) {
    if root != nil {
        InOrder(root.Left)
        fmt.Printf("%d ", root.Key)
        InOrder(root.Right)
    }
}

func main() {
    var root *Node

    root = Insert(root, 50)
    Insert(root, 30)
    Insert(root, 20)
    Insert(root, 40)
    Insert(root, 70)
    Insert(root, 60)
    Insert(root, 80)

    fmt.Println("二叉搜索树中序遍历结果：")
    InOrder(root)
    fmt.Println()

    fmt.Println("搜索键值 60：")
    if Search(root, 60) != nil {
        fmt.Println("键值 60 存在")
    } else {
        fmt.Println("键值 60 不存在")
    }

    fmt.Println("删除键值 40：")
    root = Delete(root, 40)
    fmt.Println("二叉搜索树中序遍历结果：")
    InOrder(root)
    fmt.Println()
}
