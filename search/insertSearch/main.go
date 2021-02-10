package main

import "fmt"

type InsertSearch struct {
}

func (i *InsertSearch) Search(target int, arr []int, left, right int) (idx int) {
	idx = -1
	if left > right {
		return
	}
	mid := left + (target-arr[left])*(right-left)/(arr[right]-arr[left]) // 注意公式
	if mid > right || mid < left {                                       // 插值查找有可能会越界，这里需要判断退出
		return
	}
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		idx = i.Search(target, arr, left, mid-1)
	} else {
		idx = i.Search(target, arr, mid+1, right)
	}
	return
}

func main() {
	var arr []int
	for i := 0; i < 9; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	is := InsertSearch{}
	idx := is.Search(5, arr, 0, len(arr)-1)
	fmt.Println(idx)
	idx = is.Search(-10, arr, 0, len(arr)-1)
	fmt.Println(idx)
}
