package main

import (
	"al/sort/common"
	"fmt"
	"time"
)

type RadixSort struct {
	bucket           [10][]int
	bucketDataCntArr [10]int
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

func (r *RadixSort) cleanBucketDataCntArr() {
	for i, _ := range r.bucketDataCntArr {
		r.bucketDataCntArr[i] = 0
	}
}

func (r *RadixSort) Sort(arr []int) {
	var radixIdx, arrIdx int
	var haveBiggerData bool
	for pos := 1; ; pos *= 10 {
		haveBiggerData = false
		for _, data := range arr {
			if data > pos {
				haveBiggerData = true
			}
			radixIdx = data / pos % 10
			r.bucket[radixIdx][r.bucketDataCntArr[radixIdx]] = data
			r.bucketDataCntArr[radixIdx]++
		}
		arrIdx = 0
		for j, tmpArr := range r.bucket {
			// fmt.Println(j, tmpArr)
			for i := 0; i < r.bucketDataCntArr[j]; i++ {
				arr[arrIdx] = tmpArr[i]
				arrIdx++
			}
		}
		// fmt.Println("radix:", arr)
		if !haveBiggerData {
			break
		}
		r.cleanBucketDataCntArr()
	}
}

func main() {
	arr := common.GenRandInt(100000, 1000000, 1)
	// fmt.Println(arr)
	rs := NewRadixSort(len(arr))
	t := time.Now()
	rs.Sort(arr)
	fmt.Println(time.Now().Sub(t))
}
