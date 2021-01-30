package common

import (
	"math/rand"
)

// GenRandInt 一个seed对应一个一样的数组
func GenRandInt(numCnt int, numRange int, seed int64) (res []int) {
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(seed)
	for i := 0; i < numCnt; i++ {
		res = append(res, rand.Intn(numRange))
	}
	return
}
