package main

import (
	"al/sort/common"
	"fmt"
	"time"
)

// 把数组调整为堆：按照大顶堆(升序)或者小顶堆（降序）进行排列，此时堆顶就是最大数
// 将堆顶元素与数组末尾元素交换，相当于选择排序，将最大数位置固定
// 继续处理数组[0:len-1],重复以上步骤

// 堆：一个完全二叉树
// 大顶堆：中间节点大于他的左右节点，左右节点大小没有要求
// 小顶堆：中间节点小于他的左右节点

type HeapSort struct {
}

func (h *HeapSort) adjustSubTree(arr []int, subTreeRootIdx int, length int) {
	if 2*subTreeRootIdx+1 >= length { // 说明这个节点已经到了叶子节点，不用再继续调整，因为堆是完全二叉树，叶子左连续，所以判断左侧即可
		return
	}
	var adjustedNodeIdx int = subTreeRootIdx // 记录每次上层节点调整的是那个节点，用于往下重新调整子树为堆
	// 左节点是2*i+1，右节点是2*i+2
	if 2*subTreeRootIdx+2 < length {
		if arr[2*subTreeRootIdx+1] > arr[2*subTreeRootIdx+2] && arr[2*subTreeRootIdx+1] > arr[subTreeRootIdx] {
			arr[2*subTreeRootIdx+1], arr[subTreeRootIdx] = arr[subTreeRootIdx], arr[2*subTreeRootIdx+1]
			adjustedNodeIdx = 2*subTreeRootIdx + 1
		} else if arr[2*subTreeRootIdx+2] > arr[2*subTreeRootIdx+1] && arr[2*subTreeRootIdx+2] > arr[subTreeRootIdx] {
			arr[2*subTreeRootIdx+2], arr[subTreeRootIdx] = arr[subTreeRootIdx], arr[2*subTreeRootIdx+2]
			adjustedNodeIdx = 2*subTreeRootIdx + 2
		}
	} else if 2*subTreeRootIdx+1 < length {
		if arr[2*subTreeRootIdx+1] > arr[subTreeRootIdx] {
			arr[2*subTreeRootIdx+1], arr[subTreeRootIdx] = arr[subTreeRootIdx], arr[2*subTreeRootIdx+1]
			adjustedNodeIdx = 2*subTreeRootIdx + 1
		}
	}
	if adjustedNodeIdx == subTreeRootIdx { // 如果调整的节点没变，说明不需要调整，退出即可
		return
	}
	// 因为adjustedNodeIdx子树的根变化，所以递归调整adjustedNodeIdx为根的子树为堆
	h.adjustSubTree(arr, adjustedNodeIdx, length)
}

// 将无序序列转换成大顶堆，
func (h *HeapSort) Sort(arr []int) {

	for length := len(arr); length > 1; {
		// 1 从右至左，从下往上，非叶子节点依次为len/2-1递减，逐个调整成堆
		for i := length/2 - 1; i >= 0; i-- {
			// 2 调整以当前节点为根的子树，成为大顶堆或小顶堆（调整有可能破坏两个左右节点原本已经调整成堆的子树！！！）
			// h.adjustSubTree(arr, i, length)
			h.adjustSubTree2(arr, i, length)
			// fmt.Println("i:", i, "arr:", arr)
		}
		// 3 交换首尾，每次把当轮最大数放置末尾
		arr[0], arr[length-1] = arr[length-1], arr[0]
		// fmt.Println("swap:", arr)
		// 4 因为每次完成后，最大数已经算出，所以调整下一次操作的数组边界
		length--
	}

}

func (h *HeapSort) adjustSubTree2(arr []int, i int, length int) {
	tmp := arr[i]
	for k := i*2 + 1; k < length; {
		// 1 取左右节点中较大的那个
		if k+1 < length && arr[k] < arr[k+1] {
			// 因为完全二叉树叶子左连续，所以一个非叶子节点至少会有左节点k，所以判断越界，判断右节点即可k+1
			k++
		}
		// 2 判断是否需要交换当前子树的根和叶子成为堆
		if arr[k] > arr[i] {
			// fmt.Printf("arr[i]=%d, arr[k]=%d\n", arr[i], arr[k])
			// 2.1 更新子树的根为子树的最大值
			// arr[i], arr[k] = arr[k], arr[i]
			arr[i] = arr[k]
			// fmt.Printf("arr[i]=%d, arr[k]=%d\n", arr[i], arr[k])

			// 2.2 子树更新后，由于其中一个叶子也被调整了，所以以这个叶子为根子树的需要进一步调整回来成为堆
			// 所以下一个被调整的子树的根就是k
			i = k
		} else {
			// 3 如果当前子树不需要调整，那他的叶子其实也没有动，所以可以退出循环
			break
		}
		// 4 继续下一个子树
		k = 2*i + 1 // 下一个子树的左节点
	}
	arr[i] = tmp
}

func (h *HeapSort) Sort2(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.adjustSubTree2(arr, i, len(arr))
	}
	for j := len(arr) - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		h.adjustSubTree2(arr, 0, j)
	}
}

func main() {
	// arr := []int{4, 6, 8, 5, 9}
	arr := common.GenRandInt(100, 1000000, 1)
	hs := HeapSort{}
	t1 := time.Now()
	hs.Sort2(arr)
	fmt.Println(time.Now().Sub(t1))
	fmt.Println("res:", arr)
	// fmt.Println(arr)
}
