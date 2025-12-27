package main

import (
	"bufio"
	"fmt"
	"os"
)

func v1() {
	file, err := os.Open("../test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := [][]string{}

	xMax, yMax := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line: ", line)

		yMax = 0
		row := make([]string, len(line))
		for _, char := range line {
			row[yMax] = string(char)
			yMax++
		}
		field = append(field, row)
		xMax++
	}

	outOfBounds := func(x, y int) bool {
		if x < 0 || y < 0 || x >= xMax || y >= yMax {
			//println("OOB Check: ", x, y)
			return true
		}
		return false
	}

	adjacentRolls := func(x, y int) int {
		count := 0
		//println("checking adj FROM", x, " ", y)
		for _, dx := range []int{-1, 0, 1} {
			for _, dy := range []int{-1, 0, 1} {
				//println("checking adj ", x+dx, " ", y+dy)
				if outOfBounds(x+dx, y+dy) || (dx == 0 && dy == 0) {
					count += 0
				} else if field[x+dx][y+dy] == "@" || field[x+dx][y+dy] == "x" {
					//println("found adj @ at ", x+dx, " ", y+dy)
					count++
				}
			}
		}
		//println("found ", count)
		return count
	}

	println("NEW FIELD:")

	total := 0
	for x, row := range field {
		for y := range row {
			if field[x][y] == "@" && adjacentRolls(x, y) < 4 {
				field[x][y] = "x"
				total++
			}
		}
		fmt.Println(field[x])
	}

	println("Total: ", total)
}
