package main

import (
	"regexp"
	"strconv"
)

var reNum = regexp.MustCompile(`\d+`)

func isSymbolAdjacent(row, col int) bool {
	isSymbol := func(r, c int) bool {
		char := rows[r][c]
		return char != '.' && (char < '0' || char > '9')
	}

	isInBounds := func(r, c int) bool {
		return r >= 0 && r < len(rows) &&
			c >= 0 && c < len(rows[r])
	}

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			adRow := row + di
			adCol := col + dj

			if isInBounds(adRow, adCol) && isSymbol(adRow, adCol) {
				return true
			}
		}
	}
	return false
}

func part1() int {
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
