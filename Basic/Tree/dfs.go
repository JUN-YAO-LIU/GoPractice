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
        // fmt.Println("----indert")
        return &Node{Key: key, Left: nil, Right: nil}
    }

    if key < root.Key {
        // fmt.Println("----left")
        root.Left = Insert(root.Left, key)
    } else if key > root.Key {
        // fmt.Println("----right")
        root.Right = Insert(root.Right, key)
    }

    return root
}

func dfs(root *Node,p *int, count int) {

    if root != nil {
		dfs(root.Left, p,count+1)
		dfs(root.Right,p ,count+1)
	}
    
	if count > *p {
		*p = count
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
    Insert(root, 10)


    var p *int
    c := 0
    p = &c
    dfs(root,p,0)
    fmt.Println(c)
}
