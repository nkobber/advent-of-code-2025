package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Direction string

const (
	DirectionLeft  = "L"
	DirectionRight = "R"
)

func main() {
	// Split by line
	lines := strings.Split(input, "\n")

	part1 := part1(lines)
	fmt.Printf("Part 1 count: %d\n", part1)

	part2 := part2(lines)
	fmt.Printf("Part 2 count: %d\n", part2)
}

func part1(lines []string) int {
	currentPos := 50
	count := 0

	// Iterate lines and parse
	for _, line := range lines {
		if line == "" {
			continue
		}

		direction, clicks := parseLine(line)

		if direction == DirectionRight {
			currentPos = currentPos + clicks
		} else {
			currentPos = currentPos - clicks
		}

		// If divisible by 100, that means we hit 0
		if currentPos%100 == 0 {
			count++
		}

	}

	return count
}

func part2(lines []string) int {
	currentPos := 50
	count := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction, clicks := parseLine(line)
		for range clicks {
			if direction == DirectionRight {
				currentPos = currentPos + 1
			} else {
				currentPos = currentPos - 1
			}
			// If currentPos is divisible by 100, that means we hit 0
			if currentPos%100 == 0 {
				count = count + 1
			}
		}
	}

	return count
}

func parseLine(line string) (Direction, int) {
	direction := string(line[0])
	clicks, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}
	return Direction(direction), clicks
}
