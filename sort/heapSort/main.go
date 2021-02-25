package main

import (
	"al/sort/common"
)

// 把数组调整为堆：按照大顶堆(升序)或者小顶堆（降序）进行排列，此时堆顶就是最大数
// 将堆顶元素与数组末尾元素交换，相当于选择排序，将最大数位置固定
// 继续处理数组[0:len-1],重复以上步骤

// 堆：一个完全二叉树
// 大顶堆：中间节点大于他的左右节点，左右节点大小没有要求
// 小顶堆：中间节点小于他的左右节点

type HeapSort struct {
}

func (h *HeapSort) adjust(arr []int, length int) {
	if length <= 1 {
		return
	}
	// 1 从末尾节点开始找到第一个非叶子节点，逐层调整成堆
	// fmt.Println("===========")
	var k int
	for i := length - 1; i-1 >= 0; i -= 2 {
		tmpHeapIdx := (i - 1) / 2
		k = 2*tmpHeapIdx + 1

		if 2*tmpHeapIdx+2 < length {
			if arr[2*tmpHeapIdx+1] < arr[2*tmpHeapIdx+2] {
				k++
			}
		}
		if arr[k] > arr[tmpHeapIdx] {
			arr[k], arr[tmpHeapIdx] = arr[tmpHeapIdx], arr[k]
		}

		// if 2*tmpHeapIdx+1 < length && arr[tmpHeapIdx] < arr[2*tmpHeapIdx+1] {
		// 	arr[tmpHeapIdx], arr[2*tmpHeapIdx+1] = arr[2*tmpHeapIdx+1], arr[tmpHeapIdx]
		// }
		// if 2*tmpHeapIdx+2 < length && arr[tmpHeapIdx] < arr[2*tmpHeapIdx+2] {
		// 	arr[tmpHeapIdx], arr[2*tmpHeapIdx+2] = arr[2*tmpHeapIdx+2], arr[tmpHeapIdx]
		// }
		// fmt.Println(arr)
	}
	// 2 交换堆的根节点和数组末尾元素，因为是顺序存储的树，所以就是交换第一个和第length-1个
	arr[0], arr[length-1] = arr[length-1], arr[0]
	// fmt.Println(arr)
	h.adjust(arr, length-1)
}

func main() {
	// arr := []int{4, 0, 6, 8, 5, 9, 10, 0, 90, 2}
	arr := common.GenRandInt(100000, 1000000, 1)
	hs := HeapSort{}
	hs.adjust(arr, len(arr))
	// fmt.Println("res:", arr)
}
