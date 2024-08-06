package main

import (
	"fmt"
	"strconv"
	"sort"
	"adventofcode/input"
)

func getCardsArr(lines []string) ([]string, []int) {
	cards := make([]string, len(lines))
	bids := make([]int, len(lines))

	for i, v := range lines {
		cards[i] = v[:5]
		bids[i], _ = strconv.Atoi(v[6:])
	}

	return cards, bids
}

func getCharIndex(ch byte) int {
	if ch >= '2' && ch <= '9' {
		return int(ch - '0')
	} else if ch == 'T' {
		return 10
	} else if ch == 'J' {
		return 11
	} else if ch == 'Q' {
		return 12
	} else if ch == 'K' {
		return 13
	} else {
		return 14
	}
}

func getTypeIndex(card string) int {
	freq := make([]int, 15)
	for i:=0; i<5; i++ {
		index := getCharIndex(card[i])
		freq[index] += 1
	}

	times := make([]int, 6)
	for i:=2;i<15;i++ {
		times[freq[i]] += 1
	}

	if times[5] == 1 { 
		return 6
	} else if times[4] == 1 {
		return 5
	} else if times[3] == 1 && times[2] == 1 {
		return 4
	} else if times[3] == 1 && times[1] == 2 {
		return 3
	} else if times[2] == 2 {
		return 2
	} else if times[2] == 1 && times[1] == 3 {
		return 1
	} else {
		return 0
	}
}


func main1() {
	lines := input.GetLines("./puzzle-input.txt")

	cards, bids := getCardsArr(lines)
	fmt.Println(cards, bids)

	mp := make(map[string]int)
	for i, v:= range cards {
		mp[v] = bids[i]
	}

	cardTypes := make([][]string, 7)
	for _, v:= range cards {
		index := getTypeIndex(v)
		cardTypes[index] = append(cardTypes[index], v)
	}

	for _,v := range cardTypes {
		sort.SliceStable(v, func(i, j int) bool {
			a := v[i]
			b := v[j]
			var ch1, ch2 int
			for i:=0; i<5; i++ {
				ch1 = getCharIndex(a[i])
				ch2 = getCharIndex(b[i])
				if (ch1 == ch2) {
					continue
				}
				break
			}
			return ch1 < ch2
		})
	}

	answer := 0
	index := 1
	for _,cardType := range cardTypes {
		for _,v := range cardType {
			answer += mp[v] * index
			index++
		}
	}

	fmt.Println("answer", answer)
}