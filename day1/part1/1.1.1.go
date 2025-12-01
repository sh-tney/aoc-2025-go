package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var position int = 50
	var count = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0:1]
		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			panic("Invalid magnitude")
		}

		pre := position

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

		fmt.Println(direction, magnitude, " | ", pre, "->", position)
	}

	fmt.Println("Count of times position was 0:", count)
}

func moveLeft(position int, magnitude int) int {
	m := magnitude % 100
	p := position - m

	if p < 0 {
		return 100 + p
	}

	return p
}

func moveRight(position int, magnitude int) int {
	m := magnitude % 100
	p := position + m

	if p > 99 {
		return p - 100
	}

	return p
}
