package main

import (
	"fmt"
	"sort"
)

// huffmanTree 带权路径最小的树，最优二叉树
// 生成的树原数组的节点都在叶子上

type Node struct {
	num   int
	left  *Node
	right *Node
}

func NewNode(num int) *Node {
	return &Node{num: num}
}

func (n *Node) prefixWalk() {
	fmt.Println(n.num)
	if n.left != nil {
		n.left.prefixWalk()
	}
	if n.right != nil {
		n.right.prefixWalk()
	}
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

func (h *HuffManTree) GenHuffManTree(arr []int) *Node {
	var list []*Node
	for _, v := range arr {
		list = append(list, NewNode(v))
	}

	for len(list) > 1 {
		// 1 从小到大排序
		sort.Sort(ByNum(list)) // 需要类型转换
		// 2 最小的两个元素构建一颗二叉树
		tmpRoot := NewNode(list[0].num + list[1].num)
		tmpRoot.left = list[0]
		tmpRoot.right = list[1]
		// 3 把生成的二叉树的根节点放回list，并删除list中这两个叶子节点对应的数据
		list = append(list, tmpRoot)
		list = list[2:]
	}
	// 4 最后list中只剩一个节点，就是huffman树的根节点
	return list[0]
}

func (h *HuffManTree) PreFixWalk(root *Node) {
	root.prefixWalk()
}

func main() {
	arr := []int{13, 7, 8, 3, 29, 6, 1}
	ht := HuffManTree{}
	root := ht.GenHuffManTree(arr)
	ht.PreFixWalk(root)

}
