package main

import (
	"aoc2021/shared"
	"log"
	"math"
)

func Part1(depths []int) int {
	depthIncreases := 0
	lastDepth := math.MaxInt

	for _, depth := range depths {
		if depth > lastDepth {
			depthIncreases++
		}
		lastDepth = depth
	}

	return depthIncreases
}

func Part2(depths []int) int {
	var windowSums []int
	depthIncreases := 0

	for i, depth := range depths {
		if i >= 2 {
			windowSum := depth + depths[i-1] + depths[i-2]

			countPreviousSums := len(windowSums)
			if countPreviousSums > 0 {
				if windowSum > windowSums[countPreviousSums-1] {
					depthIncreases++
				}
			}

			windowSums = append(windowSums, windowSum)
		}
	}
	return depthIncreases
}

func main() {
	numbers := shared.ReadNumberFile("day01/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
