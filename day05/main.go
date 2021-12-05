package main

import (
	"aoc2021/shared"
	"log"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func Part1(lines []string) int {
	coords := make(map[Coordinate]int)

	for _, line := range lines {
		points := strings.Split(line, " -> ")

		point1 := strings.Split(points[0], ",")
		point2 := strings.Split(points[1], ",")

		x1, err := strconv.Atoi(point1[0])
		if err != nil {
			panic(err)
		}
		y1, err := strconv.Atoi(point1[1])
		if err != nil {
			panic(err)
		}

		x2, err := strconv.Atoi(point2[0])
		if err != nil {
			panic(err)
		}
		y2, err := strconv.Atoi(point2[1])
		if err != nil {
			panic(err)
		}

		if x1 == x2 {
			yStart := math.Min(float64(y1), float64(y2))
			yEnd := math.Max(float64(y1), float64(y2))

			for y := yStart; y <= yEnd; y++ {
				coords[Coordinate{
					x: x1,
					y: int(y),
				}]++
			}
		} else if y1 == y2 {
			xStart := math.Min(float64(x1), float64(x2))
			xEnd := math.Max(float64(x1), float64(x2))

			for x := xStart; x <= xEnd; x++ {
				coords[Coordinate{
					x: int(x),
					y: y1,
				}]++
			}
		}
	}

	dupeCount := 0
	for _, count := range coords {
		if count > 1 {
			dupeCount++
		}
	}

	return dupeCount
}

func Part2(lines []string) int {
	coords := make(map[Coordinate]int)

	for _, line := range lines {
		points := strings.Split(line, " -> ")

		point1 := strings.Split(points[0], ",")
		point2 := strings.Split(points[1], ",")

		x1, err := strconv.Atoi(point1[0])
		if err != nil {
			panic(err)
		}
		y1, err := strconv.Atoi(point1[1])
		if err != nil {
			panic(err)
		}

		x2, err := strconv.Atoi(point2[0])
		if err != nil {
			panic(err)
		}
		y2, err := strconv.Atoi(point2[1])
		if err != nil {
			panic(err)
		}

		if x1 == x2 {
			yStart := math.Min(float64(y1), float64(y2))
			yEnd := math.Max(float64(y1), float64(y2))

			for y := yStart; y <= yEnd; y++ {
				coords[Coordinate{
					x: x1,
					y: int(y),
				}]++
			}
		} else if y1 == y2 {
			xStart := math.Min(float64(x1), float64(x2))
			xEnd := math.Max(float64(x1), float64(x2))

			for x := xStart; x <= xEnd; x++ {
				coords[Coordinate{
					x: int(x),
					y: y1,
				}]++
			}
		} else {
			length := int(math.Abs(float64(x1 - x2)))

			xDirection := (x2 - x1) / length
			yDirection := (y2 - y1) / length

			for i := 0; i <= length; i++ {
				coords[Coordinate{
					x: x1 + (xDirection * i),
					y: y1 + (yDirection * i),
				}]++
			}
		}
	}

	dupeCount := 0
	for _, count := range coords {
		if count > 1 {
			dupeCount++
		}
	}

	return dupeCount
}

func main() {
	numbers := shared.ReadLinesFromFile("day05/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
