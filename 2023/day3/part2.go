package main

import (
	"strconv"
)

func isStarAdjacent(row, col int) ([2]int, bool) {
	isStar := func(r, c int) bool {
		return rows[r][c] == '*'
	}

	isInBounds := func(r, c int) bool {
		return r >= 0 && r < len(rows) &&
			c >= 0 && c < len(rows[r])
	}

	var pos = [2]int{}
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			adRow := row + di
			adCol := col + dj

			if isInBounds(adRow, adCol) && isStar(adRow, adCol) {
				pos = [2]int{adRow, adCol}
				return pos, true
			}
		}
	}
	return pos, false
}

func part2() int {
	var gears = map[[2]int][]int{}
	for i, row := range rows {
		numMatch := reNum.FindAllStringIndex(row, -1)
		for _, match := range numMatch {
			numStr := row[match[0]:match[1]]
			numValue, _ := strconv.Atoi(string(numStr))

			pos, found := isStarAdjacent(i, match[0])
			firstPos := pos
			if found {
				gears[pos] = append(gears[pos], numValue)
			}

			pos, found = isStarAdjacent(i, match[1]-1)
			if found && pos != firstPos {
				gears[pos] = append(gears[pos], numValue)
			}
		}
	}

	ratios := 0
	for _, nums := range gears {
		if len(nums) == 2 {
			ratios += nums[0] * nums[1]
		}
	}

	return ratios
}
