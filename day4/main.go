package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	lines := ReadFileAsLines("day4/test_input")
	part2(lines)

}

func part2(lines []string) {
	cards := make([]int, len(lines))

	// we start with 1 card
	for i := range cards {
		cards[i] = 1
	}

	// note game number is i+1
	for i, line := range lines {
		game := strings.Split(line, ":")
		split := strings.Split(game[1], " | ")
		winning := strings.Fields(split[0])
		myNum := strings.Fields(split[1])

		var found int
		for _, v := range winning {
			for _, m := range myNum {
				if v == m {
					found++
				}
			}
		}
		for k := 0; k < cards[i]; k++ {
			for j := 1; j <= found; j++ {
				cards[i+j] += 1
			}
		}
	}

	var result int

	for _, c := range cards {
		result += c
	}

	fmt.Println("Part 2:", result)
}

func part1(lines []string) {
	var result float64
	for _, line := range lines {
		game := strings.Split(line, ":")
		split := strings.Split(game[1], " | ")
		winning := strings.Fields(split[0])
		myNum := strings.Fields(split[1])

		var found int
		for _, v := range winning {
			for _, m := range myNum {
				if v == m {
					found++
				}
			}
		}

		if found > 0 {
			result += math.Pow(2, float64(found)-1)
		}
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
