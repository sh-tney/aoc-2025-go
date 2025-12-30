package main

import (
	"bufio"
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

	validRanges := make([]Pair, 0)

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

		validRanges = append(validRanges, Pair{start: rangeStart, end: rangeEnd})
	}

	println("Finished Ranges")

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		println("Checking Node: ", line)

		node, err := strconv.Atoi(line)
		if err != nil {
			println("Error parsing node:", err.Error())
			return
		}

		for _, validRange := range validRanges {
			if node >= validRange.start && node <= validRange.end {
				println("Node ", node, " is valid in range ", validRange.start, "-", validRange.end)
				count++
				break
			}
		}
	}

	println("Total valid nodes: ", count)
}
