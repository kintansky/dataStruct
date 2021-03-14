package main

// 二叉排序树：左节点比根小，右节点比根大
// 中序遍历结果就是从小到大排列

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(i int) *Node {
	return &Node{data: i}
}

type BST struct {
	root *Node
}

func (b *BST) insertNode(n *Node, pos *Node) {
	if n.data < pos.data {
		if pos.left == nil {
			pos.left = n
			return
		}
		b.insertNode(n, pos.left)
	} else if n.data > pos.data {
		if pos.right == nil {
			pos.right = n
			return
		}
		b.insertNode(n, pos.right)
	} else {
		fmt.Printf("node %v exists\n", n.data)
	}
}

func (b *BST) InsertNode(n *Node) {
	if b.root == nil {
		b.root = n
		return
	}
	b.insertNode(n, b.root)
}

func (b *BST) Walk() {
	if b.root == nil {
		return
	}
	b.walk(b.root)
}

// BST 的中序遍历就是从小到大排列
func (b *BST) walk(n *Node) {
	if n.left != nil {
		b.walk(n.left)
	}
	fmt.Println(n.data)
	if n.right != nil {
		b.walk(n.right)
	}
}

// 非递归查找
func (b *BST) FindNode(i int) (n *Node, pn *Node) {
	tmp := b.root
	if tmp.data == i {
		n = tmp
		return
	}
	for tmp != nil {
		if i < tmp.data {
			if tmp.left == nil {
				return
			}
			if tmp.left.data == i {
				n = tmp.left
				pn = tmp
				return
			}
			tmp = tmp.left
		} else if i > tmp.data {
			if tmp.right == nil {
				return
			}
			if tmp.right.data == i {
				n = tmp.right
				pn = tmp
				return
			}
			tmp = tmp.right
		}
	}
	return
}

// 递归查找
func (b *BST) findNode2(i int, node *Node) (n *Node, pn *Node) {
	if i < node.data {
		if node.left == nil {
			return
		}
		if node.left.data == i {
			n = node.left
			pn = node
			return
		}
		return b.findNode2(i, node.left)
	} else if i > node.data {
		if node.right == nil {
			return
		}
		if node.right.data == i {
			n = node.right
			pn = node
			return
		}
		return b.findNode2(i, node.right)
	}
	return
}

func (b *BST) FindNode2(i int) (n *Node, pn *Node) {
	if b.root.data == i {
		n = b.root
		return
	}
	return b.findNode2(i, b.root)
}

func main() {
	arr := []int{7, 3, 10, 12, 5, 1, 9}
	bst := BST{}
	for _, d := range arr {
		bst.InsertNode(NewNode(d))
	}
	bst.Walk()
	node, parentNode := bst.FindNode(9)
	fmt.Println(node, parentNode)
	node, parentNode = bst.FindNode2(9)
	fmt.Println(node, parentNode)
}
