package main

import (
	"log"
)

func Part1(vectors []Vector) int {
	x := 0
	y := 0

	for _, v := range vectors {
		x += v.X
		y += v.Y
	}

	return x * y
}

func Part2(vectors []Vector) int {
	x := 0
	y := 0
	aim := 0

	for _, v := range vectors {
		x += v.X
		aim += v.Y

		y += v.X * aim
	}

	return x * y
}

func main() {
	numbers := ReadVectorsFromFile("day02/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
