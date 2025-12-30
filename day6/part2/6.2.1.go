package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func v1() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic("file error")
	}
	defer file.Close()

	rows := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, line)
	}
	fmt.Println(rows)

	signs := rows[len(rows)-1]
	fmt.Println("Signs: ", signs)
	rows = rows[:len(rows)-1]
	fmt.Println("Numbers:", rows)

	total := 0
	op := sum
	nums := make([]int, 0)
	for pos := len(signs) - 1; pos >= 0; pos-- {
		// read column
		str := ""
		for _, r := range rows {
			str += string(r[pos])
		}

<<<<<<< HEAD
		// if column is empty, process current nums and reset
=======
		// if column is empty, skip
>>>>>>> 0809ee9 (day 6)
		if strings.TrimSpace(str) == "" {
			continue
		} else {
			println("Num:", str)
			num, err := strconv.Atoi(strings.TrimSpace(str))
			if err != nil {
				panic("parse error")
			}
			nums = append(nums, num)
		}

		// process on operator change
		if signs[pos] != ' ' {
			if signs[pos] == '+' {
				op = sum
			} else {
				op = prod
			}
			fmt.Println("Finished column", nums, "total:", op(nums))
			total += op(nums)
			nums = make([]int, 0)
		}
	}

	fmt.Println("Grand Total: ", total)
}

func sum(inp []int) int {
	total := 0
	for _, val := range inp {
		total += val
	}
	return total
}

func prod(inp []int) int {
	total := 1
	for _, val := range inp {
		total *= val
	}
	return total
}
