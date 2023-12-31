package main

import (
	"bufio"
	"log"
)

func part1(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var first, last int
		firstFound := false

		for _, char := range line {
			if char >= '0' && char <= '9' {
				if !firstFound {
					first = int(char - '0')
					firstFound = true
				}
				last = int(char - '0')
			}
		}
		sum += first*10 + last
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}
