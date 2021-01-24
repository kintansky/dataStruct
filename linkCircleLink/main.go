package main

import "fmt"

type Node struct {
	num  int
	next *Node
}

func NewNode(n int) *Node {
	return &Node{num: n}
}

type CircleLink struct {
	first *Node
}

func NewCircleLink() *CircleLink {
	return &CircleLink{
		first: nil,
	}
}

func (c *CircleLink) Add(node *Node) {
	if c.IsEmpty() {
		c.first = node
		node.next = node
		return
	}
	cur := c.first
	for cur.next != c.first {
		cur = cur.next
	}
	node.next = c.first
	cur.next = node
}

func (c *CircleLink) IsEmpty() (empty bool) {
	if c.first == nil {
		empty = true
	}
	return
}

func (c *CircleLink) Show() {
	if c.IsEmpty() {
		return
	}
	cur := c.first
	for {
		fmt.Printf("[%d]->", cur.num)
		if cur.next == c.first {
			break
		}
		cur = cur.next
	}
	fmt.Println()
}

func (c *CircleLink) GetLast() (node *Node, length int) {
	if c.IsEmpty() {
		return
	}
	node = c.first
	for node.next != c.first {
		node = node.next
		length++
	}
	return
}

func Joseph(cl *CircleLink, beginNum int, cnt int) {
	// 单向循环链表只能通过前一个节点，删除下一个节点，所以辅助指针先指向最后一个节点
	// 通过一个length值，可以减少循环的次数
	preNode, length := cl.GetLast()
	// 定位起始位置
	for i := 0; i < (beginNum-1)%length; i++ {
		preNode = preNode.next
	}
	// 开始出队
	for preNode.next != preNode {
		// 起始位开始，第cnt个出队，然后下一个又继续开始
		for i := 0; i < (cnt-1)%length; i++ {
			preNode = preNode.next
		}
		fmt.Printf("pop [%d]\n", preNode.next.num)
		preNode.next = preNode.next.next
		length--
	}
	fmt.Printf("pop [%d]\n", preNode.num)
}

func main() {
	cl := NewCircleLink()
	for i := 0; i < 10; i++ {
		cl.Add(NewNode(i))
	}
	cl.Show()
	Joseph(cl, 3, 200)
}
