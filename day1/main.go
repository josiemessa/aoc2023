package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//var numbersRegex = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
var numbersMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	lines := ReadFileAsLines("day1/input")

	// PART 1
	digits := make([]int, len(lines))
	for i, line := range lines {
		// part 1
		var digit string
		for j, c := range line {
			if c > 47 && c < 58 {
				// c is a number
				digit += string(c)
				continue
			}
			for s, val := range numbersMap {
				if j+len(s) > len(line) {
					continue
				}
				if line[j:j+len(s)] == s {
					digit += val
				}
			}
		}

		// PART 2
		//digit := numbersRegex.FindAll([]byte(line), -1)
		//for j, d := range digit {
		//	if v, ok := numbersMap[string(d)]; ok {
		//		digit[j] = []byte(v)
		//	}
		//}
		//fmt.Println(digit)
		//
		d, err := strconv.Atoi(string(digit[0]) + string(digit[len(digit)-1]))
		if err != nil {
			log.Fatalln("failed to find digits and parse them in line", line, err)
		}
		digits[i] = d
	}

	var result int

	fmt.Println(digits)
	for _, digit := range digits {
		result += digit
	}

	fmt.Println("Success:", result)
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
