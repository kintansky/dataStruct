package main

import (
	"al/sort/common"
	"fmt"
)

// ShellSort 希尔排序
type ShellSort struct {
}

// 0 1 2 3 4 5 6 7 8 9 10 11 12 共13个
// step = 6
// 0,6 1,7 2,8 3,9 4,10 5,11 12
// step = 6/2 = 3
// 0,3,6,9,12 1,4,7,10 2,5,8,11
// step = 3/2 = 1

// Sort 排序
func (s *ShellSort) Sort(arr []int, step int) {
	if step == 0 {
		return
	}
	for i := 0; i < step; i++ {
		// 每隔step取出一个值，进行插入排序
		for j := 1; i+step*j < len(arr); j++ {
			// 取出的值与前面step间隔的值，进行比较进行插入排序，注意运行条件curIdx > i
			for curIdx := i + step*j; curIdx > i; curIdx = curIdx - step {
				if arr[curIdx] >= arr[curIdx-step] {
					break
				}
				fmt.Printf("i=%d, step=%d, swap {idx:%d, data:%d}<=>{idx:%d, data:%d}\n", i, step, curIdx, arr[curIdx], curIdx-step, arr[curIdx-step])
				arr[curIdx], arr[curIdx-step] = arr[curIdx-step], arr[curIdx] // 交换的位置是往前step位置的值
				fmt.Println(arr)
			}
		}
	}
	s.Sort(arr, step/2)
}

func main() {
	arr := common.GenRandInt(10, 100)
	fmt.Println(arr)
	ss := ShellSort{}
	ss.Sort(arr, len(arr)/2)
	fmt.Println(arr)
}
