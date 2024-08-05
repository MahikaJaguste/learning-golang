package main

import (
	"fmt"
	"strings"
	"strconv"
	
	"adventofcode/input"
)

func isColorCountSatisfied(count int, color string) bool {
	if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
		return false
	}
	return true
}

func processSet(currLine string) bool {
	result := true

	for strings.Contains(currLine, ",") {
		spaceIndex := strings.Index(currLine, " ")
		count, _ := strconv.Atoi(currLine[:spaceIndex])
		 
		commaIndex := strings.Index(currLine, ",")
		color := currLine[spaceIndex+1: commaIndex]

		if !isColorCountSatisfied(count, color) {
			result = false
			break
		}
		
		currLine = currLine[commaIndex+2:]
	}

	spaceIndex := strings.Index(currLine, " ")
	count, _ := strconv.Atoi(currLine[:spaceIndex])
	color := currLine[spaceIndex+1: ]

	if !isColorCountSatisfied(count, color) {
		result = false
	}

	return result
}

func isValidGame(line string) (int, bool) {
	var (
		gameNum int 
		isValid bool = true
	)

	colonIndex := strings.Index(line, ":")
	gameNum, _ = strconv.Atoi(line[5:colonIndex])

	line = line[colonIndex+2:]

	for strings.Contains(line, ";") {
		semiColonIndex := strings.Index(line, ";")
		currLine := line[:semiColonIndex]
		line = line[semiColonIndex+2:]

		isValidSet := processSet(currLine)
		if !isValidSet {
			isValid = false
			break
		}

	}

	isValidSet := processSet(line)
	if !isValidSet {
		isValid = false	
	}

	return gameNum, isValid
}

func solve(lines []string) int {
	answer := 0

	for _, s := range lines {
		gameNum, isValid := isValidGame(s)
		if isValid {
			answer += gameNum
		}
	}

	return answer
	
}


func main1() {
	
	lines := input.GetLines("./puzzle-input.txt")
	// fmt.Println("lines:", lines)

	answer := solve(lines)
	fmt.Println("answer:", answer)

}