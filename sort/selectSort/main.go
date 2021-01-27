package main

import (
	"al/sort/common"
	"fmt"
)

type SelectSort struct {
	// 时间复杂度O(n^2)
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
	arr := common.GenRandInt(10, 100)
	fmt.Println(arr)

	ss := SelectSort{}
	ss.Sort(arr)
	fmt.Println(arr)

}
