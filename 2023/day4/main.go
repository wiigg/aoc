package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const filename = "input.txt"

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result1 := part1(scanner)
	fmt.Println("Part 1 result:", result1)

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	result2 := part2(scanner)
	fmt.Println("Part 2 result:", result2)
}
