package main

func transformRange(initialRange Range) []Range {
	currentRanges := []Range{initialRange}
	for _, m := range mappings {
		newRanges := []Range{}
		for _, currentRange := range currentRanges {
			newRanges = append(newRanges, mapRange(currentRange, m)...)
		}
		currentRanges = newRanges
	}
	return currentRanges
}

func mapRange(inputRange Range, mappings []Mapping) []Range {
	outputRanges := []Range{}
	unmappedMin := inputRange.min

	for unmappedMin <= inputRange.max {
		mapped := false
		for _, mapping := range mappings {
			if unmappedMin > mapping.src.max {
				continue
			}
			if unmappedMin >= mapping.src.min {
				mapped = true
				mappedMax := min(inputRange.max, mapping.src.max)
				outputMin := mapping.dest + (unmappedMin - mapping.src.min)
				outputMax := mapping.dest + (mappedMax - mapping.src.min)
				outputRanges = append(outputRanges, Range{outputMin, outputMax})
				unmappedMin = mappedMax + 1
				break
			}
		}
		if !mapped {
			outputRanges = append(outputRanges, Range{unmappedMin, inputRange.max})
			break
		}
	}
	return outputRanges
}

func part2() int {
	minLocation := 1 << 31
	for i := 0; i < len(seeds); i += 2 {
		seedRange := Range{seeds[i], seeds[i] + seeds[i+1] - 1}
		locationRange := transformRange(seedRange)

		// find the smallest location for this seed
		location := locationRange[0].min
		for _, r := range locationRange {
			location = min(location, r.min)
		}
		minLocation = min(minLocation, location)
	}

	return minLocation
}
