package main

import "log"

// Node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	log.Println(tree)
	tree.Insert(50)
	log.Println(tree)
	tree.Insert(200)
	tree.Insert(250)
	log.Println(tree)
	log.Println("search 50", tree.Search(50))
	log.Println("search 88", tree.Search(88))
	log.Println("search 200", tree.Search(200))
	log.Println("search 300", tree.Search(300))
}
