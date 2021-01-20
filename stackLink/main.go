package main

import "fmt"

type Node struct {
	num int
	pre *Node
}

func NewNode(n int) *Node {
	return &Node{num: n, pre: nil}
}

type LinkStack struct {
	top *Node
}

func NewLinkStack() *LinkStack {
	return &LinkStack{top: nil}
}

func (l *LinkStack) IsEmpty() (empty bool) {
	if l.top == nil {
		empty = true
	}
	return
}

func (l *LinkStack) Push(node *Node) {
	if l.IsEmpty() {
		l.top = node
		return
	}
	node.pre = l.top
	l.top = node
}

func (l *LinkStack) Pop() (node *Node) {
	if l.IsEmpty() {
		return
	}
	node = l.top
	l.top = l.top.pre
	return
}

func (l *LinkStack) Show() {
	for tmp := l.top; tmp != nil; tmp = tmp.pre {
		fmt.Printf("<-[%d]", tmp.num)
	}
	fmt.Println()
}

func main() {
	ls := NewLinkStack()
	for i := 0; i < 10; i++ {
		ls.Push(NewNode(i))
	}
	ls.Show()
	for i := 0; i < 12; i++ {
		ls.Pop()
		ls.Show()
	}
	ls.Show()

}
