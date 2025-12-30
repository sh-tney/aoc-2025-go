package main

import (
	"bufio"
	"fmt"
	"os"
)

func v1() {
	file, err := os.Open("../test.txt")
	if err != nil {
		panic("file error")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// init
	scanner.Scan()
	line := scanner.Text()
	beams := make([]int, len(line))
	for i, x := range line {
		if x == 'S' {
			beams[i] = 1
		} else {
			beams[i] = 0
		}
	}

	fmt.Println(beams)
	for scanner.Scan() {
		line = scanner.Text()
		for i, x := range line {
			if x == '^' && beams[i] > 0 {
				beams[i-1] += beams[i]
				beams[i+1] += beams[i]
				beams[i] = 0
			}
		}
		fmt.Println(beams)
	}

	total := 0
	for _, n := range beams {
		total += n
	}
	fmt.Println("Timelines:", total)
}
