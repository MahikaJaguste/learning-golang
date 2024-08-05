package input

import (
	"os"
	"log"
	"bufio"
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