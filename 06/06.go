package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1 := part1()
	log.Printf("Part1: %d\n", part1)
	part2 := part2()
	log.Printf("Part2: %d\n", part2)
}

func part1() int {
	split := strings.Split(input, "\n")
	re := regexp.MustCompile(`[+*]`)
	operators := re.FindAllString(split[len(split)-1], -1)

	var numbers [][]string
	re = regexp.MustCompile(`\d+`)
	for i := 0; i < len(split)-1; i++ {
		numbers = append(numbers, re.FindAllString(split[i], -1))
	}

	totalSum := 0

	for x := range numbers[0] {
		sum := 0
		operator := operators[x]
		if operator == "*" {
			sum = 1
		}
		for y := range len(numbers) {
			parsed, err := strconv.Atoi(numbers[y][x])
			if err != nil {
				panic(err)
			}

			switch operator {
			case "*":
				sum *= parsed
			case "+":
				sum += parsed
			default:
				panic("unknown operator")
			}

		}

		totalSum += sum

	}

	return totalSum
}

func part2() int64 {
	split := strings.Split(input, "\n")
	numbers := split[0 : len(split)-1]
	operators := split[len(split)-1]

	var totalSum int64 = 0

	// iterate each column at a time, use operator to check if we are done with the "column block"
	numbersInCurrentBlock := []int{}
	for x := len(numbers[0]) - 1; x >= 0; x-- {
		numAsString := ""
		for y := range numbers {
			if string(numbers[y][x]) != " " {
				numAsString += string(numbers[y][x])
			}
		}

		// if we hit col with no numbers
		if numAsString == "" {
			continue
		}

		asInt, err := strconv.Atoi(numAsString)
		if err != nil {
			panic(err)
		}

		numbersInCurrentBlock = append(numbersInCurrentBlock, asInt)

		// find the operator
		operator := string(operators[x])
		if operator == " " {
			continue
		}

		// if we found an operator, we calculate the block, and empty it
		switch operator {
		case "*":
			totalSum += int64(multiply(numbersInCurrentBlock))
		case "+":
			totalSum += int64(add(numbersInCurrentBlock))
		}
		numbersInCurrentBlock = []int{}
	}

	return totalSum
}

func multiply(nums []int) int {
	res := 1
	for _, num := range nums {
		res *= num
	}
	return res
}

func add(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
