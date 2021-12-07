package main

import (
	"aoc2021/shared"
	"log"
	"math"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	xStrings := strings.Split(lines[0], ",")

	xMax := 0
	xMin := math.MaxInt

	var positions []int

	for _, xString := range xStrings {
		x, err := strconv.Atoi(xString)
		if err != nil {
			panic(err)
		}

		if x > xMax {
			xMax = x
		}

		if x < xMin {
			xMin = x
		}

		positions = append(positions, x)
	}

	shortest := math.MaxInt
	for meetX := xMin; meetX <= xMax; meetX++ {
		cost := 0
		for _, x := range positions {
			distance := x - meetX
			if distance < 0 {
				distance *= -1
			}
			cost += distance
		}

		if cost < shortest {
			shortest = cost
		}
	}

	return shortest
}

func Part2(lines []string) int {
	xStrings := strings.Split(lines[0], ",")

	xMax := 0
	xMin := math.MaxInt

	var positions []int

	for _, xString := range xStrings {
		x, err := strconv.Atoi(xString)
		if err != nil {
			panic(err)
		}

		if x > xMax {
			xMax = x
		}

		if x < xMin {
			xMin = x
		}

		positions = append(positions, x)
	}

	shortest := math.MaxInt
	for meetX := xMin; meetX <= xMax; meetX++ {
		cost := 0
		for _, x := range positions {
			distance := x - meetX
			if distance < 0 {
				distance *= -1
			}

			for d := 1; d <= distance; d++ {
				cost += d
			}
		}

		if cost < shortest {
			shortest = cost
		}
	}

	return shortest
}

func main() {
	numbers := shared.ReadLinesFromFile("day07/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
