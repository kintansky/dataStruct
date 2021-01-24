package main

import "fmt"

// 8皇后问题，8*8棋盘，放置8个皇后，任意两个皇后不能在同一直线（行列）和斜线上

type Queue struct {
	size     int
	resArray []int // 结果数组，数组index对应第几个皇后（即第几行），值对应摆放的列位置
}

func NewQueue(s int) *Queue {
	return &Queue{
		size:     s,
		resArray: make([]int, s),
	}
}

func (q *Queue) Show() {
	fmt.Println(q.resArray)
}

// 对第queueNum个皇后进行conflict检测
func (q *Queue) DetectConflict(queueNum int) (conflict bool) {
	for i := 0; i < queueNum; i++ {
		// 同一列
		if q.resArray[i] == q.resArray[queueNum] {
			conflict = true
			break
		}
		// 同斜线，两个方向
		if (queueNum-i) == (q.resArray[queueNum]-q.resArray[i]) || -1*(queueNum-i) == (q.resArray[queueNum]-q.resArray[i]) {
			conflict = true
			break
		}
	}
	return
}

// Put 放置第queueNum个皇后
func (q *Queue) Put(queueNum int) {
	// 计算完所有皇后位置退出
	if queueNum >= q.size {
		q.Show()
		return
	}
	// 递归求解，递归过程其中一层q.DetectConflict为真，都会继续这一层的下一个循环
	for i := 0; i < q.size; i++ {
		q.resArray[queueNum] = i
		if q.DetectConflict(queueNum) {
			continue
		}
		q.Put(queueNum + 1)
	}
}

func main() {
	q := NewQueue(8)
	q.Put(0)
}
