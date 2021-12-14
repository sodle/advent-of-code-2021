package main

import (
	"aoc2021/shared"
	"log"
	"math"
	"strconv"
	"strings"
)

type Sheet struct {
	dots map[shared.Coordinate]bool
}

func NewSheet() Sheet {
	return Sheet{
		dots: map[shared.Coordinate]bool{},
	}
}

func (s *Sheet) AddDot(dot shared.Coordinate) {
	s.dots[dot] = true
}

func (s Sheet) FoldX(axis int) (folded Sheet) {
	folded = NewSheet()
	for dot, _ := range s.dots {
		x := dot.X
		if x > axis {
			x = 2*axis - x
		}

		folded.AddDot(shared.Coordinate{
			X: x,
			Y: dot.Y,
		})
	}
	return
}

func (s Sheet) FoldY(axis int) (folded Sheet) {
	folded = NewSheet()
	for dot, _ := range s.dots {
		y := dot.Y
		if y > axis {
			y = 2*axis - y
		}

		folded.AddDot(shared.Coordinate{
			X: dot.X,
			Y: y,
		})
	}
	return
}

func (s Sheet) CountDots() (count int) {
	for _, _ = range s.dots {
		count++
	}
	return
}

func Part1(lines []string) int {
	s := NewSheet()

	for _, line := range lines {
		if strings.HasPrefix(line, "fold") {
			instruction := strings.Split(line, "=")

			axis, err := strconv.Atoi(instruction[1])
			if err != nil {
				panic(err)
			}

			if instruction[0] == "fold along x" {
				s = s.FoldX(axis)
			} else {
				s = s.FoldY(axis)
			}

			break
		} else {
			pair := strings.Split(line, ",")

			x, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(pair[1])
			if err != nil {
				panic(err)
			}

			s.AddDot(shared.Coordinate{
				X: x,
				Y: y,
			})
		}
	}

	return s.CountDots()
}

func Part2(lines []string) (out string) {
	s := NewSheet()

	for _, line := range lines {
		if strings.HasPrefix(line, "fold") {
			instruction := strings.Split(line, "=")

			axis, err := strconv.Atoi(instruction[1])
			if err != nil {
				panic(err)
			}

			if instruction[0] == "fold along x" {
				s = s.FoldX(axis)
			} else {
				s = s.FoldY(axis)
			}
		} else {
			pair := strings.Split(line, ",")

			x, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(pair[1])
			if err != nil {
				panic(err)
			}

			s.AddDot(shared.Coordinate{
				X: x,
				Y: y,
			})
		}
	}

	right := 0
	bottom := 0
	left := math.MaxInt
	top := math.MaxInt
	for dot, _ := range s.dots {
		if dot.X > right {
			right = dot.X
		}
		if dot.Y > bottom {
			bottom = dot.Y
		}
		if dot.X < left {
			left = dot.X
		}
		if dot.Y < top {
			top = dot.Y
		}
	}

	for y := top; y <= bottom; y++ {
		out += "\n"
		for x := left; x <= right; x++ {
			if s.dots[shared.Coordinate{
				X: x,
				Y: y,
			}] {
				out += "#"
			} else {
				out += " "
			}
		}
	}

	return
}

func main() {
	numbers := shared.ReadLinesFromFile("day13/test_input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %s\n", Part2(numbers))
}
