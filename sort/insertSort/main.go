package main

import (
	"al/sort/common"
	"fmt"
)

type InsertSort struct {
	// 时间复杂度O(n^2)
	// 1 将数组分成一个有序表，一个待排序的无序表
	// 2 第一次直接把第一个元素放入有序表，其余作为无序表
	// 3 后续将无序表的元素取出与有序表的最后一个元素比较
	// 3.1 如果符合，则不动，继续进行下一个插入无需元素的插入比较
	// 3.2 如果不符合，先交换两者位置，则该元素继续跟它在有序表的前一个元素进行比较，直到符合
}

func (i *InsertSort) Sort(arr []int) {
	for idx, _ := range arr {
		if idx == 0 {
			continue
		}
		for j := idx; j > 0; j-- {
			if arr[j] >= arr[j-1] {
				break
			}
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

func main() {
	arr := common.GenRandInt(10, 100)
	fmt.Println(arr)

	is := InsertSort{}
	is.Sort(arr)
	fmt.Println(arr)
}
