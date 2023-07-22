package main

import "fmt"

type Node struct {
    value int
    next  *Node
}

func reverseList(head *Node) *Node {
    var prev *Node

	// cache the current head
    current := head

	// use loop to change all nodes of direction
    for current != nil {
        next := current.next

		// because we reverse node's direction, current's next should be null.
        current.next = prev

		// 現在成為下一個節點的
        prev = current

		// 執行到下個節點了
        current = next
    }

    return prev
}

func printList(head *Node) {
    current := head

    for current != nil {
        fmt.Printf("%d ", current.value)
        current = current.next
    }

    fmt.Println()
}

func main() {
    // Create a sample linked list
    head := &Node{value: 1}
    node2 := &Node{value: 2}
    node3 := &Node{value: 3}
    node4 := &Node{value: 4}

    head.next = node2
    node2.next = node3
    node3.next = node4

    fmt.Println("Original list:")
    printList(head)

    // Reverse the linked list
    reversed := reverseList(head)

    fmt.Println("Reversed list:")
    printList(reversed)
}
