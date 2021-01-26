package common

import (
	"math/rand"
	"time"
)

func GenRandInt(numCnt int, numRange int) (res []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numCnt; i++ {
		res = append(res, rand.Intn(numRange))
	}
	return
}
