package main

import (
	"bufio"
	"log"
	"strings"
)

func part2(scanner *bufio.Scanner) int {
	wordToNum := []string{"_", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var first, last int
		firstFound := false

		for i, char := range line {
			if char >= '0' && char <= '9' {
				if !firstFound {
					first = int(char - '0')
					firstFound = true
				}
				last = int(char - '0')
			} else {
				for j, word := range wordToNum {
					if strings.HasPrefix(line[i:], word) {
						if !firstFound {
							first = j
							firstFound = true
						}
						last = j
					}
				}
			}
		}
		sum += first*10 + last
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}
