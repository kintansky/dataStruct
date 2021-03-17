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

// RotateLeft 左旋
func (n *Node) RotateLeft() {
	// 1 以当前根的值复制出一个节点
	newNode := NewNode(n.data)
	// 2 这个复制节点的左子树指向原根的左子树
	newNode.left = n.left
	// 3 这个复制节点的右子树指向原根的右子树的左子树
	newNode.right = n.right.left
	// 4 替换原根节点的值为右节点的值
	n.data = n.right.data
	// 5 原根的右子树指向原根的右子树的右子树
	n.right = n.right.right
	// 6 原根的左子树指向复制节点
	n.left = newNode
}

func (n *Node) RotateRight() {
	newNode := NewNode(n.data)
	newNode.right = n.right
	newNode.left = n.left.right
	n.data = n.left.data
	n.left = n.left.left
	n.right = newNode
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
	a.Rotate()
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

func (a *AvlTree) Rotate() {
	if a.root.right.Height()-a.root.left.Height() > 1 {
		if a.root.right != nil && a.root.right.left.Height() > a.root.left.Height() {
			a.root.right.RotateRight()
			a.root.RotateLeft()
		} else {
			a.root.RotateLeft()
		}
		fmt.Printf("root:%d, height:%d, leftTreeHeight:%d, rightTreeHeight:%d\n", a.root.data, a.root.Height(), a.root.left.Height(), a.root.right.Height())
		return
	}
	if a.root.right.Height()-a.root.left.Height() < -1 {
		if a.root.left != nil && a.root.left.right.Height() > a.root.right.Height() {
			a.root.left.RotateLeft()
			a.root.RotateRight()
		} else {
			a.root.RotateRight()
		}
		fmt.Printf("root:%d, height:%d, leftTreeHeight:%d, rightTreeHeight:%d\n", a.root.data, a.root.Height(), a.root.left.Height(), a.root.right.Height())
		return
	}
}

func main() {
	// arr := []int{4, 3, 6, 5, 7, 8}
	arr := []int{10, 11, 7, 6, 8, 9}
	avl := AvlTree{}
	for _, d := range arr {
		avl.InsertNode(NewNode(d))
	}
	avl.Walk()
	h := avl.root.Height()
	fmt.Printf("root:%d, height:%d, leftTreeHeight:%d, rightTreeHeight:%d", avl.root.data, h, avl.root.left.Height(), avl.root.right.Height())
}
