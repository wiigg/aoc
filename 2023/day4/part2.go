package main

import (
	"bufio"
	"strconv"
	"strings"
)

func part2(scanner *bufio.Scanner) int {
	total := 0
	copies := make([]int, 196)
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ": ")
		numbers := strings.Split(card[1], " | ")

		winning := map[int]bool{}
		for _, numStr := range strings.Fields(numbers[0]) {
			if num, err := strconv.Atoi(numStr); err == nil {
				winning[num] = true
			}
		}

		count := 0
		for _, numStr := range strings.Fields(numbers[1]) {
			if num, err := strconv.Atoi(numStr); err == nil && winning[num] {
				copies[count+1] += copies[0] + 1
				count++
			}
		}
		total += copies[0] + 1
		copies = copies[1:]
	}
	return total
}
