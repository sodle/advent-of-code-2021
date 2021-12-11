package main

import (
	"aoc2021/shared"
	"log"
	"strconv"
)

func Part1(lines []string) (flashes int) {
	energyLevels := map[shared.Coordinate]int{}
	width := 0
	height := len(lines)

	for y, line := range lines {
		width = len(line)
		for x, char := range line {
			energy, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			energyLevels[shared.Coordinate{X: x, Y: y}] = energy
		}
	}

	for i := 0; i < 100; i++ {
		flashers := map[shared.Coordinate]bool{}
		for c, energy := range energyLevels {
			energyLevels[c] = energy + 1
		}

		for true {
			doneFlashing := true

			for c, energy := range energyLevels {
				if energy > 9 && !flashers[c] {
					doneFlashing = false
					flashers[c] = true
					flashes++
					for _, neighbor := range c.Neighbors(width, height, true) {
						energyLevels[neighbor]++
					}
				}
			}

			if doneFlashing {
				break
			}
		}

		for c, flashed := range flashers {
			if flashed {
				energyLevels[c] = 0
			}
		}
	}

	return
}

func Part2(lines []string) int {
	energyLevels := map[shared.Coordinate]int{}
	width := 0
	height := len(lines)

	for y, line := range lines {
		width = len(line)
		for x, char := range line {
			energy, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			energyLevels[shared.Coordinate{X: x, Y: y}] = energy
		}
	}

	i := 0
	for true {
		i++
		flashers := map[shared.Coordinate]bool{}
		for c, energy := range energyLevels {
			energyLevels[c] = energy + 1
		}

		for true {
			doneFlashing := true

			for c, energy := range energyLevels {
				if energy > 9 && !flashers[c] {
					doneFlashing = false
					flashers[c] = true
					for _, neighbor := range c.Neighbors(width, height, true) {
						energyLevels[neighbor]++
					}
				}
			}

			if doneFlashing {
				break
			}
		}

		for c, flashed := range flashers {
			if flashed {
				energyLevels[c] = 0
			}
		}

		allFlashed := true
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if !flashers[shared.Coordinate{
					X: x,
					Y: y,
				}] {
					allFlashed = false
				}
			}
		}
		if allFlashed {
			return i
		}
	}

	return -1
}

func main() {
	numbers := shared.ReadLinesFromFile("day11/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
