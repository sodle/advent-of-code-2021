package main

import (
	"aoc2021/shared"
	"log"
	"math"
	"sort"
	"strconv"
)

func Part1(lines []string) int {
	grid := map[shared.Coordinate]int{}

	height := len(lines)
	width := 0
	for i, line := range lines {
		width = len(line)
		for j, cell := range line {
			elevation, err := strconv.Atoi(string(cell))
			if err != nil {
				panic(err)
			}

			grid[shared.Coordinate{
				X: j,
				Y: i,
			}] = elevation
		}
	}

	risk := 0

	for coord, elevation := range grid {
		left := math.MaxInt
		right := math.MaxInt
		up := math.MaxInt
		down := math.MaxInt

		if coord.X > 0 {
			left = grid[shared.Coordinate{
				coord.X - 1,
				coord.Y,
			}]
		}
		if coord.Y > 0 {
			up = grid[shared.Coordinate{
				coord.X,
				coord.Y - 1,
			}]
		}
		if coord.X < width-1 {
			right = grid[shared.Coordinate{
				coord.X + 1,
				coord.Y,
			}]
		}
		if coord.Y < height-1 {
			down = grid[shared.Coordinate{
				coord.X,
				coord.Y + 1,
			}]
		}

		if elevation < left && elevation < right && elevation < up && elevation < down {
			risk += elevation + 1
		}
	}

	return risk
}

// Part2 not yet working
func Part2(lines []string) int {
	grid := map[shared.Coordinate]int{}

	height := len(lines)
	width := 0
	for i, line := range lines {
		width = len(line)
		for j, cell := range line {
			elevation, err := strconv.Atoi(string(cell))
			if err != nil {
				panic(err)
			}

			grid[shared.Coordinate{
				X: j,
				Y: i,
			}] = elevation
		}
	}

	regionMap := map[shared.Coordinate]int{}
	regionSizes := map[int]int{}
	currentRegion := 1
	for coord, elevation := range grid {
		if elevation == 9 {
			continue
		}

		region := 0
		if regionMap[coord] > 0 {
			region = regionMap[coord]
		} else {
			if coord.X > 0 {
				left := shared.Coordinate{
					coord.X - 1,
					coord.Y,
				}
				if regionMap[left] > 0 {
					region = regionMap[left]
				}
			}
			if coord.Y > 0 {
				up := shared.Coordinate{
					coord.X,
					coord.Y - 1,
				}
				if regionMap[up] > 0 {
					region = regionMap[up]
				}
			}
			if coord.X < width-1 {
				right := shared.Coordinate{
					coord.X + 1,
					coord.Y,
				}
				if regionMap[right] > 0 {
					region = regionMap[right]
				}
			}
			if coord.Y < height-1 {
				down := shared.Coordinate{
					coord.X,
					coord.Y + 1,
				}
				if regionMap[down] > 0 {
					region = regionMap[down]
				}
			}
		}

		if region == 0 {
			region = currentRegion
			currentRegion++

			regionMap[coord] = region
			regionSizes[region]++

			if coord.X > 0 {
				left := shared.Coordinate{
					coord.X - 1,
					coord.Y,
				}
				if grid[left] < 9 && regionMap[left] == 0 {
					regionMap[left] = region
					regionSizes[region]++
				}
			}
			if coord.Y > 0 {
				up := shared.Coordinate{
					coord.X,
					coord.Y - 1,
				}
				if grid[up] < 9 && regionMap[up] == 0 {
					regionMap[up] = region
					regionSizes[region]++
				}
			}
			if coord.X < width-1 {
				right := shared.Coordinate{
					coord.X + 1,
					coord.Y,
				}
				if grid[right] < 9 && regionMap[right] == 0 {
					regionMap[right] = region
					regionSizes[region]++
				}
			}
			if coord.Y < height-1 {
				down := shared.Coordinate{
					coord.X,
					coord.Y + 1,
				}
				if grid[down] < 9 && regionMap[down] == 0 {
					regionMap[down] = region
					regionSizes[region]++
				}
			}
		}
	}

	var regionSizeList []int
	for _, size := range regionSizes {
		regionSizeList = append(regionSizeList, size)
	}
	sort.Ints(regionSizeList)
	regionCount := len(regionSizeList)

	return regionSizeList[regionCount-1] * regionSizeList[regionCount-2] * regionSizeList[regionCount-3]
}

func main() {
	numbers := shared.ReadLinesFromFile("day09/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
