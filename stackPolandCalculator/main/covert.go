package main

import "strconv"

// 中缀表达式转后缀表达式
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

// 1. 使用一个缓存栈+一个队列
// 2. 数字：直接进入输出队列
// 3. 操作符
// 	3.1 缓存栈为空，操作符直接入缓存栈
// 	3.2 缓存栈非空
// 		3.2.1 当前操作符优先级<栈顶操作符：缓存栈栈顶出栈进入输出队列，然后继续循环3.2.1，直到不成立
// 		3.2.2 当前操作符优先级>=栈顶操作符：当前操作符直接入缓存栈
// 	3.3 ()改变运算顺序的情况
// 		3.3.1 ( 直接入缓存栈，其他按照2、3.1 和3.2继续入栈，直到碰到 )
// 		3.3.2 ) 不入栈，并将缓存栈的符号依次弹出添加到输出队列，直到碰到第一个配对的 (，并出栈丢弃 (
// 4. 遍历完，缓存栈依次弹出到输出队列

var operPriorityMap map[string]int = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
}

func convertToSuffix(infixArray []string) (suffixArray []*Node) {
	bufferStack := NewStack()
	for _, str := range infixArray {
		if isNum(str) {
			num, _ := strconv.ParseInt(str, 10, 64)
			suffixArray = append(suffixArray, NewNode(num, "", true))
			continue
		}
		// 操作符
		if str == "(" {
			bufferStack.Push(NewNode(0, str, false))
			continue
		}
		if str == ")" {
			for bufferStack.top.oper != "(" {
				suffixArray = append(suffixArray, bufferStack.Pop())
			}
			bufferStack.Pop()
			continue
		}
		if bufferStack.IsEmpty() {
			bufferStack.Push(NewNode(0, str, false))
			continue
		}
		for !bufferStack.IsEmpty() && operPriorityMap[str] < operPriorityMap[bufferStack.top.oper] {
			suffixArray = append(suffixArray, bufferStack.Pop())
		}
		bufferStack.Push(NewNode(0, str, false))
	}
	for !bufferStack.IsEmpty() {
		suffixArray = append(suffixArray, bufferStack.Pop())
	}
	return
}
