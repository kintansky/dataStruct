package main

import "fmt"

type ArrayStack struct {
	maxSize int
	data    []int
	top     int
}

func NewArrayStack(m int) *ArrayStack {
	return &ArrayStack{
		maxSize: m,
		data:    make([]int, m),
		top:     -1,
	}
}

func (a *ArrayStack) IsFull() (full bool) {
	if a.top == a.maxSize-1 {
		full = true
	}
	return
}

func (a *ArrayStack) IsEmpty() (empty bool) {
	if a.top == -1 {
		empty = true
	}
	return
}

func (a *ArrayStack) Push(n int) {
	if a.IsFull() {
		fmt.Println("Stack full")
		return
	}
	a.top++
	a.data[a.top] = n
}

func (a *ArrayStack) Pop() (v int) {
	if a.IsEmpty() {
		fmt.Println("Stack empty")
		return
	}
	v = a.data[a.top]
	a.top--
	return
}

func (a *ArrayStack) Show() {
	for tmp := a.top; tmp >= 0; tmp-- {
		fmt.Printf("<-[%d]", a.data[tmp])
	}
	fmt.Println()
}

func main() {
	s := NewArrayStack(10)
	for i := 0; i < 12; i++ {
		s.Push(i)
	}
	s.Show()
	s.Pop()
	s.Show()
}
