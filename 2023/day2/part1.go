package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	limits := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], ";")

		if processRounds(rounds, limits) {
			gameId := strings.Split(game[0], " ")[1]
			total, _ := strconv.Atoi(gameId)
			sum += total
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func processRounds(rounds []string, limits map[string]int) bool {
	for _, round := range rounds {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			numType := strings.Split(cube, " ")
			num, _ := strconv.Atoi(strings.TrimSpace(numType[0]))
			if num > limits[numType[1]] {
				return false
			}
		}
	}
	return true
}
