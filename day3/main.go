package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := ReadFileAsLines("day3/test_input")

	part2(lines)

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

func part2(lines []string) {
	type pair struct{ r, c int }
	gears := make(map[pair][]int)

	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			var (
				numberStr string
				start     int
				end       int
			)

			if c := line[col]; c > 47 && c < 58 {
				numberStr = string(c)
				start = col
				end = col

				// lookahead for the rest of the number
				for i := 1; col+i < len(line); i++ {
					if c := line[col+i]; c > 47 && c < 58 {
						numberStr += string(c)
						end = col + i
					} else {
						break
					}
				}

				if numberStr != "" && end > 0 {
					col = end
				}

				// look around for symbols
				var found bool
				var p pair
				for r := -1; r <= 1; r++ {
					if row+r < 0 || row+r > len(line)-1 {
						continue
					}
					for i := start - 1; i <= end+1; i++ {
						if i < 0 || i > len(lines)-1 {
							continue
						}

						if c := lines[row+r][i]; c == '*' {
							// we have found a gear
							found = true
							p = pair{r: row + r, c: i}
						}
					}
					if found {
						break
					}
				}
				if found {
					n, err := strconv.Atoi(numberStr)
					if err != nil {
						log.Fatalf("Could not parse number %q, with error: %q, from line %q\n", numberStr, err, line)
					}
					gears[p] = append(gears[p], n)
				}
			}
		}
	}

	var result int

	for _, partNos := range gears {
		if len(partNos) == 2 {
			result += partNos[0] * partNos[1]
		}
	}
	fmt.Println("Part 2:", result)

}

func part1(lines []string) {
	parts := make([]int, 0)

	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			var (
				numberStr string
				start     int
				end       int
			)

			if c := line[col]; c > 47 && c < 58 {
				numberStr = string(c)
				start = col
				end = col

				// lookahead for the rest of the number
				for i := 1; col+i < len(line); i++ {
					if c := line[col+i]; c > 47 && c < 58 {
						numberStr += string(c)
						end = col + i
					} else {
						break
					}
				}

				if numberStr != "" && end > 0 {
					col = end
				}

				// look around for symbols
				var found bool
				for r := -1; r <= 1; r++ {
					if row+r < 0 || row+r > len(line)-1 {
						continue
					}
					for i := start - 1; i <= end+1; i++ {
						if i < 0 || i > len(lines)-1 {
							continue
						}

						// symbols are in ascii range 33-47
						re := regexp.MustCompile(`[^0-9]`)
						if c := lines[row+r][i]; re.Match([]byte{c}) && c != '.' {
							// we have found a part number
							found = true
							break
						}
					}
					if found {
						break
					}
				}
				if found {
					n, err := strconv.Atoi(numberStr)
					if err != nil {
						log.Fatalf("Could not parse number %q, with error: %q, from line %q\n", numberStr, err, line)
					}
					parts = append(parts, n)
				}
			}
		}
	}
	sort.Ints(parts)
	fmt.Println(parts)

	var result int
	for _, v := range parts {
		result += v
	}

	fmt.Println("Part 1:", result)
}
