package main

import "fmt"

// FIFO
type ArrayQueue struct {
	maxSize int
	front   int // 指向第一个元素的前一个index
	rear    int // 指向最后一个元素
	array   []int
}

func NewArrayQueue(ms int) *ArrayQueue {
	return &ArrayQueue{
		maxSize: ms,
		front:   -1,
		rear:    -1,
		array:   make([]int, ms),
	}
}

func (a *ArrayQueue) Add(data int) {
	if a.IsFull() {
		fmt.Println("queue full")
		return
	}
	a.rear++
	a.array[a.rear] = data
}

func (a *ArrayQueue) Get() (data int) {
	if a.IsEmpty() {
		fmt.Println("queue empty")
		return
	}
	a.front++
	data = a.array[a.front]
	return
}

func (a *ArrayQueue) showData(pos int) {
	if pos == a.rear {
		return
	}
	pos++
	fmt.Printf("<-%d", a.array[pos])
	a.showData(pos)
}

func (a *ArrayQueue) Show() {
	a.showData(a.front)
	fmt.Println()
}

func (a *ArrayQueue) IsFull() bool {
	return a.rear == a.maxSize-1
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.front == a.rear
}

func (a *ArrayQueue) GetFront() int {
	return a.front
}

func (a *ArrayQueue) GetRear() int {
	return a.rear
}

func main() {
	q := NewArrayQueue(10)
	for i := 0; i < 11; i++ {
		q.Add(i)
	}
	q.Show()
	for i := 0; i < 11; i++ {
		q.Get()
		q.Show()
	}
}
