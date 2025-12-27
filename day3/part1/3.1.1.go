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

		a, err := strconv.Atoi(string(line[0]))
		if err != nil {
			panic("Invalid number")
		}
		b := 0

		for i := 1; i < len(line); i++ {
			x, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic("Invalid number")
			}

			if x > a && i != len(line)-1 {
				a = x
				b = 0
				println("  New max a: ", a)
				continue
			}

			if x > b {
				b = x
				println("  New max b: ", b)
			}
		}

		c := strconv.Itoa(a) + strconv.Itoa(b)
		println("  Pair: ", c)
		converted, err := strconv.Atoi(c)
		if err != nil {
			panic("Invalid pair number")
		}
		total += converted
	}

	fmt.Println("Total sum of pairs: ", total)
}
