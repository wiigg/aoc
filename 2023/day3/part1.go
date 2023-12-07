package main

import (
	"bufio"
	"regexp"
	"strconv"
)

var reNum = regexp.MustCompile(`\d+`)

var rows = []string{}

func isSymbolAdjacent(row, col int) bool {
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			adRow := row + di
			adCol := col + dj

			if adRow >= 0 && adRow < len(rows) &&
				adCol >= 0 && adCol < len(rows[adRow]) &&
				rows[adRow][adCol] != '.' &&
				(rows[adRow][adCol] < '0' || rows[adRow][adCol] > '9') {
				return true
			}
		}
	}
	return false
}

func part1(scanner *bufio.Scanner) int {
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	sum := 0
	for i, row := range rows {
		numMatch := reNum.FindAllStringIndex(row, -1)
		for _, match := range numMatch {
			numStr := row[match[0]:match[1]]
			numValue, _ := strconv.Atoi(string(numStr))
			if isSymbolAdjacent(i, match[0]) || isSymbolAdjacent(i, match[1]-1) {
				sum += numValue
			}
		}
	}

	return sum
}
