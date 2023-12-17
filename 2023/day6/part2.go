package main

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func part2(scanner *bufio.Scanner) int {
	var time, distance int
	readTimes := true
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())[1:]
		value := strings.Join(values, "")
		if readTimes {
			time, _ = strconv.Atoi(value)
		} else {
			distance, _ = strconv.Atoi(value)
		}
		readTimes = false
	}

	x1 := float64(time) + math.Sqrt(float64(time*time)-4*float64(distance))/2
	x2 := float64(time) - math.Sqrt(float64(time*time)-4*float64(distance))/2

	return int(x1 - x2)
}
