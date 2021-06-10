package main

import "fmt"

type CircleArrayQueue struct {
	maxSize int // 实际总可用是maxSize-1，因为预留了一个空位置给rear用来判断满
	front   int // 指向第一个元素
	rear    int // 指向最后一个元素的后一个index
	array   []int
}

func NewCircleArrayQueue(ms int) *CircleArrayQueue {
	return &CircleArrayQueue{
		maxSize: ms,
		front:   0,
		rear:    0,
		array:   make([]int, ms),
	}
}

func (c *CircleArrayQueue) IsFull() bool {
	// c.rear指向最后一个元素的后一个位置，而且这个位置是一个预留空位置，不能插入
	// 所以如果满，则意味着这个空位的下一位就是队头
	return (c.rear+1)%c.maxSize == c.front
}

func (c *CircleArrayQueue) IsEmpty() bool {
	return c.rear == c.front
}

func (c *CircleArrayQueue) Add(data int) {
	if c.IsFull() {
		fmt.Println("queue full")
		return
	}
	c.array[c.rear] = data
	c.rear = (c.rear + 1) % c.maxSize
}

func (c *CircleArrayQueue) Get() (data int) {
	if c.IsEmpty() {
		fmt.Println("queue empty")
		return
	}
	data = c.array[c.front]
	c.front = (c.front + 1) % c.maxSize
	return
}

func (c *CircleArrayQueue) showData(pos int) {
	if c.rear == pos {
		return
	}
	fmt.Printf("<-%d", c.array[pos])
	pos = (pos + 1) % c.maxSize
	c.showData(pos)
}

func (c *CircleArrayQueue) Show() {
	c.showData(c.front)
	fmt.Println()
}

func (c *CircleArrayQueue) Size() int {
	return (c.rear + c.maxSize - c.front) % c.maxSize
}

func main() {
	c := NewCircleArrayQueue(10)
	for i := 0; i < 11; i++ {
		c.Add(i)
	}
	c.Get()
	c.Get()
	c.Show()

	c.Add(100)
	c.Show()

	c.Add(101)
	c.Show()

	c.Get()
	c.Show()

	c.Add(102)
	c.Show()
}
