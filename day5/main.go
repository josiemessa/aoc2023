package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := ReadFileAsLines("day5/input")
	seedsStr := strings.Fields(strings.TrimPrefix(lines[0], "seeds: "))
	seeds := make([]uint64, len(seedsStr))
	for i, s := range seedsStr {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			log.Fatalln("Could not parse seed ID", v, err)
		}
		seeds[i] = v
	}

	// 0: seed -> soil
	// 1: soil -> fertilizer
	// 2: fertilizer -> water
	// 3: water -> light
	// 4: light -> temp
	// 5: temp -> hum
	// 6: hum -> loc
	maps := make([]map[[2]uint64]uint64, 0)
	current := make(map[[2]uint64]uint64)
	for _, line := range lines[2:] {
		if line == "" {
			maps = append(maps, current)
			continue
		}
		if strings.HasSuffix(line, ":") {
			// title line
			current = make(map[[2]uint64]uint64)
			continue
		}

		// the destination range start, the source range start, and the range length.
		entry := [3]uint64{}
		vals := strings.Fields(line)
		for i, val := range vals {
			num, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				log.Fatalln("Error parsing line", line, err)
			}
			entry[i] = num
		}
		//             source  , source end              --> destination
		current[[2]uint64{entry[1], (entry[1] + entry[2] - 1)}] = entry[0]
	}
	maps = append(maps, current)

	fmt.Println(maps)

	minLoc := uint64(math.MaxUint64)
	for _, seed := range seeds {
		var val uint64 = seed
		for _, m := range maps {
			for src, dst := range m {
				if src[0] <= val && val <= src[1] {
					val = dst + (val - src[0])
					break
				}
			}
			// if we don't find val, it just stays the same
			fmt.Println(val)
		}

		fmt.Println(val)
		if val < minLoc {
			minLoc = val
		}
		fmt.Println()
	}

	fmt.Println("Part 1:", minLoc)
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
