package main

import (
	"bufio"
	"fmt"
	"os"
)

func v1() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic("file error")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// init
	scanner.Scan()
	line := scanner.Text()
	beams := make([]bool, len(line))
	for i, x := range line {
		beams[i] = x == 'S'
	}

	splitCount := 0
	for scanner.Scan() {
		line = scanner.Text()
		for i, x := range line {
			if x == '^' && beams[i] {
				splitCount++
				beams[i-1] = true
				beams[i] = false
				beams[i+1] = true
			}
		}
	}

	fmt.Println("Final Beams:")
	fmt.Println(beams)
	fmt.Println("Total splits:", splitCount)
}
