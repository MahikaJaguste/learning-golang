package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode/input"
)

func getArrayFromString(numString string, numArr *[]int) {
	for strings.Contains(numString, " ") {
		i := strings.Index(numString, " ")
		currNum, _ := strconv.Atoi(numString[:i])
		numString = strings.TrimSpace(numString[i+1:])
		*numArr = append(*numArr, currNum)
	}
	currNum, _ := strconv.Atoi(numString)
	*numArr = append(*numArr, currNum)
}

func getArrays(line string) (*[]int, *[]int) {
	winningNumbers := strings.TrimSpace(line[strings.Index(line, ":") + 2 : strings.Index(line, "|") - 1])
	myNumbers := strings.TrimSpace(line[strings.Index(line, "|") + 2 : ])

	winningNumbersArr := make([]int, 0, 20)
	myNumbersArr := make([]int, 0, 50)

	getArrayFromString(winningNumbers, &winningNumbersArr)
	getArrayFromString(myNumbers, &myNumbersArr)

	return &winningNumbersArr, &myNumbersArr
}

func getCardValue(winningNumbersArr []int, myNumbersArr []int) int {
	mp := make(map[int]bool, 10)
	for _, v:= range winningNumbersArr {
		mp[v] = true
	}

	product := 0

	for _, v:= range myNumbersArr {
		if mp[v] {
			if product == 0 {
				product = 1
			} else {
				product *= 2
			}
		}
	}

	return product
}




func main1() {
	lines := input.GetLines("./puzzle-input.txt")

	answer := 0
	for _, line := range lines {
		winningNumbersArr, myNumbersArr := getArrays(line)
		answer += getCardValue(*winningNumbersArr, *myNumbersArr)
	}

	fmt.Println("answer:", answer)
}