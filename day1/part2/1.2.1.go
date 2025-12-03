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

	var position int = 50
	var count = 0

	moveLeft := func(position int, magnitude int) int {
		count += magnitude / 100
		m := magnitude % 100
		p := position - m

		if p < 0 {
			if position != 0 {
				count++
			}
			return 100 + p
		}

		return p
	}

	moveRight := func(position int, magnitude int) int {
		count += magnitude / 100
		m := magnitude % 100
		p := position + m

		if p > 99 {
			if position != 0 && p-100 != 0 {
				count++
			}
			return p - 100
		}

		return p
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0:1]
		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			panic("Invalid magnitude")
		}

		switch direction {
		case "L":
			position = moveLeft(position, magnitude)
		case "R":
			position = moveRight(position, magnitude)
		default:
			panic("Invalid direction")
		}

		if position == 0 {
			count++
		}

	}

	fmt.Println("Count of times position was 0:", count)

}
