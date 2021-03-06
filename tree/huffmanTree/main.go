package main

import (
	"fmt"
	"sort"
)

type Node struct {
	num   int
	left  *Node
	right *Node
}

func NewNode(num int) *Node {
	return &Node{num: num}
}

type ByNum []*Node

func (b ByNum) Len() int {
	return len(b)
}

func (b ByNum) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByNum) Less(i, j int) bool {
	return b[i].num < b[j].num
}

type HuffManTree struct {
}

func (h *HuffManTree) GenHuffManTree(arr []int) {
	var list []*Node
	for _, v := range arr {
		list = append(list, NewNode(v))
	}
	// 从小到大排序
	sort.Sort(ByNum(list))
	for _, d := range list {
		fmt.Println(d.num)
	}
}

func main() {
	arr := []int{13, 7, 8, 3, 29, 6, 1}
	ht := HuffManTree{}
	ht.GenHuffManTree(arr)
}
