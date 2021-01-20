package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	num       int64
	oper      string
	numOrOper bool // 数字：true，操作符：false
	pre       *Node
}

func NewNode(n int64, o string, numOrOper bool) *Node {
	return &Node{
		num:       n,
		oper:      o,
		numOrOper: numOrOper,
	}
}

type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return &Stack{
		top: nil,
	}
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

func (s *Stack) Show() {
	for tmp := s.top; tmp != nil; tmp = tmp.pre {
		if tmp.numOrOper {
			fmt.Printf("<-[%d]", tmp.num)
			continue
		}
		fmt.Printf("<-[%s]", tmp.oper)
	}
	fmt.Println()
}

func getNodeFromStr(nodeChan chan<- *Node, s string) {
	var tmpNumStr string
	for _, c := range s {
		// 数字先缓存
		if c >= 48 && c <= 57 {
			tmpNumStr += string(c)
			continue
		}
		// 遇到操作符，先把缓存下来的数字解析成数字，推入chan
		n, _ := strconv.ParseInt(tmpNumStr, 10, 64)
		nodeChan <- NewNode(n, "", true)
		tmpNumStr = ""
		// 然后将操作符推出chan
		nodeChan <- NewNode(0, string(c), false)
	}
	n, _ := strconv.ParseInt(tmpNumStr, 10, 64)
	nodeChan <- NewNode(n, "", true)
	close(nodeChan)
}

// 1 如果是符号栈空，数字入数字栈，符号栈入符号栈
// 2 符号栈不空，
// 2.1 数字可直接入数字栈；
// 2.2 符号栈比较准备入栈的符号与栈顶元素优先级：
// 2.2.1 栈顶优先级低的，符号直接入栈
// 2.2.2 栈顶优先级高的，数字栈弹出两个数字，符号栈弹出栈顶符号，进行计算，结果压入数字栈，然后再入栈准备入栈的符号

var operPriorityMap map[string]int = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
}

func cal(num1, num2 int64, oper string) (res int64) {
	switch oper {
	case "+":
		res = num2 + num1
	case "-":
		res = num2 - num1
	case "*":
		res = num2 * num1
	case "/":
		res = num2 / num1
	}
	return
}

func run(nodeChan <-chan *Node) (ret int64) {
	numStack := NewStack()
	operStack := NewStack()
	var (
		num1   int64
		num2   int64
		tmpRes int64
	)
	for node := range nodeChan {
		if node.numOrOper {
			numStack.Push(node)
			continue
		}
		if operStack.IsEmpty() {
			operStack.Push(node)
			continue
		}
		if operPriorityMap[node.oper] <= operPriorityMap[operStack.top.oper] {
			num1 = numStack.Pop().num
			num2 = numStack.Pop().num
			tmpRes = cal(num1, num2, operStack.Pop().oper)
			numStack.Push(NewNode(tmpRes, "", true))
		}
		operStack.Push(node)
	}
	// 字符串遍历后两个栈的结果
	numStack.Show()
	operStack.Show()
	// 最后出栈所有符号，计算结果
	for operStack.top != nil {
		num1 = numStack.Pop().num
		num2 = numStack.Pop().num
		tmpRes = cal(num1, num2, operStack.Pop().oper)
		numStack.Push(NewNode(tmpRes, "", true))
	}
	ret = numStack.Pop().num
	return
}

func main() {
	str := "30+2*3-40"
	nodeChan := make(chan *Node, 0)
	go getNodeFromStr(nodeChan, str)
	ret := run(nodeChan)
	fmt.Printf("%s=%d\n", str, ret)
}
