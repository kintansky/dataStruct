package main

import (
	"al/sort/common"
	"fmt"
)

type MergeSort struct {
}

func (m *MergeSort) Split(arr []int, headIdx, tailIdx int) (newArr []int) {
	// 不断划分直至只剩一个元素开始return，返回时还是以数组形式返回[a,]
	// 然后进行有序合并
	if headIdx+1 == tailIdx {
		return arr[headIdx:tailIdx]
	}
	midIdx := (headIdx + tailIdx) / 2
	arr1 := m.Split(arr, headIdx, midIdx)
	arr2 := m.Split(arr, midIdx, tailIdx)
	// fmt.Printf("arr1:%v, arr2:%v\n", arr1, arr2)
	newArr = m.Merge(arr1, arr2)
	// fmt.Printf("merge res:%v\n", newArr)
	return
}

func (m *MergeSort) Merge(arr1 []int, arr2 []int) (newArr []int) {
	newArr = make([]int, len(arr1)+len(arr2))
	var i, j, idx int
	for ; i < len(arr1) && j < len(arr2); idx++ {
		if arr1[i] < arr2[j] {
			newArr[idx] = arr1[i]
			i++
		} else {
			newArr[idx] = arr2[j]
			j++
		}
	}
	// 剩余数据加入
	for ; i < len(arr1); i++ {
		newArr[idx] = arr1[i]
		idx++
	}
	for ; j < len(arr2); j++ {
		newArr[idx] = arr2[j]
		idx++
	}
	return
}

// 不使用返回值，减少内存申请的方法
func (m *MergeSort) Merge2(arr []int, left, mid, right int, tmpArr []int) {
	var (
		i   = left
		j   = mid
		idx int
	)
	fmt.Printf("merge=>arr1:%v, arr2:%v\n", arr[left:mid], arr[mid:right])
	for ; i < mid && j < right; idx++ {
		if arr[i] < arr[j] {
			tmpArr[idx] = arr[i]
			i++
		} else {
			tmpArr[idx] = arr[j]
			j++
		}
	}
	// 剩余数据加入
	for ; i < mid; i++ {
		tmpArr[idx] = arr[i]
		idx++
	}
	for ; j < right; j++ {
		tmpArr[idx] = arr[j]
		idx++
	}
	// 拷贝元素回去arr，这里拷贝覆盖的部分只有[left:right]范围，
	// 从tmpArr取出的数据量夜只有right-left个，而不是整个tmpArr
	// 每次递归之后merge都可能改变之前的数据位置，因为递归在不断返回过程merge在不断扩大
	oldLeft := left
	for idx = 0; left < right; idx++ {
		arr[left] = tmpArr[idx]
		left++
	}
	fmt.Printf("tmp:%v copy range [%d:%d] to arr range [%d:%d], finally arr:%v\n", tmpArr, 0, right-oldLeft, oldLeft, right, arr)
}

func (m *MergeSort) Split2(arr []int, headIdx, tailIdx int, tmpArr []int) {
	if headIdx+1 == tailIdx {
		return
	}
	midIdx := (headIdx + tailIdx) / 2
	m.Split2(arr, headIdx, midIdx, tmpArr)
	m.Split2(arr, midIdx, tailIdx, tmpArr)
	m.Merge2(arr, headIdx, midIdx, tailIdx, tmpArr)
}

func main() {
	arr := common.GenRandInt(9, 100, 1)
	fmt.Println(arr)
	// t1 := time.Now()
	ms := MergeSort{}
	// newArr := ms.Split(arr, 0, len(arr))
	// fmt.Println(newArr)
	// fmt.Println(time.Now().Sub(t1))
	tmpArr := make([]int, len(arr)) // 临时数组的最长长度就是原数组长度，但每次使用是的有效长度都不一样
	ms.Split2(arr, 0, len(arr), tmpArr)
	fmt.Println(arr)
}
