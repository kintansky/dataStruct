package main

// 层序遍历
type queue struct {
	arr  []interface{}
	head int
	tail int
}

func NewQueue() *queue {
	return &queue{
		arr:  make([]interface{}, 0),
		head: -1,
		tail: -1,
	}
}

func (q *queue) add(i interface{}) {
	q.tail++
	q.arr = append(q.arr, i)
}

func (q *queue) pop() (i interface{}) {
	if q.isEmpty() {
		return
	}
	q.head++
	i = q.arr[q.head]
	return
}

func (q *queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *queue) size() int {
	return q.tail - q.head
}

func levelOrder(root *Node) [][]int { // BFS思想
	if root == nil {
		return nil
	}
	q := NewQueue()
	q.add(root) // 根节点加入队列
	var res [][]int
	for !q.isEmpty() { // 如果当前层次还有数据的，继续出队
		var levelData []int
		size := q.size() // 因为在每层节点出队的同时，其子节点也同时入队，所以只能通过当前层次的个数，判断是否已经遍历完一层的数据
		for i := 0; i < size; i++ {
			node := q.pop().(*Node)
			levelData = append(levelData, node.num)
			// 当前节点的左右节点就是下一层的数据，当前节点出队后加入队列
			if node.left != nil {
				q.add(node.left)
			}
			if node.right != nil {
				q.add(node.right)
			}
		}
		// 一层数据遍历完后加入结果集
		res = append(res, levelData)
	}
	return res
}
