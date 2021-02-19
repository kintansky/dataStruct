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
	a := NewArrBinaryTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(a)
	a.PrefixWalk()
	fmt.Println("+++++++++++++")
	a.MiddleWalk()
	fmt.Println("+++++++++++++")
	a.SuffixWalk()
}
