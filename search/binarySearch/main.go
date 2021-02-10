package main

import "fmt"

type BinarySearch struct {
}

func (b *BinarySearch) Search(target int, arr []int, left, right int) (idx int) {
	idx = -1
	if left > right || target < arr[0] || target > arr[len(arr)-1] {
		return
	}
	mid := (left + right) / 2
	if target == arr[mid] {
		return mid
	}
	if target > arr[mid] {
		idx = b.Search(target, arr, mid+1, right)
	} else {
		idx = b.Search(target, arr, left, mid-1)
	}
	return
}

// 查找所有目标值的index
func (b *BinarySearch) SearchAllIndex(target int, arr []int, left, right int, idxArr *[]int) {
	if left > right || target < arr[0] || target > arr[len(arr)-1] {
		return
	}
	mid := (left + right) / 2
	// fmt.Println(left, mid, right, arr[mid])
	if target == arr[mid] {
		*idxArr = append(*idxArr, mid)
		// fmt.Println(idxArr)
	}
	b.SearchAllIndex(target, arr, mid+1, right, idxArr)
	b.SearchAllIndex(target, arr, left, mid-1, idxArr)
}

func main() {
	var arr []int
	for i := 0; i < 9; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	bs := BinarySearch{}
	idx := bs.Search(5, arr, 0, len(arr)-1)
	fmt.Println(idx)
	idx = bs.Search(50, arr, 0, len(arr)-1)
	fmt.Println(idx)

	arr = append(arr, 8)
	arr[7] = 8
	fmt.Println(arr)
	var idxArr1 []int
	bs.SearchAllIndex(8, arr, 0, len(arr)-1, &idxArr1)
	fmt.Println(idxArr1)
	var idxArr2 []int
	bs.SearchAllIndex(90, arr, 0, len(arr)-1, &idxArr2)
	fmt.Println(idxArr2)

}
