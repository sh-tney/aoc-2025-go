package main

import (
	"fmt"
	"os"
)

func v2() {

	file, err := os.Open("input.txt")
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

		for i := 0; i < magnitude; i++ {
			if direction == 'L' {
				position--
			} else {
				position++
			}
			if position == 100 {
				position = 0
			}
			if position == -1 {
				position = 99
			}
			if position == 0 {
				count++
			}
		}
	}

	fmt.Println("Count of times position was 0:", count)
}
