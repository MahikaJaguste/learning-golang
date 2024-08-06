package main

import (
	"fmt"
	"adventofcode/input"
)

type node struct {
	data string
	left *node
	right *node
}

var mp map[string]*node

func getNodeData(line string) (string, string, string) {
	elem := line[:3]
	l := line[7:10]
	r := line[12:15]
	return elem, l, r
}

func formNodes(lines []string) {
	mp = make(map[string]*node)

	for _, line := range lines {
		elem, l, r := getNodeData(line)

		if _, ok := mp[elem]; !ok {
			mp[elem] = &node{data: elem}
		}

		if _, ok := mp[l]; !ok {
			mp[l] = &node{data: l}
		}

		if _, ok := mp[r]; !ok {
			mp[r] = &node{data: r}
		}

		currNode := mp[elem]
		currNode.left = mp[l]
		currNode.right = mp[r]
		mp[elem] = currNode
	}
}


func performTraversal(directionLine string) int {
	index := 0
	count := 0
	currNode := mp["AAA"]
	for true {
		count++

		if index >= len(directionLine) {
			index = 0
		}

		if directionLine[index] == 'R' {
			currNode = currNode.right
		} else {
			currNode = currNode.left
		}

		if currNode.data == "ZZZ" {
			break
		}

		index++
	}

	return count
}

func main1() {
	lines := input.GetLines("./puzzle-input.txt")

	directionLine := lines[0]
	formNodes(lines[2:])

	answer := performTraversal(directionLine)
	fmt.Println("answer", answer)
}