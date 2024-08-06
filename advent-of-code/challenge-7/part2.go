package main

import (
	"fmt"
	"sort"
	"adventofcode/input"
)


func getCharIndexAdvanced(ch byte) int {
	if ch >= '2' && ch <= '9' {
		return int(ch - '0')
	} else if ch == 'T' {
		return 10
	} else if ch == 'J' {
		return 1
	} else if ch == 'Q' {
		return 11
	} else if ch == 'K' {
		return 12
	} else {
		return 13
	}
}

func getTypeIndexAdvanced(card string) int {
	freq := make([]int, 14)
	for i:=0; i<5; i++ {
		index := getCharIndexAdvanced(card[i])
		freq[index] += 1
	}

	maxFreq := 1;
	times := make([]int, 6)
	for i:=2;i<len(freq);i++ {
		times[freq[i]] += 1
		if freq[i] > maxFreq {
			maxFreq = freq[i]
		}
	}

	jTimes := freq[1]
	if(jTimes == 5) {
		return 6
	}

	times[maxFreq] -= 1
	times[maxFreq + jTimes] += 1


	if times[5] == 1 { 
		return 6
	} else if times[4] == 1 {
		return 5
	} else if times[3] == 1 && times[2] == 1 {
		return 4
	} else if times[3] == 1 {
		return 3
	} else if times[2] == 2 {
		return 2
	} else if times[2] == 1 {
		return 1
	} else {
		return 0
	}
}


func main() {
	lines := input.GetLines("./puzzle-input.txt")

	cards, bids := getCardsArr(lines)
	fmt.Println(cards, bids)

	mp := make(map[string]int)
	for i, v:= range cards {
		mp[v] = bids[i]
	}

	cardTypes := make([][]string, 7)
	for _, v:= range cards {
		index := getTypeIndexAdvanced(v)
		cardTypes[index] = append(cardTypes[index], v)
	}

	for _,v := range cardTypes {
		sort.SliceStable(v, func(i, j int) bool {
			a := v[i]
			b := v[j]
			var ch1, ch2 int
			for i:=0; i<5; i++ {
				ch1 = getCharIndexAdvanced(a[i])
				ch2 = getCharIndexAdvanced(b[i])
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