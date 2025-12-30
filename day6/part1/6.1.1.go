package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func v1() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic("file error")
	}
	defer file.Close()

	stuff := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`[\s]+`)
		stuff = append(stuff, re.Split(strings.TrimSpace(line), -1))
	}
	fmt.Println(stuff)

	grandTotal := 0
	for j := range stuff[0] {
		total := 0
		if stuff[len(stuff)-1][j] == "+" {
			println("adding")
			total = 0
			for i := 0; i < len(stuff)-1; i++ {
				num, _ := strconv.Atoi(stuff[i][j])
				println(num, total)
				total = total + num
			}
		} else if stuff[len(stuff)-1][j] == "*" {
			println("multiplying", len(stuff)-1)
			total = 1
			for i := 0; i < len(stuff)-1; i++ {
				num, _ := strconv.Atoi(stuff[i][j])
				println(num, total)
				total = total * num
			}
		}
		fmt.Println("Column ", j, " total: ", total)
		grandTotal = grandTotal + total
	}
	fmt.Println("Grand Total: ", grandTotal)
}
