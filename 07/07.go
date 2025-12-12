package main

import (
	_ "embed"
	"log"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	part1 := part1(lines)
	log.Printf("part1: %d\n", part1)
	part2 := part2(lines)
	log.Printf("part1: %d\n", part2)
}

func part1(lines []string) int {
	splitCount := 0
	laserPositionsAt := map[int][]int{}

	isHitByLaser := func(y, x int) bool {
		lasersOnPreviousLine := laserPositionsAt[y-1]
		return slices.Contains(lasersOnPreviousLine, x)
	}

	for y, line := range lines {
		for x, char := range line {
			// on finding start, set laser on next line
			if char == 'S' {
				laserPositionsAt[y+1] = append(laserPositionsAt[y+1], x)
			}

			// if splitter hit, add lasers on either side of splitter
			if char == '^' && isHitByLaser(y, x) {
				laserPositionsAt[y] = append(laserPositionsAt[y], []int{x - 1, x + 1}...)
				splitCount += 1
			}

			// if we hit nothing, and laser on above line, continue laser on this line
			if char == '.' && isHitByLaser(y, x) {
				laserPositionsAt[y] = append(laserPositionsAt[y], x)
			}
		}
	}
	return splitCount
}

func part2(lines []string) int {
	height := len(lines)

	// map for storing timelines we already reached
	timelineStore := make(map[[2]int]int)

	var timelinesFrom func(y, x int) int
	timelinesFrom = func(y, x int) int {
		// exit if last line
		if y == height-1 {
			return 1
		}

		// if we already hit this timeline, just return it
		key := [2]int{y, x}
		if v, exists := timelineStore[key]; exists {
			return v
		}

		result := 0

		switch lines[y][x] {
		case 'S', '.':
			// on s or dot, continue
			result = timelinesFrom(y+1, x)
		case '^':
			// on split, timeline from left and right of split
			left := timelinesFrom(y+1, x-1)
			right := timelinesFrom(y+1, x+1)
			result = left + right
		}

		timelineStore[key] = result
		return result
	}

	// find start
	startX := 0
	startY := 0
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				startY = y
				startX = x
				break
			}
		}
	}

	return timelinesFrom(startY, startX)
}
