package main

import (
	"aoc2021/shared"
	"log"
)

func Part1(lines []string) int {
	return len(lines)
}

func Part2(numbers []int) int {
	acc := 0
	for _, n := range numbers {
		acc += n
	}
	return acc
}

func main() {
	lines := shared.ReadLinesFromFile("day00/input.txt")
	log.Printf("Part 1: there are %d numbers\n", Part1(lines))

	numbers := shared.ReadNumberFile("day00/input.txt")
	log.Printf("Part 2: sum of the numbers is %d\n", Part2(numbers))
}
