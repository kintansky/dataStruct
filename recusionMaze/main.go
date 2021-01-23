package main

import "fmt"

func createMaze(col, row int) (maze [][]int) {
	for i := 0; i < col; i++ {
		colData := make([]int, row)
		switch {
		case i == 0 || i == col-1:
			for idx, _ := range colData {
				colData[idx] = 1
			}
		case i == 2:
			colData[3] = 1
		case i == 3:
			for idx, _ := range colData {
				if idx < 4 {
					colData[idx] = 1
					continue
				}
				break
			}
		}
		colData[0] = 1
		colData[row-1] = 1
		maze = append(maze, colData)
	}
	return
}

func showMaze(maze [][]int) {
	for _, colData := range maze {
		fmt.Println(colData)
	}
}

func findPath(x, y int, targetPos [2]int, maze [][]int) {
	// 逆时针寻路，走过的位置置为2，死路置为4
	if maze[x][y] == 0 {
		maze[x][y] = 2
	} else {
		maze[x][y] = 4
	}
	showMaze(maze)
	if x == targetPos[0] && y == targetPos[1] {
		fmt.Println("success!")
		return
	}
	fmt.Println("========================================")
	switch {
	case maze[x-1][y] != 1 && maze[x-1][y] != 4: // 向上寻路
		fmt.Printf("up: (%d, %d)\n", x-1, y)
		findPath(x-1, y, targetPos, maze)
	case maze[x][y-1] != 1 && maze[x][y-1] != 4: // 向左寻路
		fmt.Printf("left: (%d, %d)\n", x, y-1)
		findPath(x, y-1, targetPos, maze)
	case maze[x+1][y] != 1 && maze[x+1][y] != 4: //向下寻路
		fmt.Printf("down: (%d, %d)\n", x+1, y)
		findPath(x+1, y, targetPos, maze)
	case maze[x][y+1] != 1 && maze[x][y+1] != 4: //向右寻路
		fmt.Printf("right: (%d, %d)\n", x, y+1)
		findPath(x, y+1, targetPos, maze)
	default:
		fmt.Println("unreachable!")
	}
}

func main() {
	col := 7
	row := 8
	maze := createMaze(col, row)
	showMaze(maze)
	// 起始位置（1，1），目的位置（col-1， row-1）
	fmt.Println()
	findPath(1, 1, [2]int{col - 2, row - 2}, maze)

}
