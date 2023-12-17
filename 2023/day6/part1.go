package main

import (
	"bufio"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	var times, distances []int
	readTimes := true
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())[1:]
		for _, v := range values {
			num, _ := strconv.Atoi(v)
			if readTimes {
				times = append(times, num)
			} else {
				distances = append(distances, num)
			}
		}
		readTimes = false
	}

	ways := 1
	for i, T := range times {
		count := 0
		for x := 0; x < T; x++ {
			d := float64(x) * float64(T-x)
			if d > float64(distances[i]) {
				count++
			}
		}
		ways *= count
	}
	return ways
}
