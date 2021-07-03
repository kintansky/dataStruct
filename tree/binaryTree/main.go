package main

import "fmt"

type Node struct {
	num   int
	left  *Node
	right *Node
}

func NewNode(num int) *Node {
	return &Node{num: num}
}

func (n *Node) Show() {
	fmt.Printf("num:%d,", n.num)
}

// PrefixWalk 前序遍历
func (n *Node) PrefixWalk() {
	n.Show()
	if n.left != nil {
		n.left.PrefixWalk()
	}
	if n.right != nil {
		n.right.PrefixWalk()
	}
}

// 中序遍历
func (n *Node) MiddleWalk() {
	if n.left != nil {
		n.left.MiddleWalk()
	}
	n.Show()
	if n.right != nil {
		n.right.MiddleWalk()
	}
}

// 后序遍历
func (n *Node) SuffixWalk() {
	if n.left != nil {
		n.left.SuffixWalk()
	}
	if n.right != nil {
		n.right.SuffixWalk()
	}
	n.Show()
}

type Tree struct {
	root *Node
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

func main() {
	// t := NewTree()
	// for i := 5; i < 8; i++ {
	// 	t.Add(NewNode(i))
	// }
	// for i := 1; i < 5; i++ {
	// 	t.Add(NewNode(i))
	// }
	// // node1 := NewNode(1)
	// // node2 := NewNode(2)
	// // node3 := NewNode(3)
	// // node4 := NewNode(4)
	// // node1.left = node2
	// // node1.right = node3
	// // node3.right = node4
	// // t.root = node1
	// t.PrefixWalk()
	// t.MiddleWalk()
	// t.SuffixWalk()

	// fmt.Println(t.Search(7))
	// fmt.Println(t.Search(2))
	// fmt.Println(t.Search(5))
	// fmt.Println(t.Search(20))
	t := NewTree()
	t.root = NewNode(3)
	t.root.left = NewNode(9)
	t.root.right = NewNode(9)
	// r := NewNode(20)
	// r.left = NewNode(15)
	// r.right = NewNode(7)
	// t.root.right = r
	// res := levelOrder(t.root)
	// fmt.Println(res)

	// maxDeep := maxDepth(t.root)
	// fmt.Println("maxDeep:", maxDeep)
	// maxDeep = maxDepth2(t.root)
	// fmt.Println("maxDeep:", maxDeep)
	fmt.Println(isSymmetric(t.root))
}
