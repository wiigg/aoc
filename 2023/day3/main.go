package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const filename = "input.txt"

var rows = []string{}

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	result1 := part1()
	fmt.Println("Part 1 result:", result1)

	result2 := part2()
	fmt.Println("Part 2 result:", result2)
}
