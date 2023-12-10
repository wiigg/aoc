package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Range struct {
	min, max int
}

type Mapping struct {
	src  Range
	dest int
}

var seeds []int
var mappings [][]Mapping

func mapSeed(seed int) int {
	for _, mapSet := range mappings {
		for _, m := range mapSet {
			if m.src.min <= seed && seed <= m.src.max {
				seed = m.dest + (seed - m.src.min)
				break
			}
		}
	}
	return seed
}

func part1(scanner *bufio.Scanner) int {
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "seeds:"):
			parts := strings.Fields(line)
			for _, part := range parts[1:] {
				seed, _ := strconv.Atoi(part)
				seeds = append(seeds, seed)
			}
		case strings.Contains(line, "map:"):
			var mapSet []Mapping
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					break
				}

				parts := strings.Fields(line)
				dest, _ := strconv.Atoi(parts[0])
				src, _ := strconv.Atoi(parts[1])
				rge, _ := strconv.Atoi(parts[2])
				mapSet = append(mapSet, Mapping{Range{src, src + rge - 1}, dest})
			}
			mappings = append(mappings, mapSet)
		}
	}

	minLocation := 1 << 31
	for _, seed := range seeds {
		minLocation = min(minLocation, mapSeed(seed))
	}

	return minLocation
}
