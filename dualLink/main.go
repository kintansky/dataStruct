package main

import "fmt"

type Node struct {
	num  int
	name string
	pre  *Node
	next *Node
}

func NewNode(n int, name string) *Node {
	return &Node{
		num:  n,
		name: name,
	}
}

type DualLink struct {
	head *Node
}

func NewDualLink() *DualLink {
	return &DualLink{
		head: NewNode(-1, ""),
	}
}

func (d *DualLink) AddByOrder(node *Node) {
	if d.IsEmpty() {
		d.head.next = node
		node.pre = d.head
		return
	}
	tmp := d.head.next
	for ; tmp != nil; tmp = tmp.next {
		if tmp.num == node.num {
			fmt.Println("node already exists")
			return
		}
		if tmp.num > node.num {
			break
		}
	}
	tmp.pre.next = node
	node.pre = tmp.pre
	node.next = tmp
	tmp.pre = node
}

func (d *DualLink) IsEmpty() (empty bool) {
	if d.head.next == nil {
		empty = true
	}
	return
}

func (d *DualLink) Update(node *Node) {
	tmp := d.head.next
	for ; tmp != nil; tmp = tmp.next {
		if tmp.num == node.num {
			node.next = tmp.next
			node.pre = tmp.pre
			tmp.pre.next = node
			tmp.next.pre = node
			return
		}
	}
	fmt.Println("node not exists")
}

func (d *DualLink) Del(n int) {
	tmp := d.head.next
	for ; tmp != nil; tmp = tmp.next {
		if tmp.num == n {
			tmp.pre.next = tmp.next
			tmp.next.pre = tmp.pre
			return
		}
	}
	fmt.Println("node not exists")
}

func (d *DualLink) Show() {
	tmp := d.head.next
	for ; tmp != nil; tmp = tmp.next {
		fmt.Printf("[%d:%s]<->", tmp.num, tmp.name)
		// fmt.Printf("[%d:%s(pre:%p,next:%p)]<->", tmp.num, tmp.name, tmp.pre, tmp.next)
	}
	fmt.Println()
}

func (d *DualLink) reverse(node *Node) {
	if node == nil {
		return
	}
	d.reverse(node.next)
	tmp := node.pre
	node.pre = node.next
	node.next = tmp
	if node.pre == nil { // 如果node是最后一个节点，反转后要修正成head的next为这个node
		node.pre = d.head
		d.head.next = node
	}
	if node.next == d.head { // 如果node是第一个节点，反转后修正node.next=nil，否则陷入head的死循环
		node.next = nil
	}
}

func (d *DualLink) ReverseDualLink() {
	d.reverse(d.head.next)
}

func main() {
	dl := NewDualLink()
	for i := 0; i < 3; i++ {
		dl.AddByOrder(NewNode(10-i, fmt.Sprintf("node_%d", 10-i)))
	}
	dl.Show()
	dl.Update(NewNode(5, "test_5"))
	dl.Show()
	dl.Del(3)
	dl.Show()
	dl.ReverseDualLink()
	dl.Show()
}
