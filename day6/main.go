package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := ReadFileAsLines("day6/input")
	timeStr := strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time: "), " ", "")
	distStr := strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance: "), " ", "")

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Fatalln("Could not parse time", timeStr)
	}
	dist, err2 := strconv.Atoi(distStr)
	if err2 != nil {
		log.Fatalln("could not parse dist", distStr)
	}
	// distance = (time - i)*i = time*i - i^2
	// find the min i

	var wins int
	for i := 1; i < time; i++ {
		distance := (time - i) * i
		if distance > dist {
			wins++
		}
	}

	fmt.Println("Part 2", wins)

}

func part1(lines []string) {
	times := strings.Fields(strings.TrimPrefix(lines[0], "Time: "))
	dists := strings.Fields(strings.TrimPrefix(lines[1], "Distance: "))
	races := make([][2]int, len(times))
	for i, t := range times {
		n, err := strconv.Atoi(t)
		if err != nil {
			log.Fatalln("Could not parse time", t)
		}
		races[i] = [2]int{n, 0}
	}

	for i, d := range dists {
		n, err := strconv.Atoi(d)
		if err != nil {
			log.Fatalln("Could not parse distance", d)
		}
		races[i][1] = n
	}

	fmt.Println(races)

	var result = 1
	for _, race := range races {
		var wins int
		for i := 1; i < race[0]; i++ {
			distance := (race[0] - i) * i
			fmt.Println(i, distance)
			if distance > race[1] {
				wins++
			}
		}
		if wins != 0 {
			result *= wins
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
