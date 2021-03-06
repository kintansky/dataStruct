package main

import (
	"al/sort/common"
	"fmt"
	"time"
)

type RadixSort struct {
	// 依次遍历数组，按照个进位有低至高（个十百千万...）进行分类
	// 1 先依次按照个位，把所有数据依次放入[0-9]的桶
	// 1.1 遍历[0-9]桶，取回数据放回arr
	// 2 按照十位，重复以上步骤，如果<10的，归为0
	// 3 最后完成的结束条件：绝对值最大的数的最高位也完成了分类
	bucket           [10][]int // bucket,存放数据的桶
	bucketDataCntArr [10]int   // 每个桶有效数据长度
}

func NewRadixSort(size int) *RadixSort {
	var bucket [10][]int
	for i := 0; i < 10; i++ {
		tmpArr := make([]int, size)
		bucket[i] = tmpArr
	}
	return &RadixSort{
		bucket: bucket,
	}
}

func (r *RadixSort) Sort(arr []int) {
	var radixIdx, arrIdx int
	var haveBiggerData bool
	for pos := 1; ; pos *= 10 {
		haveBiggerData = false
		for _, data := range arr {
			// 正负数分开处理
			// 1 正数采用升序填入bucket
			// 2 负数采用绝对值的升序填入bucket（填入时仍然时负数本身而不是绝对值）
			// 3 到了最后一轮填入bucket的时候，正数以升序集中在bucket0, 负数以升序集中在bucket9
			if data >= 0 {
				if data > pos {
					haveBiggerData = true
				}
				radixIdx = data / pos % 10
				r.bucket[radixIdx][r.bucketDataCntArr[radixIdx]] = data
				r.bucketDataCntArr[radixIdx]++
			} else {
				if -1*data > pos {
					haveBiggerData = true
				}
				radixIdx = -1 * data / pos % 10
				r.bucket[9-radixIdx][r.bucketDataCntArr[9-radixIdx]] = data // 注意这里
				r.bucketDataCntArr[9-radixIdx]++                            // 注意这里
			}
		}
		arrIdx = 0
		// 最后一次取出数据放回arr时，因为负数集中在bucket9,正数集中在bucket0，所以先取bucket9最后才取出bucket0
		if !haveBiggerData {
			for i := len(r.bucketDataCntArr) - 1; i >= 0; i-- {
				for j := 0; j < r.bucketDataCntArr[i]; j++ {
					arr[arrIdx] = r.bucket[i][j]
					arrIdx++
				}
			}
			// fmt.Println("radix:", arr)
			break
		}
		// 不是最后一次取出的时候，正常按顺序回填即可
		for i, cnt := range r.bucketDataCntArr {
			if cnt == 0 {
				continue
			}
			// fmt.Println(i, r.bucket[i])
			for j := 0; j < cnt; j++ {
				arr[arrIdx] = r.bucket[i][j]
				arrIdx++
			}
			r.bucketDataCntArr[i] = 0
		}
		// fmt.Println("radix:", arr)
	}
}

func main() {
	// arr1 := common.GenRandInt(5, 100, 1)
	// for i := range arr1 {
	// 	arr1[i] = -1 * arr1[i]
	// }
	// arr2 := common.GenRandInt(4, 100, 2)
	// arr := append(arr1, arr2...)
	// arr := []int{1, -1, 2, -7, -8, 0}
	arr := common.GenRandInt(100000, 1000000, time.Now().UnixNano())
	// fmt.Println(arr)
	rs := NewRadixSort(len(arr))
	t := time.Now()
	rs.Sort(arr)
	fmt.Println(time.Now().Sub(t))
}
