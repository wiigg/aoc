package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func part2(scanner *bufio.Scanner) int {
	powers := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], ";")

		powers += revisedProcessRounds(rounds)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return powers
}

func revisedProcessRounds(rounds []string) int {
	fewest := map[string]int{}
	for _, round := range rounds {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			numType := strings.Split(cube, " ")
			num, _ := strconv.Atoi(strings.TrimSpace(numType[0]))
			fewest[numType[1]] = max(num, fewest[numType[1]])
		}
	}
	power := fewest["blue"] * fewest["green"] * fewest["red"]
	return power
}
