package main

import "fmt"

type node struct {
	key   int
	value string
	left  *node
	right *node
}

func NewNode(k int) *node {
	return &node{key: k, left: nil, right: nil}
}

func (n *node) setKey(k int) {
	n.key = k
}

func (n *node) setValue(v string) {
	n.value = v
}

func (n *node) setLeft(l *node) {
	n.left = l
}

func (n *node) setRight(r *node) {
	n.right = r
}

func (n *node) getKey() int {
	return n.key
}

func (n *node) getValue() string {
	return n.value
}

func (n *node) getLeft() *node {
	return n.left
}

func (n *node) getRight() *node {
	return n.right
}

func main() {
	// root node
	root := NewNode(8)

	// node
	n1 := NewNode(4)
	root.setLeft(n1)

	n2 := NewNode(11)
	root.setRight(n2)

	n3 := NewNode(2)
	n1.setLeft(n3)

	n4 := NewNode(6)
	n1.setRight(n4)

	n5 := NewNode(10)
	n2.setLeft(n5)

	n6 := NewNode(13)
	n2.setRight(n6)

	n7 := NewNode(5)
	n4.setLeft(n7)

	n8 := NewNode(7)
	n4.setRight(n8)

	key := 9
	root.insert(key)

	n := root.search(key)
	fmt.Printf("n.getKey() = %+v\n", n)
}

func (n *node) insert(key int) {
	if key > n.getKey() {
		if n.right == nil {
			n.setRight(NewNode(key))
		} else {
			n.getRight().insert(key)
		}
	} else {
		if n.left == nil {
			n.setLeft(NewNode(key))
		} else {
			n.getLeft().insert(key)
		}
	}
}

func (n *node) search(key int) *node {
	if n.getKey() == key {
		return n
	} else if key > n.getKey() {
		if n.getRight() == nil {
			return nil
		}
		return n.getRight().search(key)
	} else {
		if n.getLeft() == nil {
			return nil
		}
		return n.getLeft().search(key)
	}
}

// func insertNode(before *node, after *node, x int) *node {
//   var p *node
//   if after == nil {
//     n := NewNode(x)
//     p = &n
//     if x < before.getKey() {
//       before.setLeft(p)
//     } else {
//       before.setRight(p)
//     }
//     return p
//   } else if after.getKey() == x {
//     return after
//   } else {
//     before = after
//     if x < after.getKey() {
//       after.setLeft(insertNode(before, after.getLeft(), x))
//     } else {
//       after.setRight(insertNode(before, after.getRight(), x))
//     }
//     return after
//   }
// }

// func search(key int, node *node) *node {
//   for node != nil && key != node.getKey() {
//     if key < node.getKey() {
//       node = node.getLeft()
//     } else {
//       node = node.getRight()
//     }
//   }
//   return node
// }
