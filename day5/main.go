package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := ReadFileAsLines("day5/input")
	seedsStr := strings.Fields(strings.TrimPrefix(lines[0], "seeds: "))
	seeds1 := make([]uint64, len(seedsStr))
	seeds2 := make([][2]uint64, len(seedsStr)/2)
	for i, s := range seedsStr {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			log.Fatalln("Could not parse seed ID", v, err)
		}
		seeds1[i] = v
		if i%2 == 0 {
			seeds2[i/2][0] = v
		} else {
			seeds2[i/2][1] = seeds2[i/2][0] + v - 1
		}
	}

	// 0: seed -> soil
	// 1: soil -> fertilizer
	// 2: fertilizer -> water
	// 3: water -> light
	// 4: light -> temp
	// 5: temp -> hum
	// 6: hum -> loc
	mappings := make([]map[[2]uint64]uint64, 0)
	current := make(map[[2]uint64]uint64)
	for _, line := range lines[2:] {
		if line == "" {
			mappings = append(mappings, current)
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
	mappings = append(mappings, current)

	//minLoc := uint64(math.MaxUint64)
	//for _, seed := range seeds1 {
	//	var val uint64 = seed
	//	for _, m := range mappings {
	//		for src, dst := range m {
	//			if src[0] <= val && val <= src[1] {
	//				val = dst + (val - src[0])
	//				break
	//			}
	//		}
	//		// if we don't find val, it just stays the same
	//		fmt.Println(val)
	//	}
	//
	//	if val < minLoc {
	//		minLoc = val
	//	}
	//}
	//
	//fmt.Println("Part 1:", minLoc)

	/////////////////////////////////
	// PART 2 - with ranges... haven't figured this out yet, need to figure out how to store ranges
	// that aren't explicitly mapped
	/////////////////////////////////
	//for _, sRange := range seeds2 {
	//	start := sRange[0]
	//	end := sRange[1]
	//	// e.g. iterating over the seed -> soil map
	//	for _, currRange := range mappings {
	//		mappedRanges := make([][2]uint64, 0)
	//		keys := sortKeys(currRange)
	//		var dstEnd uint64
	//		// e.g. iterating over each range of mappings [x_1, ..., x_n]
	//		for _, src := range keys {
	//			// if the start of the current "seed" is within the current range
	//			// figure out how much of this range is usable
	//			if src[0] <= start && start <= src[1] {
	//				if end <= src[1] {
	//					dstEnd = (end - start) + src[0]
	//				} else {
	//					// end goes past the end of the current range we're looking in
	//					// so go round the loop again, setting our new start to be 1 passed the
	//					// end of the current range, and setting the end of this dst map to the
	//					// end of the current range.
	//					start = src[1] + 1
	//					dstEnd = src[1]
	//				}
	//				dst := currRange[src]
	//				mappedRanges = append(mappedRanges, [2]uint64{dst, dstEnd})
	//			}
	//		}
	//		// if we don't find val, it just stays the same
	//		fmt.Println(val)
	//
	//	}
	//
	//	if val < minLoc {
	//		minLoc = val
	//	}
	//}

	////////////////
	// PART 2 - brute force
	////////////////

	minLoc := uint64(math.MaxUint64)
	for rangeNo, seed := range seeds2 {
		fmt.Println("calc", rangeNo)
		for i := seed[0]; i <= seed[1]; i++ {
			var val uint64 = i
			for _, m := range mappings {
				for src, dst := range m {
					if src[0] <= val && val <= src[1] {
						val = dst + (val - src[0])
						break
					}
				}
			}
			// if we don't find val, stays the same
			if val < minLoc {
				minLoc = val
			}
		}
	}

	fmt.Println("Part 2:", minLoc)
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

func sortKeys(m map[[2]uint64]uint64) [][2]uint64 {
	keys := make([][2]uint64, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}

	slices.SortFunc(keys, func(a, b [2]uint64) int {
		return cmp.Compare(a[0], b[0])
	})

	return keys
}
