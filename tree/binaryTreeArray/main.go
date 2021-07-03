package main

import "fmt"

// 顺序存储的二叉树，通过数组存储二叉树，主要应用在堆排序中
// 通常是一个完全二叉树
// arr中下标i的元素的左节点下标是2*i+1
// arr中下标i的元素的右节点下标是2*i+2
// arr中下标i的元素的父节点是(i-1)/2

type ArrBinaryTree struct {
	arr []int
}

func NewArrBinaryTree(arr []int) *ArrBinaryTree {
	return &ArrBinaryTree{arr: arr}
}

func (a *ArrBinaryTree) prefixWalk(idx int) {
	if len(a.arr) == 0 {
		fmt.Println("empty")
		return
	}
	fmt.Println(a.arr[idx])
	if idx*2+1 < len(a.arr) {
		a.prefixWalk(idx*2 + 1)
	}
	if idx*2+2 < len(a.arr) {
		a.prefixWalk(idx*2 + 2)
	}
}

func (a *ArrBinaryTree) PrefixWalk() {
	a.prefixWalk(0)
}

func (a *ArrBinaryTree) middleWalk(idx int) {
	if len(a.arr) == 0 {
		fmt.Println("empty")
		return
	}
	if idx*2+1 < len(a.arr) {
		a.middleWalk(idx*2 + 1)
	}
	fmt.Println(a.arr[idx])
	if idx*2+2 < len(a.arr) {
		a.middleWalk(idx*2 + 2)
	}
}

func (a *ArrBinaryTree) MiddleWalk() {
	a.middleWalk(0)
}

func (a *ArrBinaryTree) suffixWalk(idx int) {
	if len(a.arr) == 0 {
		fmt.Println("empty")
		return
	}
	if idx*2+1 < len(a.arr) {
		a.suffixWalk(idx*2 + 1)
	}
	if idx*2+2 < len(a.arr) {
		a.suffixWalk(idx*2 + 2)
	}
	fmt.Println(a.arr[idx])
}

func (a *ArrBinaryTree) SuffixWalk() {
	a.suffixWalk(0)
}

func main() {
	// a := NewArrBinaryTree([]int{1, 2, 3, 4, 5, 6, 7})
	// fmt.Println(a)
	// a.PrefixWalk()
	// fmt.Println("+++++++++++++")
	// a.MiddleWalk()
	// fmt.Println("+++++++++++++")
	// a.SuffixWalk()
	root := &TreeNode{Val: 4}
	nodem7 := &TreeNode{Val: -7}
	nodem3 := &TreeNode{Val: -3}
	root.Left = nodem7
	root.Right = nodem3
	nodem9 := &TreeNode{Val: -9}
	nodem3.Left = nodem9
	nodem32 := &TreeNode{Val: -3}
	nodem3.Right = nodem32
	node9 := &TreeNode{Val: 9}
	nodem9.Left = node9
	nodem72 := &TreeNode{Val: -7}
	nodem9.Right = nodem72
	nodem4 := &TreeNode{Val: -4}
	nodem32.Left = nodem4
	node6 := &TreeNode{Val: 6}
	node9.Left = node6
	node0 := &TreeNode{Val: 0}
	node6.Left = node0
	node62 := &TreeNode{Val: 6}
	node6.Right = node62
	nodem6 := &TreeNode{Val: -6}
	nodem72.Left = nodem6
	nodem62 := &TreeNode{Val: -6}
	nodem72.Right = nodem62
	nodem1 := &TreeNode{Val: -1}
	node0.Right = nodem1
	nodem42 := &TreeNode{Val: -4}
	node62.Left = nodem42

	prefixWalkFromRoot(root)
	c := Constructor()
	arr := c.serialize(root)
	fmt.Println(arr)
	org := c.deserialize(arr)
	prefixWalkFromRoot(org)

}
