package main

import "fmt"

// Node
type Node struct {
	val   int
	left  *Node
	right *Node
}

// Insert
func (n *Node) Insert(val int) {
	if val > n.val {
		if n.right == nil {
			n.right = &Node{val: val}
		} else {
			n.right.Insert(val)
		}
	} else if val < n.val {
		if n.left == nil {
			n.left = &Node{val: val}
		} else {
			n.left.Insert(val)
		}
	}
}

// Search
func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}

	if val > n.val {
		if n.right == nil {
			return false
		}

		return n.right.Search(val)
	} else if val < n.val {
		if n.left == nil {
			return false
		}

		return n.left.Search(val)
	}

	return true
}

func main() {
	t := &Node{val: 100}

	t.Insert(150)
	t.Insert(150)

	fmt.Println(t)
	t.Insert(4150)
	fmt.Println(t.right)

	t.Insert(10)
	fmt.Println(t)

	s := t.Search(4150)
	fmt.Println(s)
}
