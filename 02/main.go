package main

import (
	_ "embed"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type idRange struct {
	start int
	end   int
}

func main() {
	ranges := parseInput(input)
	log.Printf("Part 1: %s", part01(ranges))
	log.Printf("Part 2: %s", part02(ranges))
}

func part01(ranges []idRange) string {
	var sum big.Int
	sum.SetInt64(0)
	// iterate the ranges
	for _, r := range ranges {
		// iterate the range between start and end
		for i := r.start; i <= r.end; i++ {
			if isInvalidID(i) {
				sum.Add(&sum, big.NewInt(int64(i)))
			}
		}
	}
	return sum.String()
}

func part02(ranges []idRange) string {
	var sum big.Int
	sum.SetInt64(0)
	// iterate the ranges
	for _, r := range ranges {
		// iterate the range between start and end
		for i := r.start; i <= r.end; i++ {
			if isInvalidIDpart2(i) {
				sum.Add(&sum, big.NewInt(int64(i)))
			}
		}
	}
	return sum.String()
}

func parseInput(s string) []idRange {
	var ranges []idRange
	ids := strings.SplitSeq(strings.TrimSpace(s), ",")
	for id := range ids {
		ranges = append(ranges, parseRange(id))
	}
	return ranges
}

func parseRange(s string) idRange {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return idRange{
		start: start,
		end:   end,
	}
}

func isInvalidID(id int) bool {
	// convert to string
	asString := fmt.Sprintf("%d", id)

	// if not divisible by 2, then the number is invalid as it is not splitable
	if len(asString)%2 != 0 {
		return false
	}

	// split in two
	split := len(asString) / 2
	part1 := asString[:split]
	part2 := asString[split:]

	return part1 == part2
}

func isInvalidIDpart2(id int) bool {
	// conver to string
	asString := fmt.Sprintf("%d", id)
	length := len(asString)

	for blockSize := 1; blockSize <= length/2; blockSize++ {
		if length%blockSize != 0 {
			continue
		}

		parts := length / blockSize
		if parts < 2 {
			continue
		}

		part := asString[:blockSize]
		isInvalidID := true

		for i := 1; i < parts; i++ {
			if asString[i*blockSize:(i+1)*blockSize] != part {
				isInvalidID = false
				break
			}
		}

		if isInvalidID {
			return true
		}
	}
	return false
}
