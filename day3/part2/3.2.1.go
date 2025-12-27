package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func v1() {

	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)

		c := recur(12, line)
		println("  Found: ", c)

		converted, err := strconv.Atoi(c)
		if err != nil {
			panic("Invalid pair number")
		}
		total += converted
	}

	fmt.Println("Total sum of pairs: ", total)
}

func recur(digits int, pool string) string {
	if digits == 0 {
		return ""
	}

	max := 0
	maxIndex := 0
	for i := 0; i <= len(pool)-digits; i++ {
		x, err := strconv.Atoi(string(pool[i]))
		if err != nil {
			panic("Invalid number")
		}
		if x > max {
			max = x
			maxIndex = i
		}
	}

	nextPool := pool[maxIndex+1:]
	return strconv.Itoa(max) + recur(digits-1, nextPool)
}
