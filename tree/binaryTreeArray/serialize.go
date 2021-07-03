package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) prefixWalk() {
	fmt.Printf("%d ", t.Val)
	if t.Left != nil {
		t.Left.prefixWalk()
	}
	if t.Right != nil {
		t.Right.prefixWalk()
	}
}

func prefixWalkFromRoot(root *TreeNode) {
	root.prefixWalk()
	fmt.Println()
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res [][]string
	if root == nil {
		return "[]"
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		var levelData []string
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				levelData = append(levelData, "null")
				continue
			}
			levelData = append(levelData, fmt.Sprintf("%v", node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		queue = queue[size:]
		res = append(res, levelData)
	}
	var s = ""
	for i, levelData := range res {
		s += strings.Join(levelData, ",")
		if i != len(res)-1 {
			s += ","
		}
	}
	s = s[:len(s)-1]
	return fmt.Sprintf("[%s]", s)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var resArr []interface{}
	var dig string
	var digCnt int
	for _, d := range data {
		if d == '[' || d == ']' {
			continue
		}
		if d != ',' {
			dig += string(d)
			continue
		}
		if dig == "null" {
			resArr = append(resArr, nil)
		} else {
			num, _ := strconv.ParseInt(dig, 10, 32)
			node := TreeNode{Val: int(num)}
			resArr = append(resArr, &node)
			rootIdx := (digCnt - 1) / 2
			if rootIdx < 0 {
				continue
			}
			for resArr[rootIdx] == nil || (resArr[rootIdx].(*TreeNode).Left != nil && resArr[rootIdx].(*TreeNode).Right != nil) {
				rootIdx++
			}
			fmt.Println(num)
			fmt.Println("\trootIdx:", rootIdx)

			if (digCnt-1)%2 == 0 {
				resArr[rootIdx].(*TreeNode).Left = &node
			} else if (digCnt-1)%2 == 1 {
				resArr[rootIdx].(*TreeNode).Right = &node
			}
		}
		dig = ""
		digCnt++
	}
	if resArr != nil {
		return resArr[0].(*TreeNode)
	}
	return nil
}
