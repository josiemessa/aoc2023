package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type handType int

const (
	hand_invalid handType = iota
	hand_high
	hand_pair
	hand_twopair
	hand_three
	hand_full
	hand_four
	hand_five
)

var card map[rune]int = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type game struct {
	hand     [5]int
	bid      int
	handType handType
}

func main() {
	lines := ReadFileAsLines("day7/input")

	games := make([]game, len(lines))

	for i, line := range lines {
		f := strings.Fields(line)
		bid, err := strconv.Atoi(f[1])
		if err != nil {
			log.Fatalln("Could not parse bid:", f[1], err)
		}
		games[i].bid = bid
		games[i].hand = [5]int{}
		for j, r := range f[0] {
			games[i].hand[j] = card[r]
		}

		games[i].handType = parseHandType(games[i].hand)
	}
	fmt.Println(games)

	slices.SortFunc(games, func(a, b game) int {
		if a.handType > b.handType {
			return 1
		}
		if a.handType < b.handType {
			return -1
		}
		if a.handType == b.handType {
			for i, r := range a.hand {
				if x := cmp.Compare(r, b.hand[i]); x != 0 {
					return x
				}
			}
		}
		return 0
	})

	//fmt.Println(games)

	var result int
	for i, g := range games {
		result += (i + 1) * g.bid
	}

	fmt.Println("Part 1:", result)
}

func ReadFileAsLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer f.Close()
	fmt.Println(f.Name())

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("scanner error", err)
	}
	return lines
}

func parseHandType(hand [5]int) handType {
	var groups = make(map[int]int)

	for _, r := range hand {
		groups[r]++
	}

	var result handType
	types := make(map[int]int)
	for _, i := range groups {
		types[i]++
	}
	for i, j := range types {
		if i == 1 && j == 5 {
			result = hand_high
			break
		}
		if i == 2 {
			if j == 1 {
				result = hand_pair
			}
			if j == 2 {
				result = hand_twopair
				break
			}
		}
		if i == 3 {
			result = hand_three
			if result == hand_pair {
				result = hand_full
			}
			break
		}
		if i == 4 {
			result = hand_four
			break
		}
		if i == 4 {
			result = hand_five
		}
	}
	return result
}
