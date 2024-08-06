package main

import (
	"fmt"
	"strings"
	"adventofcode/input"
)


func getTimeDist(lines []string) ([]int, []int) {
	line := strings.TrimSpace(lines[0][strings.Index(lines[0], ":")+1:])
	time := input.GetNumArr(line)
	line = strings.TrimSpace(lines[1][strings.Index(lines[1], ":")+1:])
	dist := input.GetNumArr(line)
	return time, dist
}

func getWays(time int, dist int) int {
	count := 0
	for t:=1;t<time;t++ {
		if (time-t)*t > dist {
			count++
		}
	}
	return count
}

func main1() {
	lines:= input.GetLines("./puzzle-input.txt")
	time, dist := getTimeDist(lines)
	fmt.Println(time, dist)

	answer := 1
	for i, _:= range time {
		answer *= getWays(time[i], dist[i])
	}
	fmt.Println("answer", answer)
}