package main

import (
	"al/sort/common"
	"fmt"
	"time"
)

type SelectSort struct {
	// 时间复杂度O(n^2)
	// 每一轮所有元素都与固定位置的元素进行比较，直到遍历完
	// 相当于每一轮都要遍历完剩下的数据，找出最大或最小的元素放在固定位置
}

func (s *SelectSort) Sort(arr []int) {
	for idx, _ := range arr {
		if idx == len(arr)-1 {
			break
		}
		for i := idx + 1; i < len(arr); i++ {
			if arr[idx] > arr[i] {
				arr[idx], arr[i] = arr[i], arr[idx]
			}
		}
	}
}

func main() {
	arr := common.GenRandInt(10, 100, time.Now().UnixNano())
	fmt.Println(arr)

	ss := SelectSort{}
	ss.Sort(arr)
	fmt.Println(arr)

}
