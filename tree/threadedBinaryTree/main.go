package main

import "fmt"

type Node struct {
	num       int
	left      *Node
	right     *Node
	leftType  int // 0代表指向左子树，1指向前驱节点
	rightType int // 0代表指向右子树，1指向后继节点
}

func NewNode(num int) *Node {
	return &Node{num: num}
}

func (n *Node) Show() {
	fmt.Printf("num:%d,", n.num)
}

// PrefixWalk 线索化后的前序遍历
func (n *Node) PrefixWalk() {
	n.Show()
	if !(n.left == nil || n.leftType == 1) {
		n.left.PrefixWalk()
	}
	if !(n.right == nil || n.rightType == 1) {
		n.right.PrefixWalk()
	}
}

// 线索化后的中序遍历
func (n *Node) MiddleWalk() {
	if !(n.left == nil || n.leftType == 1) { // 注意条件
		n.left.MiddleWalk()
	}
	n.Show()
	if !(n.right == nil || n.rightType == 1) {
		n.right.MiddleWalk()
	}
}

// 后序遍历
func (n *Node) SuffixWalk() {
	if !(n.left == nil || n.leftType == 1) {
		n.left.SuffixWalk()
	}
	if !(n.right == nil || n.rightType == 1) {
		n.right.SuffixWalk()
	}
	n.Show()
}

type Tree struct {
	root *Node
	pre  *Node // 辅助指针，总是指向当前节点的前驱节点
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Add(node *Node) {
	if t.root == nil {
		t.root = node
		return
	}
	tmp := t.root
	for tmp != nil {
		if node.num < tmp.num {
			if tmp.left == nil {
				tmp.left = node
				break
			}
			tmp = tmp.left
		} else if node.num > tmp.num {
			if tmp.right == nil {
				tmp.right = node
				break
			}
			tmp = tmp.right
		} else {
			fmt.Printf("node %v already exists\n", node)
			return
		}
	}
}

func (t *Tree) Search(num int) (node *Node) {
	if t.root == nil {
		return
	}
	tmp := t.root
	for tmp != nil {
		if num < tmp.num {
			if tmp.left == nil {
				break
			}
			tmp = tmp.left
		} else if num > tmp.num {
			if tmp.right == nil {
				break
			}
			tmp = tmp.right
		} else {
			node = tmp
			break
		}
	}
	return
}

func (t *Tree) PrefixWalk() {
	if t.root == nil {
		fmt.Println("empty")
		return
	}
	t.root.PrefixWalk()
	fmt.Println()
}

func (t *Tree) MiddleWalk() {
	if t.root == nil {
		fmt.Println("empty")
		return
	}
	t.root.MiddleWalk()
	fmt.Println()

}

func (t *Tree) SuffixWalk() {
	if t.root == nil {
		fmt.Println("empty")
		return
	}
	t.root.SuffixWalk()
	fmt.Println()

}

// 中序线索化二叉树
func (t *Tree) middleThreadedNodes(node *Node) {
	if node == nil {
		return
	}
	// 1 线索化左指数
	t.middleThreadedNodes(node.left)
	// 2 线索化当前节点的前驱指针和他的前驱点的后继指针
	// 2.1 前驱节点和类型：更新当前节点的左空指针指向前驱节点
	if node.left == nil {
		node.left = t.pre
		node.leftType = 1
	}
	// 2.2 后继节点和类型：更新当前节点的前驱节点的右空指针指向后继节点
	if t.pre != nil && t.pre.right == nil {
		t.pre.right = node
		t.pre.rightType = 1
	}
	t.pre = node
	// 3 线索化右指数
	t.middleThreadedNodes(node.right)

}

func (t *Tree) MiddleThreadedNodes() {
	t.middleThreadedNodes(t.root)
}

func main() {
	t := NewTree()
	root := NewNode(1)
	node2 := NewNode(3)
	node3 := NewNode(6)
	node4 := NewNode(8)
	node5 := NewNode(10)
	node6 := NewNode(14)
	node7 := NewNode(15)
	t.root = root
	root.left = node2
	root.right = node3
	node2.left = node4
	node2.right = node5
	node3.left = node6
	node4.right = node7

	fmt.Println(node5.left, node5.right)
	t.MiddleThreadedNodes()
	fmt.Println(node5.left, node5.right)

	t.MiddleWalk()
}
