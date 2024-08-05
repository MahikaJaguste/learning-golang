package main

import (
	"fmt"
	"strings"
	"strconv"
	
	"adventofcode/input"
)


func getSetCounts(currLine string) [3]int {
	r := [3]int{}

	mp := map[string]int {
		"red": 0,
		"green": 1,
		"blue": 2,
	}

	for strings.Contains(currLine, ",") {
		spaceIndex := strings.Index(currLine, " ")
		count, _ := strconv.Atoi(currLine[:spaceIndex])
		 
		commaIndex := strings.Index(currLine, ",")
		color := currLine[spaceIndex+1: commaIndex]

		r[mp[color]] = count
		
		currLine = currLine[commaIndex+2:]
	}

	spaceIndex := strings.Index(currLine, " ")
	count, _ := strconv.Atoi(currLine[:spaceIndex])
	color := currLine[spaceIndex+1: ]

	r[mp[color]] = count

	return r
}

func getGameProduct(line string) int {
	product := 1
	maxValues := [3]int{0,0,0}

	colonIndex := strings.Index(line, ":")
	line = line[colonIndex+2:]

	for strings.Contains(line, ";") {
		semiColonIndex := strings.Index(line, ";")
		currLine := line[:semiColonIndex]
		line = line[semiColonIndex+2:]

		r := getSetCounts(currLine)
		for i, _:= range r {
			if maxValues[i] < r[i] {
				maxValues[i] = r[i]
			}
		}
	}

	r := getSetCounts(line)
	for i, _:= range r {
		if maxValues[i] < r[i] {
			maxValues[i] = r[i]
		}
	}

	for _, v:= range maxValues {
		product *= v
	}

	return product
}

func solvePart2(lines []string) int {
	answer := 0

	for _, s := range lines {
		answer += getGameProduct(s)		
	}

	return answer
	
}


func main() {
	
	lines := input.GetLines("./puzzle-input.txt")
	// fmt.Println("lines:", lines)

	answer := solvePart2(lines)
	fmt.Println("answer:", answer)
}