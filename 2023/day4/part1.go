package main

import (
	"bufio"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ": ")
		numbers := strings.Split(card[1], " | ")

		winning := map[int]bool{}
		for _, numStr := range strings.Fields(numbers[0]) {
			if num, err := strconv.Atoi(numStr); err == nil {
				winning[num] = true
			}
		}

		score := 0
		for _, numStr := range strings.Fields(numbers[1]) {
			if num, err := strconv.Atoi(numStr); err == nil && winning[num] {
				if score == 0 {
					score = 1
				} else {
					score <<= 1
				}
			}
		}
		sum += score
	}
	return sum
}
