package main

import (
	"fmt"
	"sort"
)

// huffman编码是一个前缀编码
// 字符出现的次数为权值，生成huffman树
// 到达每个叶子节点的前缀遍历路径就是对应字符的huffman编码，不通的huffman树生成的编码表都不一样

type Node struct {
	data   byte
	weight int
	left   *Node
	right  *Node
}

func NewNode(ch byte, w int) *Node {
	return &Node{
		data:   ch,
		weight: w,
	}
}

func (n *Node) Show() {
	fmt.Printf("(%s:weight:%d)", string(n.data), n.weight)
}

func (n *Node) prefixWalk() {
	n.Show()
	if n.left != nil {
		n.left.prefixWalk()
	}
	if n.right != nil {
		n.right.prefixWalk()
	}
}

type ByWeight []*Node

func (b ByWeight) Len() int {
	return len(b)
}

func (b ByWeight) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByWeight) Less(i, j int) bool {
	return b[i].weight < b[j].weight
}

type huffman struct {
	root      *Node
	bytesArr  []byte
	encodeMap map[byte]string
}

func NewHuffman(s string) *huffman {
	return &huffman{bytesArr: []byte(s), encodeMap: make(map[byte]string)}
}

func (h *huffman) createNodeArr() []*Node {
	var res []*Node
	cntMap := make(map[byte]int)
	for _, b := range h.bytesArr {
		_, exists := cntMap[b]
		if !exists {
			cntMap[b] = 0
		}
		cntMap[b]++
	}
	for b, w := range cntMap {
		res = append(res, NewNode(b, w))
	}
	return res
}

func (h *huffman) CreateTree() {
	// 1 将字符串转换成Node列表
	arr := h.createNodeArr()
	// 2 生成huffman树
	for len(arr) > 1 {
		sort.Sort(ByWeight(arr))
		tmpRoot := NewNode(byte(0), arr[0].weight+arr[1].weight) // 非叶子节点data不存放数据
		tmpRoot.left = arr[0]
		tmpRoot.right = arr[1]
		arr = append(arr, tmpRoot)
		arr = arr[2:]
	}
	h.root = arr[0]
}

func (h *huffman) GenEncodeMap(node *Node, pathStr string) {
	// 前序遍历
	// 通过路径生成huffman编码：向左为0，向右为1
	if node.data != byte(0) {
		h.encodeMap[node.data] = pathStr
	}
	if node.left != nil {
		pathStr += "0"
		h.GenEncodeMap(node.left, pathStr)
	}
	if node.right != nil {
		pathStr += "1"
		h.GenEncodeMap(node.right, pathStr)
	}
}

func (h *huffman) PrefixWalk(root *Node) {
	root.prefixWalk()
}

func main() {
	hf := NewHuffman("test test, no answer")
	hf.CreateTree()
	// hf.PrefixWalk(root)
	hf.GenEncodeMap(hf.root, "")
	fmt.Println(hf.encodeMap)
}
