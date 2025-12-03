package main

import (
	"fmt"
	"os"
)

func v3() {

	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var position int = 50
	var count = 0

	var direction rune
	var magnitude int

	for {
		_, err := fmt.Fscanf(file, `%c%d`, &direction, &magnitude)
		if err != nil {
			break
		}

		// Accounts for full loops around the circle i.e. L485 = 4 extra loops
		count += magnitude / 100
		m := magnitude % 100

		if direction == 'L' {
			position -= m
		} else {
			position += m
		}

		// Do counting/wrapping logic only when necessary
		if position > 99 || position < 0 || position == 0 {
			// Doesn't count if we started from zero and went left
			if position != -m {
				count++
			}

			// Wraps from 118 to 18
			position = position % 100

			// Wraps around from -5 to 95
			if position < 0 {
				position = 100 + (position % 100)
			}
		}
	}

	fmt.Println("Count of times position was 0:", count)
}
