package main

import "fmt"

type FibSearch struct{}

func (f *FibSearch) Search(target int, arr []int) (idx int) {
	idx = -1
	var (
		low  = 0
		high = len(arr) - 1
		mid  = 0
		k    = 0
	)
	// 1 生成一个斐波那契数列
	fibArr := fib(len(arr))
	// fmt.Println("fib:", fibArr)
	// 2 找到满足数值大于数组长度的斐波那契数列下标
	for fibArr[k]-1 < len(arr) {
		k++
	}
	// 3 数组填充，使用数组末尾数据进行填充
	appendArr := make([]int, fibArr[k]-len(arr))
	for idx := range appendArr {
		appendArr[idx] = arr[high]
	}
	newArr := append(arr, appendArr...)
	// fmt.Println("newArr:", newArr)
	// 4 不断进行划分，F(k) = F(k-1)+F(k-2)，分成前半为F(k-1)个，后半是F(k-2)个
	for low <= high {
		mid = low + fibArr[k-1] - 1
		if target < newArr[mid] { // 前半F(k-1)
			high = mid - 1
			k -= 1
		} else if target > newArr[mid] { // 后半F(k-2)
			low = mid + 1
			k -= 2
		} else {
			if mid < len(arr) {
				idx = mid
			} else {
				idx = len(arr) - 1
			}
			break
		}
	}
	return
}

func fib(size int) (fibArr []int) {
	fibArr = make([]int, size)
	fibArr[0] = 1
	fibArr[1] = 1
	for i := 2; i < size; i++ {
		fibArr[i] = fibArr[i-1] + fibArr[i-2]
	}
	return
}

func main() {
	arr := []int{
		1, 2, 4, 6, 7, 9, 13, 16, 17, 21, 23, 25, 40, 45,
	}
	fmt.Println(arr, len(arr))
	fs := FibSearch{}
	idx := fs.Search(9, arr)
	fmt.Println(idx)
	idx = fs.Search(1, arr)
	fmt.Println(idx)
	idx = fs.Search(45, arr)
	fmt.Println(idx)
}
