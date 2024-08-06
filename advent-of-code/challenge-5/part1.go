package main 

import (
	"fmt"
	"strings"
	"adventofcode/input"

)

func getSeeds(line string) []int {
	line = strings.TrimSpace(line[strings.Index(line, ":")+1:])
	return input.GetNumArr(line)
}

func processSet(lineSet []string, currList []int) {
	lineSetNums := make([][]int, 0, 100)
	for _, v := range lineSet {
		lineSetNums = append(lineSetNums, input.GetNumArr(v))
	}

	for index, num := range currList {
		for _, lineSetNumsArr := range lineSetNums {
			if num >= lineSetNumsArr[1] && num < lineSetNumsArr[1] + lineSetNumsArr[2] {
				currList[index] = lineSetNumsArr[0] + num - lineSetNumsArr[1]
				break
			}
		}
	}
}

func processMappings(lines []string, currList []int) []int {

	lineSet := make([]string, 0, 100)
	for _, line := range lines {
		if len(line) == 0 {
			processSet(lineSet, currList)
			lineSet = make([]string, 0, 100)
			continue
		}
		if !strings.Contains(line, "map:") {
			lineSet = append(lineSet, line)
		}
	}

	processSet(lineSet, currList)
	return currList
}

func getMin(finalList []int) int {
	minVal := finalList[0]
	for _, v := range finalList {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func main() {
	lines := input.GetLines("./puzzle-input.txt")

	seeds := getSeeds(lines[0])
	fmt.Println("seeds", seeds)

	finalList := processMappings(lines[2:], seeds)
	fmt.Println("finalList", finalList)

	minVal := getMin(finalList)
	fmt.Println("minVal", minVal)
}