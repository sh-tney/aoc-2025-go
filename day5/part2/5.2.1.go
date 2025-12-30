package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	start int
	end   int
}

func v1() {
	file, err := os.Open("../input.txt")
	if err != nil {
		println("Error reading file:", err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	validRanges := map[Pair]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		println("Line: ", line)

		if line == "" {
			break
		}

		ranges := strings.Split(line, "-")
		rangeStart, err := strconv.Atoi(ranges[0])
		rangeEnd, err := strconv.Atoi(ranges[1])
		if err != nil {
			println("Error parsing range:", err.Error())
			return
		}

		validRanges[Pair{start: rangeStart, end: rangeEnd}] = true
	}

	collapsedRanges := map[Pair]bool{}
	movements := 999
	for movements != 0 {
		movements = 0
		for subject := range validRanges {
			moved := false
			for target := range collapsedRanges {
				// Check if subject overlaps or is adjacent to targetRange
				if target != subject && subject.end >= target.start-1 && subject.start <= target.end+1 {
					delete(collapsedRanges, target)
					new := Pair{start: min(subject.start, target.start), end: max(subject.end, target.end)}
					collapsedRanges[new] = true
					movements++
					moved = true
					fmt.Println("Merged:", subject, "with", target, "into", new)
					break
				}
			}
			if !moved {
				collapsedRanges[subject] = true
				fmt.Println("Added new range:", subject)
			}
		}
		clear(validRanges)
		for new := range collapsedRanges {
			validRanges[new] = true
		}
	}

	fmt.Println((collapsedRanges))
	println(len(collapsedRanges), "ranges after collapsing")

	totalLength := 0
	for r := range collapsedRanges {
		totalLength += r.end - r.start + 1
	}
	fmt.Println("Total length of all ranges:", totalLength)
}
