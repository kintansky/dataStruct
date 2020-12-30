package main

import "fmt"

func zipArray(array interface{}) (sparseArray [][3]int) {
	orgArray := array.([][]int)
	sparseArray = [][3]int{[3]int{len(orgArray), len(orgArray[0]), 0}}
	for i, line := range orgArray {
		for j, data := range line {
			if data == 0 {
				continue
			}
			sparseArray = append(sparseArray, [3]int{i, j, data})
			sparseArray[0][2]++
		}
	}
	return
}

func unzipArray(sparseArray [][3]int) (orgArray [][]int) {
	n := sparseArray[0][0]
	m := sparseArray[0][1]
	for i := 0; i < n; i++ {
		defaultLine := make([]int, m)
		orgArray = append(orgArray, defaultLine)
	}
	for _, zipData := range sparseArray[1:] {
		orgArray[zipData[0]][zipData[1]] = zipData[2]
	}
	return
}

func main() {
	orgArray := [][]int{
		[]int{0, 0, 0, 1, 4},
		[]int{0, 20, 0, 0, 0},
		[]int{3, 0, 10, -1, 0},
	}
	fmt.Println(orgArray)
	sparseArray := zipArray(orgArray)
	fmt.Println(sparseArray)
	unzipArray := unzipArray(sparseArray)
	fmt.Println(unzipArray)
}
