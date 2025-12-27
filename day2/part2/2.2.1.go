package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func v1() {

	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	total := 0
	ranges := strings.Split(line, ",")
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		low, err1 := strconv.Atoi(bounds[0])
		high, err2 := strconv.Atoi(bounds[1])
		if err1 != nil || err2 != nil {
			panic("Invalid bounds")
		}
		fmt.Println("Range:", low, "to", high)

		for i := low; i <= high; i++ {
			fmt.Println("  Value:", i)
			if checkComplexMatch(i) {
				total += i
			}
		}
	}

	fmt.Println("Total sum of matching numbers:", total)
}

func checkComplexMatch(i int) bool {
	converted := strconv.Itoa(i)

	for i := 1; i < len(converted); i++ {
		if len(converted)%i != 0 {
			println("		Skipping length", i)
		} else {
			segments := map[string]string{}
			for d := 0; d < len(converted)/i; d++ {
				segment := converted[d*i : (d+1)*i]
				segments[segment] = segment
			}
			if len(segments) == 1 {
				println("		Found ", converted, " in segments of ", i, ": ", segments)
				return true
			}
		}
	}

	return false
}
