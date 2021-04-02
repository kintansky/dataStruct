package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	arr       []string
	vertexNum int
	edges     [][]int // 顶点的邻接矩阵
	edgesNum  int
}

func NewGraph(n int) *Graph {
	g := &Graph{
		arr:   make([]string, n),
		edges: make([][]int, n),
	}
	for i := 0; i < n; i++ {
		g.edges[i] = make([]int, n)
	}
	return g
}

func (g *Graph) Insert(v string) {
	if g.vertexNum >= len(g.arr) {
		fmt.Println("full")
		return
	}
	g.arr[g.vertexNum] = v
	g.vertexNum++
}

// AddEdge 连接两个顶点，v1Idx,V2Idx是对应两个顶点的INDEX
func (g *Graph) AddEdge(v1Idx, v2Idx int, weight int) {
	g.edges[v1Idx][v2Idx] = weight
	g.edges[v2Idx][v1Idx] = weight
	g.edgesNum++
}

func (g *Graph) GetVertexIdx(v string) (idx int) {
	for i := 0; i < g.vertexNum; i++ {
		if v == g.arr[i] {
			return i
		}
	}
	return
}

func (g *Graph) GetVertexNum() int {
	return g.vertexNum
}

func (g *Graph) GetEdgesNum() int {
	return g.edgesNum
}

func (g *Graph) GetVertex(i int) string {
	return g.arr[i]
}

func (g *Graph) GetWeight(v1, v2 int) int {
	return g.edges[v1][v2]
}

func (g *Graph) Show() {
	fmt.Println(" ", strings.Join(g.arr, " "))
	for i, line := range g.edges {
		fmt.Printf(g.arr[i] + " ")
		for _, d := range line {
			fmt.Printf("%d ", d)
		}
		fmt.Println()
	}
}

type Queue struct {
	maxSize int
	front   int
	rear    int
	arr     []int
}

func NewQueue(s int) *Queue {
	return &Queue{
		maxSize: s,
		front:   -1, // 队头
		rear:    -1, // 队尾
		arr:     make([]int, s),
	}
}

func (q *Queue) Push(i int) {
	q.rear++
	q.arr[q.rear] = i
}

func (q *Queue) Pop() (i int) {
	q.front++
	i = q.arr[q.front]
	return
}

func (q *Queue) IsEmpty() bool {
	return q.rear == q.front
}

func (q *Queue) IsFull() bool {
	return q.rear == q.maxSize-1
}

func (g *Graph) BFS(startVertexIdx int) {
	visited := make([]bool, g.vertexNum)
	// 非连通图情况下的遍历
	for i := startVertexIdx; i < startVertexIdx+g.vertexNum; i++ {
		idx := i % g.vertexNum
		if visited[idx] {
			continue
		}
		g.bfs(idx, visited)
	}
}

func (g *Graph) bfs(startVertexIdx int, visited []bool) {
	queue := NewQueue(g.vertexNum)

	fmt.Println(g.GetVertex(startVertexIdx))
	visited[startVertexIdx] = true
	queue.Push(startVertexIdx) // 访问完后本顶点入队，出队的时候用于访问他的所有邻接点
	for !queue.IsEmpty() {
		idx := queue.Pop()
		// BFS优先遍历完当前顶点的所有邻接点
		for i := 0; i < g.vertexNum; i++ {
			// 如果顶点有未访问的邻接点，进行访问并入队
			if g.edges[idx][i] != 0 && !visited[i] {
				fmt.Println(g.GetVertex(i))
				visited[i] = true
				queue.Push(i)
			}
		}
	}
}

func main() {
	g := NewGraph(5)
	g.Insert("a")
	g.Insert("b")
	g.Insert("c")
	g.Insert("d")
	g.Insert("e")
	g.AddEdge(g.GetVertexIdx("a"), g.GetVertexIdx("c"), 1)
	// g.AddEdge(g.GetVertexIdx("b"), g.GetVertexIdx("c"), 1)
	// g.AddEdge(g.GetVertexIdx("a"), g.GetVertexIdx("e"), 1)
	// g.AddEdge(g.GetVertexIdx("b"), g.GetVertexIdx("d"), 1)
	// g.AddEdge(g.GetVertexIdx("a"), g.GetVertexIdx("d"), 1)
	g.AddEdge(g.GetVertexIdx("d"), g.GetVertexIdx("e"), 1)
	g.Show()
	g.BFS(4)
}
