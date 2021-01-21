package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	num int64
	pre *Node
}

func NewNode(n int64) *Node {
	return &Node{num: n}
}

type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return &Stack{top: nil}
}

func (s *Stack) IsEmpty() (empty bool) {
	if s.top == nil {
		empty = true
	}
	return
}

func (s *Stack) Push(node *Node) {
	if s.IsEmpty() {
		s.top = node
		return
	}
	node.pre = s.top
	s.top = node
}

func (s *Stack) Pop() (node *Node) {
	if s.IsEmpty() {
		return
	}
	node = s.top
	s.top = s.top.pre
	return
}

func isNum(s string) (y bool) {
	re := regexp.MustCompile(`\d+`)
	y = re.MatchString(s)
	return
}

func cal(num1, num2 int64, oper string) (ret int64) {
	switch oper {
	case "+":
		ret = num2 + num1
	case "-":
		ret = num2 - num1
	case "*":
		ret = num2 * num1
	case "/":
		ret = num2 / num1
	}
	return
}

func main() {
	//           -
	//			/ \
	//         *   6
	//        / \
	//       +   5
	//      / \
	//     3   4
	//    /     \
	//   (       )
	// 中缀表达式(3+4)*5-6的后缀表达式（逆波兰表达式）: ( 3 ) 4 + 5 * 6 -
	suffixExp := "3 4 + 5 * 6 -"
	arrayList := strings.Split(suffixExp, " ")
	s := NewStack()
	var (
		num1 int64
		num2 int64
	)
	// 与中缀相比，只需要一个栈就可以完成计算，不需要判断符号优先级
	// 1 遇到数字就入栈
	// 2 遇到操作符就弹出栈顶两个数字就行计算，结果重新入栈
	// 3 最后重zhan中取出结果
	for _, d := range arrayList {
		if isNum(d) {
			num, _ := strconv.ParseInt(d, 10, 64)
			s.Push(NewNode(num))
			continue
		}
		num1 = s.Pop().num
		num2 = s.Pop().num
		s.Push(NewNode(cal(num1, num2, d)))
	}
	res := s.Pop()
	fmt.Println(res)
}
