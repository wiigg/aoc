package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	fmt.Println("the sum is", sum)
}
