package main

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type idRange struct {
	start int64
	end   int64
}

var (
	idRanges = []idRange{}
	ids      = []int64{}
)

func main() {
	isPastRanges := false
	for line := range strings.SplitSeq(input, "\n") {
		// Change how we parse when we switch from ranges to available
		if line == "" {
			isPastRanges = true
			continue
		}

		if !isPastRanges {
			parts := strings.Split(line, "-")
			start, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}
			end, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				panic(err)
			}
			idRange := idRange{
				start: start,
				end:   end,
			}
			idRanges = append(idRanges, idRange)
		} else {
			id, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}
			ids = append(ids, id)
		}

	}

	part1 := part1()
	log.Printf("Part 1: %d\n", part1)
	part2 := part2()
	log.Printf("Part 2: %d\n", part2)
}

func part1() int {
	count := 0
	for _, id := range ids {
		for _, idRange := range idRanges {
			if id >= idRange.start && id <= idRange.end {
				count++
				break
			}
		}
	}
	return count
}

func part2() int64 {
	var sum int64 = 0
	for _, r := range mergeRanges(idRanges) {
		sum += r.end - r.start + 1
	}
	return sum
}

func mergeRanges(oldRanges []idRange) []idRange {
	// copy to avoid sorting the original array
	ranges := append([]idRange{}, oldRanges...)

	// sort by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := make([]idRange, 0, len(ranges))
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		theRange := ranges[i]

		// check overlap
		if theRange.start <= current.end {
			if theRange.end > current.end {
				current.end = theRange.end
			}
			continue
		}

		// if no overlap, add it
		merged = append(merged, current)
		current = theRange
	}

	merged = append(merged, current)

	return merged
}
