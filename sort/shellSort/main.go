package main

import (
	"al/sort/common"
	"fmt"
	"time"
)

// ShellSort 希尔排序
type ShellSort struct {
	// 为了优化插入排序时，如果出现最坏情况需要整个有序表进行遍历的问题，提出了分组实现缩小增量
	// 1 按照步长step进行分组
	// 2 对分组进行插入排序
	// 3 缩小step继续分组+插入排序（step缩小后，分组会变少，但是组内元素增加），由于在此前的插入排序中，尽量减少每次出现最坏情况的几率
	// 4 结束条件，最后只剩一个分组，即步长step=1

}

// 0 1 2 3 4 5 6 7 8 9 10 11 12 共13个
// step = 6
// 0,6 1,7 2,8 3,9 4,10 5,11 12
// step = 6/2 = 3
// 0,3,6,9,12 1,4,7,10 2,5,8,11
// step = 3/2 = 1

// Sort 嵌套一个插入排序进行排序
func (s *ShellSort) Sort(arr []int, size int) {
	swapCnt := 0
	for step := len(arr) / size; step > 0; step /= size {
		for i := 0; i < step; i++ {
			// 每隔step取出一个值，进行插入排序
			for j := 1; i+step*j < len(arr); j++ {
				// 取出的值与前面step间隔的值，进行比较进行插入排序，注意运行条件curIdx > i
				for curIdx := i + step*j; curIdx > i; curIdx -= step {
					if arr[curIdx] >= arr[curIdx-step] {
						break
					}
					// fmt.Printf("i=%d, step=%d, swap {idx:%d, data:%d}<=>{idx:%d, data:%d}\n", i, step, curIdx, arr[curIdx], curIdx-step, arr[curIdx-step])
					arr[curIdx], arr[curIdx-step] = arr[curIdx-step], arr[curIdx] // 交换的位置是往前step位置的值
					// fmt.Println(arr)
					swapCnt++
				}
			}
		}
	}
	fmt.Println("swap count:", swapCnt)
}

// Sort2 对于数据量少的情况与Sort速度相当，但数据量达到10万以上开始严重劣化
func (s *ShellSort) Sort2(arr []int) {
	swapCnt := 0
	for step := len(arr) / 2; step > 0; step /= 2 {
		for i := step; i < len(arr); i++ {
			for j := i - step; j >= 0; j -= step {
				if arr[j] > arr[j+step] {
					arr[j], arr[j+step] = arr[j+step], arr[j]
					// fmt.Println(arr)
					swapCnt++
				}
			}
		}
	}
	fmt.Println("swap count:", swapCnt)
}

func main() {
	seed := time.Now().UnixNano()
	arr := common.GenRandInt(10, 1000000, seed)
	fmt.Println(time.Now())
	ss := ShellSort{}
	ss.Sort(arr, 2)
	fmt.Println(time.Now())

	arr = common.GenRandInt(10, 1000000, seed)
	fmt.Println(time.Now())
	// ss.Sort(arr, 10)
	ss.Sort2(arr)
	fmt.Println(time.Now())
}
