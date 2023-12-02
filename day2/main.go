package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type draw struct {
	red   int
	green int
	blue  int
}

const (
	red        = "red"
	green      = "green"
	blue       = "blue"
	redCubes   = 12
	greenCubes = 13
	blueCubes  = 14
)

// Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
func main() {
	lines := ReadFileAsLines("day2/input")
	games := make([][]draw, len(lines))
	// note that index in array is game number - 1
	for i, line := range lines {
		// Split for "Game: <draws>"
		s := strings.Split(line, ": ")

		// Split for "<draw1>; <draw2>""
		results := strings.Split(s[1], "; ")
		draws := make([]draw, len(results))
		for j, res := range results {
			d := draw{}

			// Split for "\d col1, \d col2"
			col := strings.Split(res, ", ")
			for _, c := range col {
				// Split for "\d col1"
				val, err := strconv.Atoi(strings.Split(c, " ")[0])
				if err != nil {
					log.Fatalln("error on line", line)
				}

				switch {
				case strings.HasSuffix(c, red):
					d.red = val
				case strings.HasSuffix(c, blue):
					d.blue = val
				case strings.HasSuffix(c, green):
					d.green = val

				}
			}
			draws[j] = d
		}
		games[i] = draws
	}

	// fmt.Println(games)

	// PART 1
	part1 := 0
	for i, game := range games {
		valid := true
		for _, d := range game {
			if !valid {
				break
			}
			valid = valid && d.red <= redCubes && d.blue <= blueCubes && d.green <= greenCubes
		}

		if valid {
			part1 += i + 1
		}
	}

	fmt.Println("Part 1:", part1)

	// PART 2
	maxs := 0
	for _, game := range games {
		max := draw{}

		for _, d := range game {
			if d.red > max.red {
				max.red = d.red
			}
			if d.blue > max.blue {
				max.blue = d.blue
			}
			if d.green > max.green {
				max.green = d.green
			}
		}
		maxs += max.red * max.blue * max.green
	}

	fmt.Println("Part 2:", maxs)
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

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
