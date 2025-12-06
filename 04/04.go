package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string
var grid [][]rune

func main() {
	for line := range strings.SplitSeq(input, "\n") {
		grid = append(grid, []rune(line))
	}

	part1 := part1()
	log.Printf("Result: %d\n", part1)

	part2 := part2()
	log.Printf("Result: %d\n", part2)
}

func part1() int {
	count := 0
	for y, row := range grid {
		for x := range row {
			if grid[y][x] != '@' {
				continue
			}
			if checkAccessible(x, y) {
				count++
			}
		}
	}
	return count
}

func part2() int {
	count := 0

	for {
		removed := 0
		for y, row := range grid {
			for x := range row {
				if grid[y][x] != '@' {
					continue
				}
				if checkAccessible(x, y) {
					grid[y][x] = 'x'
					removed++
				}
			}
		}
		// when none are removed in the iteration, just return
		if removed == 0 {
			return count
		}
		count += removed
	}
}

func checkAccessible(x, y int) bool {
	count := 0

	check := func(x, y int) {
		if y >= 0 && y < len(grid) {
			row := grid[y]
			if x >= 0 && x < len(row) {
				if row[x] == '@' {
					count++
				}
			}
		}
	}

	// check up-left
	check(x-1, y-1)
	// check up
	check(x, y-1)
	// check up-right
	check(x+1, y-1)
	// check left
	check(x-1, y)
	// check right
	check(x+1, y)
	// check down-left
	check(x-1, y+1)
	// check down
	check(x, y+1)
	// check down-right
	check(x+1, y+1)

	return count < 4
}
