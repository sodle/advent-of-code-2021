package main

import (
	"aoc2021/shared"
	"github.com/yourbasic/graph"
	"log"
)

func RuneToInt(r rune) int64 {
	return int64(r - 48)
}

func CoordinateToFlat(c shared.Coordinate, gridWidth int) int {
	return c.Y*gridWidth + c.X
}

func Part1(lines []string) int {
	height := len(lines)
	width := len(lines[0])

	g := graph.New(height * width)

	for y, line := range lines {
		for x, _ := range line {
			c := shared.Coordinate{
				X: x,
				Y: y,
			}
			myIdx := CoordinateToFlat(c, width)
			neighbors := c.Neighbors(width, height, false)
			for _, neighbor := range neighbors {
				neighborCost := RuneToInt(rune(lines[neighbor.Y][neighbor.X]))
				neighborIdx := CoordinateToFlat(neighbor, width)
				g.AddCost(myIdx, neighborIdx, neighborCost)
			}
		}
	}

	_, cost := graph.ShortestPath(g, 0, width*height-1)

	return int(cost)
}

func Part2(lines []string) int {
	height := len(lines)
	width := len(lines[0])

	fullHeight := 5 * height
	fullWidth := 5 * width
	g := graph.New(fullHeight * fullWidth)

	for y := 0; y < fullHeight; y++ {
		for x := 0; x < fullWidth; x++ {
			c := shared.Coordinate{
				X: x,
				Y: y,
			}
			myIdx := CoordinateToFlat(c, fullWidth)
			neighbors := c.Neighbors(fullWidth, fullHeight, false)
			for _, neighbor := range neighbors {
				neighborCost := RuneToInt(rune(lines[neighbor.Y%height][neighbor.X%width]))
				widthPenalty := neighbor.X / width
				heightPenalty := neighbor.Y / height
				neighborCost += int64(widthPenalty + heightPenalty)
				for neighborCost > 9 {
					neighborCost -= 9
				}

				neighborIdx := CoordinateToFlat(neighbor, fullWidth)
				g.AddCost(myIdx, neighborIdx, neighborCost)
			}
		}
	}

	_, cost := graph.ShortestPath(g, 0, fullWidth*fullHeight-1)

	return int(cost)
}

func main() {
	numbers := shared.ReadLinesFromFile("day15/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
