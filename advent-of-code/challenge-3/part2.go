package main

import (
	"fmt"
	"adventofcode/input"
)


func getNumbersMatrix(lines []string) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = -1
		}
	}

	for index, line := range lines {
		currNum := 0
		startIndex := -1
		for i, ch := range line {
			if ch >= '0' && ch <= '9' {
				currNum = currNum*10 + int(ch - '0')
				if startIndex == -1 {
					startIndex = i
				}
			} else {
				if startIndex != -1 {
					for k := startIndex; k < i; k++ {
						matrix[index][k] = currNum
					}
				}
				currNum = 0
				startIndex = -1
			}
		}
		if startIndex != -1 {
			for k := startIndex; k < len(line); k++ {
				matrix[index][k] = currNum
			}
		}
	}

	return matrix
}

func solveGears(lines []string, matrix [][]int) int {
	answer := 0

	for index, line := range lines {
		for i, ch := range line {
			if ch != '*' {
				continue
			}

			nums := make([]int, 0, 10)

			hasU := inRange(index-1, i)
			if hasU { 
				if matrix[index-1][i] != -1 {
					nums = append(nums, matrix[index-1][i])
				} else {
					hasLeftNum := inRange(index-1, i-1) && matrix[index-1][i-1] != -1
					hasRightNum := inRange(index-1, i+1) && matrix[index-1][i+1] != -1
					if hasLeftNum {
						nums = append(nums,  matrix[index-1][i-1])
					}
					if hasRightNum {
						nums = append(nums,  matrix[index-1][i+1])
					}
				}
			}

			hasB := inRange(index+1, i)
			if hasB { 
				if matrix[index+1][i] != -1 {
					nums = append(nums, matrix[index+1][i])
				} else {
					hasLeftNum := inRange(index+1, i-1) && matrix[index+1][i-1] != -1
					hasRightNum := inRange(index+1, i+1) && matrix[index+1][i+1] != -1
					if hasLeftNum {
						nums = append(nums,  matrix[index+1][i-1])
					}
					if hasRightNum {
						nums = append(nums,  matrix[index+1][i+1])
					}
				}
			}

			if inRange(index, i-1) && matrix[index][i-1] != -1 {
				nums = append(nums, matrix[index][i-1])
			}

			if inRange(index, i+1) && matrix[index][i+1] != -1 {
				nums = append(nums, matrix[index][i+1])
			}

			if len(nums) != 2 {
				continue
			}

			answer += nums[0] * nums[1]
		}
	}

	return answer
}


func main() {
	lines := input.GetLines("./puzzle-input.txt")

	rows = len(lines)
	cols = len(lines[0])
	fmt.Println("rows:", rows, "; cols:", cols)

	matrix := getNumbersMatrix(lines)
	// fmt.Println("matrix:", matrix)

	answer := solveGears(lines, matrix)
	fmt.Println("answer:", answer)
}