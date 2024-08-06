package input

import (
	"os"
	"log"
	"bufio"
	"strings"
)

func GetLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(file)

	lines := make([]string, 0, 10)
	
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		lines = append(lines, currLine)
	}

	return lines
}

func GetNumArr(line string) []int {
	line = strings.TrimSpace(line)
	numArr := make([]int, 0, 100)
	currNum := 0
	flag := false
	for _, v:= range line {
		if v >= '0' && v <= '9' {
			currNum = currNum*10 + int(v-'0')
			flag = true
		} else {
			if flag {
				numArr = append(numArr, currNum)
			}
			flag = false
			currNum = 0
		}
	}
	numArr = append(numArr, currNum)
	return numArr;
}