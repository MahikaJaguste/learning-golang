package main

import (
	"fmt"
	"adventofcode/input"
)

var rows, cols int

func inRange(r, c int) bool {
	return r >= 0 && r < rows && c >= 0 && c < cols
}

func getSymbolMatrix(lines []string) [][]bool {
	symbolMatrix := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		symbolMatrix[i] = make([]bool, cols)
	}

	for index, line := range lines {
		for i, v := range line {
			isSymbol := !((v >= '0' && v <= '9') || v == '.')

			if !isSymbol {
				continue
			}
			
			// plus
			if inRange(index, i-1) {
				symbolMatrix[index][i-1] = true
			}
			if inRange(index, i+1) {
				symbolMatrix[index][i+1] = true
			}
			if inRange(index-1, i) {
				symbolMatrix[index-1][i] = true
			}
			if inRange(index+1, i) {
				symbolMatrix[index+1][i] = true
			}
			
			// diagonal
			if inRange(index-1, i-1) {
				symbolMatrix[index-1][i-1] = true
			}
			if inRange(index-1, i+1) {
				symbolMatrix[index-1][i+1] = true
			}
			if inRange(index+1, i-1) {
				symbolMatrix[index+1][i-1] = true
			}
			if inRange(index+1, i+1) {
				symbolMatrix[index+1][i+1] = true
			}
		} 
	}

	return symbolMatrix
}

func printSymbolMatrix(symbolMatrix [][]bool, lines []string) {
	fmt.Println("symbolMatrix:")
	for index, line := range lines {
		for i, _ := range line {
			fmt.Print(symbolMatrix[index][i], " ")
		} 
		fmt.Println()
	}
}

func getPartNumbersSum(symbolMatrix [][]bool, lines []string) int {
	answer := 0

	for index, line := range lines {
		isCurrNumEligible := false
		currNum := 0
		for i, v := range line {
			if v >= '0' && v <= '9' {
				currNum = currNum*10 + int(v - '0')
				isCurrNumEligible = isCurrNumEligible || symbolMatrix[index][i]
			} else {
				if isCurrNumEligible {
					answer += currNum
				}
				currNum = 0
				isCurrNumEligible = false
			}
		} 
		if isCurrNumEligible {
			answer += currNum
		}
	}

	return answer

}


func main1() {
	lines := input.GetLines("./puzzle-input.txt")

	rows = len(lines)
	cols = len(lines[0])
	fmt.Println("rows:", rows, "; cols:", cols)

	symbolMatrix := getSymbolMatrix(lines)
	// printSymbolMatrix(symbolMatrix, lines)

	answer := getPartNumbersSum(symbolMatrix, lines)
	fmt.Println("answer:", answer)

}