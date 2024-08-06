package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode/input"
)


func getTimeDistCombined(lines []string) (int, int) {
	line := strings.TrimSpace(lines[0][strings.Index(lines[0], ":")+1:])
	time := input.GetNumArr(line)

	line = strings.TrimSpace(lines[1][strings.Index(lines[1], ":")+1:])
	dist := input.GetNumArr(line)

	timeS := make([]string, len(line))
	distS := make([]string, len(dist))
	for i, v:= range time {
		timeS[i] = strconv.Itoa(v)
	}
	for i, v:= range dist {
		distS[i] = strconv.Itoa(v)
	}

	t, _ := strconv.Atoi(strings.Join(timeS, ""))
	d, _ := strconv.Atoi(strings.Join(distS, ""))
	return t,d
}

func satifies(time, dist, mid int) bool {
	return (time - mid) * mid > dist
}

func binarySearch(time int, dist int, lower bool) int {
	l := 1
	h := time - 1
	mid := l + (h-l)/2

	answer := 0
	for l <= h {
		mid = l + (h-l)/2
		isSatisfied := satifies(time, dist, mid)
		if isSatisfied {
			answer = mid
			if lower {
				h = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if lower {
				l = mid + 1
			} else {
				h = mid - 1
			}
		}
	}

	return answer
}


func main() {
	lines:= input.GetLines("./puzzle-input.txt")
	time, dist := getTimeDistCombined(lines)
	fmt.Println(time, dist)

	low := binarySearch(time, dist, true)
	fmt.Println("low", low)

	high := binarySearch(time, dist, false)
	fmt.Println("high", high)

	answer := high - low + 1
	fmt.Println("answer", answer)
}