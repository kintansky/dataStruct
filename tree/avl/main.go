package main

import (
	"fmt"
	"math"
)

// avl平衡二叉树是一个bst树，但是通过算法把bst树优化成左右子树高度差绝对值<=1
// 以达到左右子树平衡，优化查找速度

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(i int) *Node {
	return &Node{data: i}
}

// Height 返回当前节点为根节点的子树的高度
func (n *Node) Height() (h int) {
	if n == nil {
		return
	}
	return int(math.Max(float64(n.left.Height()), float64(n.right.Height()))) + 1
}

type AvlTree struct {
	root *Node
}

func (a *AvlTree) insertNode(n *Node, pos *Node) {
	if n.data < pos.data {
		if pos.left == nil {
			pos.left = n
			return
		}
		a.insertNode(n, pos.left)
	} else if n.data > pos.data {
		if pos.right == nil {
			pos.right = n
			return
		}
		a.insertNode(n, pos.right)
	} else {
		fmt.Printf("node %v exists\n", n.data)
	}
}

func (a *AvlTree) InsertNode(n *Node) {
	if a.root == nil {
		a.root = n
		return
	}
	a.insertNode(n, a.root)
}

func (a *AvlTree) Walk() {
	if a.root == nil {
		return
	}
	a.walk(a.root)
}

// BST 的中序遍历就是从小到大排列
func (a *AvlTree) walk(n *Node) {
	if n.left != nil {
		a.walk(n.left)
	}
	fmt.Println(n.data)
	if n.right != nil {
		a.walk(n.right)
	}
}

func main() {
	arr := []int{4, 3, 6, 5, 7, 8}
	avl := AvlTree{}
	for _, d := range arr {
		avl.InsertNode(NewNode(d))
	}
	avl.Walk()
	h := avl.root.Height()
	fmt.Println(h)
}
