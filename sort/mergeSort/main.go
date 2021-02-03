package main

import (
	"al/sort/common"
	"fmt"
)

type MergeSort struct {
}

func (m *MergeSort) Sort(arr []int, headIdx, tailIdx int) (newArr []int) {
	if headIdx+1 == tailIdx {
		return arr[headIdx:tailIdx]
	}
	midIdx := (headIdx + tailIdx) / 2
	arr1 := m.Sort(arr, headIdx, midIdx)
	arr2 := m.Sort(arr, midIdx, tailIdx)
	fmt.Printf("arr1:%v, arr2:%v\n", arr1, arr2)
	newArr = m.Merge(arr1, arr2)
	fmt.Printf("merge res:%v\n", newArr)
	return
}

func (m *MergeSort) Merge(arr1 []int, arr2 []int) (newArr []int) {
	newArr = make([]int, len(arr1)+len(arr2))
	var i, j int
	for idx := 0; idx < len(newArr); idx++ {
		if i >= len(arr1) && j >= len(arr2) {
			return
		} else if j >= len(arr2) {
			for ; i < len(arr1); i++ {
				newArr[idx] = arr1[i]
				idx++
			}
			return
		} else if i >= len(arr1) {
			for ; j < len(arr2); j++ {
				newArr[idx] = arr2[j]
				idx++
			}
			return
		}
		if arr1[i] < arr2[j] {
			newArr[idx] = arr1[i]
			i++
		} else {
			newArr[idx] = arr2[j]
			j++
		}
	}
	return
}

func main() {
	arr := common.GenRandInt(10, 100, 1)
	fmt.Println(arr)
	ms := MergeSort{}
	newArr := ms.Sort(arr, 0, len(arr))
	fmt.Println(newArr)
	// testArr := ms.Merge([]int{2, 4, 5}, []int{1, 3, 7, 9})
	// fmt.Println(testArr)
}
