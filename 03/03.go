package main

import (
	_ "embed"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	part1 := part01(lines)
	part2 := part2(lines)

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2.String())
}

func part01(lines []string) int {
	sum := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		sum += findJoltage1(line)
	}

	return sum
}

func findJoltage1(line string) int {
	max := 0

	// Brute force all combinations
	for i := 0; i < len(line); i++ {
		for j := i + 1; j < len(line); j++ {
			digit1 := int(line[i] - '0')
			digit2 := int(line[j] - '0')
			joltage, _ := strconv.Atoi(fmt.Sprintf("%d%d", digit1, digit2))

			if joltage > max {
				max = joltage
			}
		}
	}

	return max
}

func part2(lines []string) *big.Int {
	sum := big.NewInt(0)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		sum = sum.Add(sum, findJoltage2(line))
	}

	return sum
}

func findJoltage2(line string) *big.Int {
	length := len(line)
	toRemove := length - 12

	result := make([]byte, 0, 12)

	for i := range length {
		for len(result) > 0 && toRemove > 0 && result[len(result)-1] < line[i] {
			result = result[:len(result)-1]
			toRemove--
		}
		result = append(result, line[i])
	}

	result = result[:len(result)-toRemove]

	res := new(big.Int)
	res, _ = res.SetString(string(result), 10)
	return res
}
