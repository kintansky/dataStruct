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

// 非递归查找节点和他的父节点，n 查找的节点，pn 查找节点的父节点，nodeType n是pn的left还是right节点，如果是root节点，nodeType=root
func (b *BST) FindNode(i int) (n *Node, pn *Node, nodeType string) {
	tmp := b.root
	if tmp.data == i {
		n = tmp
		nodeType = "root"
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
				nodeType = "left"
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
				nodeType = "right"
				return
			}
			tmp = tmp.right
		}
	}
	return
}

// 递归查找
func (b *BST) findNode2(i int, node *Node) (n *Node, pn *Node, nodeType string) {
	if i < node.data {
		if node.left == nil {
			return
		}
		if node.left.data == i {
			n = node.left
			pn = node
			nodeType = "left"
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
			nodeType = "right"
			return
		}
		return b.findNode2(i, node.right)
	}
	return
}

func (b *BST) FindNode2(i int) (n *Node, pn *Node, nodeType string) {
	if b.root.data == i {
		n = b.root
		nodeType = "root"
		return
	}
	return b.findNode2(i, b.root)
}

func (b *BST) DelNode(i int) {
	// 0 找到对应节点的父节点
	node, parentNode, nodeType := b.FindNode(i)
	// fmt.Println(node, parentNode, nodeType)
	if node == nil {
		return
	}
	// 1 如果删除的是叶子节点，把父节点对应的子节点删掉就可
	if node.left == nil && node.right == nil {
		switch nodeType {
		case "left":
			parentNode.left = nil
		case "right":
			parentNode.right = nil
		case "root":
			b.root = nil
		}
		return
	}
	// 2 如果要删除的节点不是叶子节点，且只有一颗子树
	var moveNode *Node
	if node.left == nil && node.right != nil {
		moveNode = node.right
	} else if node.right == nil && node.left != nil {
		moveNode = node.left
	}
	if moveNode != nil {
		switch nodeType {
		case "left":
			parentNode.left = moveNode
		case "right":
			parentNode.right = moveNode
		case "root":
			b.root = moveNode
		}
		return
	}
	// 3 如果要删除的节点不是叶子节点，而且有两颗子树
	if node.left != nil && node.right != nil {
		switch nodeType {
		// 如果删除的节点位于左子树
		case "left":
			// 3.1 找到要删除节点的右子树中最小（最左）的叶子
			leftestNode, leftestParentNode := b.LeftestNode(node.right)
			// 3.1.1 将右子树中最小的叶子删除
			if leftestNode == nil {
				node.data = leftestParentNode.data
				node.right = leftestParentNode.right
				return
			}
			leftestParentNode.left = nil
			// 3.1.2 将右子树中最小的叶子替换进去要删除的节点node
			node.data = leftestNode.data
		case "right":
			// 3.2 找到要删除节点的左子树中最大（最右）的叶子
			rightestNode, rightestParentNode := b.RightestNode(node.left)
			if rightestNode == nil {
				node.data = rightestParentNode.data
				node.left = rightestParentNode.left
				return
			}
			rightestParentNode.right = nil
			node.data = rightestNode.data
		case "root": // 根节点按照left或者right任一处理方式都可以
			rightestNode, rightestParentNode := b.RightestNode(node.left)
			rightestParentNode.right = nil
			b.root.data = rightestNode.data
		}
		return
	}
}

func (b *BST) LeftestNode(node *Node) (n *Node, pn *Node) {
	pn = node
	if pn.left == nil {
		return
	}
	for ; pn.left.left != nil; pn = pn.left {
	}
	n = pn.left
	return
}

func (b *BST) RightestNode(node *Node) (n *Node, pn *Node) {
	pn = node
	if pn.right == nil {
		return
	}
	for ; pn.right.right != nil; pn = pn.right {
	}
	n = pn.right
	return
}

func main() {
	arr := []int{10, 5, 1, 8, 7, 6, 9, 12, 11, 13}
	bst := BST{}
	for _, d := range arr {
		bst.InsertNode(NewNode(d))
	}
	bst.Walk()
	node, parentNode, nodeType := bst.FindNode(9)
	fmt.Println(node, parentNode, nodeType)
	// node, parentNode, nodeType = bst.FindNode2(9)
	// fmt.Println(node, parentNode, nodeType)
	bst.DelNode(10)
	bst.Walk()

}
