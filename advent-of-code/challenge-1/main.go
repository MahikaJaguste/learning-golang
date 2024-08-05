package main

import (
	"fmt"
	"strings"
	"strconv"
	"slices"
	"math"
	
	"adventofcode/input"
)

func getNums(lines []string) []int {
	nums := make([]int, 0, len(lines))

	for _, s := range lines {
		chars := strings.Split(s, "")

		var firstDigit, lastDigit int

		for _, v := range chars {
			if v >= "0" && v <= "9" {
				firstDigit, _ = strconv.Atoi(v)
				break
			}
		}

		slices.Reverse(chars)
		for _, v := range chars {
			if v >= "0" && v <= "9" {
				lastDigit, _ = strconv.Atoi(v)
				break
			}
		}

		newNum := firstDigit * 10 + lastDigit
		nums = append(nums, newNum)
	}

	return nums
	
}

func getNumAdvanced(s string) int {

	validNums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	validNumsIndex := make([]int, 9, 9)
	for i, v := range validNums {
		validNumsIndex[i] = strings.Index(s, v)
	}

	validNIndex := make([]int, 10, 10)
	for i:= 0; i<= 9; i++ {
		validNIndex[i] = strings.Index(s, strconv.Itoa(i))
	}

	var firstDigit int
	
	var minIndex, minValue int
	minValue = math.MaxInt64

	flag := false
	for i, v := range validNumsIndex {
		if v < minValue && v > -1 {
			flag = true
			minValue = v
			minIndex = i
		}
	}

	if flag {
		firstDigit = minIndex + 1
	}

	flag = false
	for i, v := range validNIndex {
		if v < minValue && v > -1 {
			flag = true
			minValue = v
			minIndex = i
		}
	}

	if flag {
		firstDigit = minIndex
	}

	validNumsLastIndex := make([]int, 9, 9)
	for i, v := range validNums {
		validNumsLastIndex[i] = strings.LastIndex(s, v)
	}

	validNLastIndex := make([]int, 10, 10)
	for i:= 0; i<= 9; i++ {
		validNLastIndex[i] = strings.LastIndex(s, strconv.Itoa(i))
	}

	var lastDigit int
	
	var maxIndex, maxValue int
	maxValue = math.MinInt64

	maxFlag := false
	for i, v := range validNumsLastIndex {
		if v > maxValue && v > -1 {
			maxFlag = true
			maxValue = v
			maxIndex = i
		}
	}

	if maxFlag {
		lastDigit = maxIndex + 1
	}

	maxFlag = false
	for i, v := range validNLastIndex {
		if v > maxValue && v > -1 {
			maxFlag = true
			maxValue = v
			maxIndex = i
		}
	}

	if maxFlag {
		lastDigit = maxIndex
	}
	
	return firstDigit*10 + lastDigit

}

func getNumsAdvanced(lines []string) []int {

	nums := make([]int, 0, len(lines))

	for _, s := range lines {
		newNum := getNumAdvanced(s)
		nums = append(nums, newNum)
	}

	return nums
}

func sumNums(nums []int) int {
	var answer int
	for _, v := range nums {
		answer += v
	}
	return answer
}

func main() {
	
	lines := input.GetLines("./puzzle-input.txt")
	// fmt.Println("lines:", lines)

	// nums := getNums(lines)
	nums := getNumsAdvanced(lines)
	// fmt.Println("nums:", nums)

	answer := sumNums(nums)
	fmt.Println("answer:", answer)


}