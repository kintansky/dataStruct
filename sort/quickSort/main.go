package main

import (
	"al/sort/common"
	"fmt"
	"time"
)

type QuickSort struct {
	// quickSort实际上是一个冒泡排序的优化
}

func (q *QuickSort) Sort(arr []int, left, right int) {
	if left >= right { // 退出条件
		return
	}
	mid := (left + right) / 2
	leftIdx := mid
	rightIdx := mid
	var (
		leftFound  bool
		rightFound bool
	)
	for leftIdx >= left || rightIdx <= right {
		// 1 找到左侧比arr[mid]大的数
		leftFound = false
		for ; leftIdx >= left; leftIdx-- {
			if arr[leftIdx] > arr[mid] {
				leftFound = true
				break
			}
		}
		rightFound = false
		// 2 找到右侧比arr[mid]小的数
		for ; rightIdx <= right; rightIdx++ {
			if arr[rightIdx] < arr[mid] {
				rightFound = true
				break
			}
		}
		// fmt.Println(leftFound, leftIdx, rightFound, rightIdx)
		// 3 数据调整，进行冒泡
		if leftFound && rightFound { // 如果左右都存在冲突的数，交换二者
			arr[leftIdx], arr[rightIdx] = arr[rightIdx], arr[leftIdx]
		} else if leftFound { // 如果只有左侧有冲突数据，则进行冒泡至mid位置
			for swapIdx := leftIdx; swapIdx < mid; swapIdx++ {
				arr[swapIdx], arr[swapIdx+1] = arr[swapIdx+1], arr[swapIdx]
			}
			mid--
		} else if rightFound { // 如果只有右侧有冲突数据，则进行冒泡至mid位置
			for swapIdx := rightIdx; swapIdx > mid; swapIdx-- {
				arr[swapIdx], arr[swapIdx-1] = arr[swapIdx-1], arr[swapIdx]
			}
			mid++
		} else { // 两边都没有冲突，说明已经整理好，可以直接break
			break
		}
	}
	// 递归进行左右侧排序
	q.Sort(arr, left, mid-1)
	q.Sort(arr, mid+1, right)
}

func main() {

	arr := common.GenRandInt(100000, 1000000, time.Now().UnixNano())
	// fmt.Println(arr)
	t1 := time.Now()
	qs := QuickSort{}
	qs.Sort(arr, 0, len(arr)-1)
	fmt.Println(time.Now().Sub(t1))
	// fmt.Println(arr)

}
