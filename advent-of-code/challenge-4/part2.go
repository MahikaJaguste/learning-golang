package main

import (
	"fmt"
	"adventofcode/input"
)

func getCardMatchingEntries(winningNumbersArr []int, myNumbersArr []int) int {
	mp := make(map[int]bool, 10)
	for _, v:= range winningNumbersArr {
		mp[v] = true
	}

	matches := 0
	for _, v:= range myNumbersArr {
		if mp[v] {
			matches++
		}
	}
	return matches
}

func updateCopies(index int, matches int, copiesArr []int) {
	for i := index+1; i<= index+matches; i++ {
		copiesArr[i] += copiesArr[index]
	}
}


func main() {
	lines := input.GetLines("./puzzle-input.txt")

	copiesArr := make([]int, len(lines))
	for index, line := range lines {
		copiesArr[index] += 1;
		winningNumbersArr, myNumbersArr := getArrays(line)
		matches := getCardMatchingEntries(*winningNumbersArr, *myNumbersArr)
		updateCopies(index, matches, copiesArr)
	}

	answer := 0
	for _, v:= range copiesArr {
		answer += v
	}

	fmt.Println("answer:", answer)
}