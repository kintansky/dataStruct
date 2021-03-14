package main

import (
	"fmt"
	"sort"
	"strconv"
)

// huffman编码是一个前缀编码
// 字符出现的次数为权值，生成huffman树
// 到达每个叶子节点的前缀遍历路径就是对应字符的huffman编码，不通的huffman树生成的编码表都不一样

type Node struct {
	data   byte // 对应的字符
	weight int  // 出现的次数作为权值
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
	root          *Node           // huffman树的根节点
	encodeMap     map[byte]string // 生成的编码表
	huffmanStrLen int             // huffman编码后的二进制数据长度，用于解码
}

func NewHuffman() *huffman {
	return &huffman{encodeMap: make(map[byte]string)}
}

func (h *huffman) createNodeArr(s string) []*Node {
	var res []*Node
	cntMap := make(map[byte]int)
	for _, b := range []byte(s) {
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

// createTree 生成huffman树
func (h *huffman) createTree(s string) {
	// 1 将字符串转换成Node列表
	arr := h.createNodeArr(s)
	// 2 生成huffman树
	for len(arr) > 1 {
		sort.Sort(ByWeight(arr))
		tmpRoot := NewNode(byte(0), arr[0].weight+arr[1].weight) // 非叶子节点data不存放数据
		// tmpRoot := &Node{weight: arr[0].weight + arr[1].weight}
		tmpRoot.left = arr[0]
		tmpRoot.right = arr[1]
		arr = arr[2:]
		arr = append(arr, tmpRoot)
	}
	h.root = arr[0]
}

// genEncodeMap 按照路径生成字符编码
func (h *huffman) genEncodeMap(node *Node, code, pathStr string) {
	// 前序遍历
	// 通过路径生成huffman编码：向左为0，向右为1
	pathStr += code
	if node.data != byte(0) {
		h.encodeMap[node.data] = pathStr
		return
	}
	if node.left != nil {
		h.genEncodeMap(node.left, "0", pathStr)
	}
	if node.right != nil {
		h.genEncodeMap(node.right, "1", pathStr)
	}
}

func (h *huffman) PrefixWalk() {
	h.root.prefixWalk()
}

// archive 组合成对应的
func (h *huffman) archive(s string) (huffmanByte []byte) {
	var huffmanStr string
	for _, b := range []byte(s) {
		huffmanStr += h.encodeMap[b]
	}
	h.huffmanStrLen = len(huffmanStr)
	// fmt.Println(huffmanStr, len(huffmanStr))
	// 开始压缩
	length := (len(huffmanStr) + 7) / 8 // 向上取整
	huffmanByte = make([]byte, length)
	var tmp string
	for i := 0; i < length; i++ {
		if (i+1)*8 < len(huffmanStr) {
			tmp = huffmanStr[i*8 : (i+1)*8]
		} else {
			tmp = huffmanStr[i*8:]
		}
		b, _ := strconv.ParseInt(tmp, 2, 32)
		huffmanByte[i] = byte(b)
	}
	// fmt.Println(huffmanByte)
	return
}

// HuffmanArchive huffman编码
func (h *huffman) HuffmanArchive(s string) (huffmanByte []byte, length int) {
	h.createTree(s)
	h.genEncodeMap(h.root, "", "")
	huffmanByte = h.archive(s)
	length = h.huffmanStrLen
	return
}

// HuffmanUnArchive huffman解码
func (h *huffman) HuffmanUnArchive(huffmanByte []byte, encodeMap map[byte]string, orgLength int) (res string) {
	// 1 转二进制，拼接成str
	var huffmanStr string
	for i, b := range huffmanByte {
		if i == len(huffmanByte)-1 { // 最后一位是否补0，补多少个0，由原来的huffmanStr长度决定
			realCnt := fmt.Sprintf("%%%db", orgLength-i*8)
			huffmanStr += fmt.Sprintf(realCnt, b)
			continue
		}
		huffmanStr += fmt.Sprintf("%08b", b) // 其余字符都要补0
	}
	// fmt.Println(huffmanStr, len(huffmanStr))
	// 2 根据字典翻译出结果
	// 2.1 反转encodeMap 的 k，v
	decodeMap := make(map[string]byte)
	for k, v := range encodeMap {
		decodeMap[v] = k
	}
	// 2.2 逐个扫描huffmanStr，然后拼接回去
	var key []byte
	for _, b := range []byte(huffmanStr) {
		key = append(key, b)
		str, exists := decodeMap[string(key)]
		if exists {
			res += string(str)
			key = []byte{}
		}
	}
	return
}

func main() {
	s := "i like like like java do you like a java"
	hf := NewHuffman()
	huffmanByte, length := hf.HuffmanArchive(s)
	fmt.Println("after encode:", huffmanByte)
	res := hf.HuffmanUnArchive(huffmanByte, hf.encodeMap, length)
	fmt.Println(res)
}
