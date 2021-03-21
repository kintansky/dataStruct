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

func main() {
	g := NewGraph(5)
	g.Insert("a")
	g.Insert("b")
	g.Insert("c")
	g.AddEdge(g.GetVertexIdx("a"), g.GetVertexIdx("c"), 1)
	g.AddEdge(g.GetVertexIdx("b"), g.GetVertexIdx("c"), 1)
	g.Show()
}
