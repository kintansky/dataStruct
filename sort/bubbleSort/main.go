package main

import (
	"al/sort/common"
	"fmt"
	"time"
)

type BubbleSort struct {
	// 两两比较，符合条件的往上升
	// 时间复杂度O(n^2)
}

func (b *BubbleSort) Sort(arr []int) {
	var cnt int // 统计一轮中的交换次数，如果交换次数=0，说明本轮未排的数据实际上已经是排好序的数据，后面无需继续冒泡
	// 每轮比较结果是靠近顶部的数据逐渐排好序
	for i := 0; i < len(arr)-1; i++ {
		cnt = 0
		// 每轮都需要比较len(arr-i-1)次，因为顶层已经排好
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] <= arr[j+1] {
				continue
			}
			arr[j], arr[j+1] = arr[j+1], arr[j]
			cnt++
		}
		// 交换次数=0，说明本轮未排的数据实际上已经是排好序的数据，后面无需继续冒泡
		if cnt == 0 {
			fmt.Println("already satisfied:", arr)
			break
		}
		// fmt.Println(arr)

	}
	return
}

func main() {
	arr := common.GenRandInt(100, 1000, time.Now().UnixNano())
	fmt.Println("org:", arr)
	b := BubbleSort{}
	b.Sort(arr)
	fmt.Println(arr)
}
