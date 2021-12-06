package main

import (
	"aoc2021/shared"
	"log"
	"strconv"
	"strings"
)

type LanternFish struct {
	age int
}

func NewLanternFish() LanternFish {
	return LanternFish{age: 8}
}

func (l *LanternFish) BreedingCycle() bool {
	if l.age == 0 {
		l.age = 6
		return true
	} else {
		l.age--
		return false
	}
}

func Part1(lines []string) int {
	line := lines[0]
	ages := strings.Split(line, ",")

	var fish []LanternFish
	for _, age := range ages {
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			panic(err)
		}

		fish = append(fish, LanternFish{
			age: ageInt,
		})
	}

	for day := 0; day < 80; day++ {
		var newFish []LanternFish

		for i := 0; i < len(fish); i++ {
			if fish[i].BreedingCycle() {
				newFish = append(newFish, NewLanternFish())
			}
		}

		fish = append(fish, newFish...)
	}

	return len(fish)
}

// Part2 Used a hint for part 2
func Part2(lines []string) int {
	line := lines[0]
	ages := strings.Split(line, ",")

	calendar := [9]int{}

	for _, age := range ages {
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			panic(err)
		}

		calendar[ageInt]++
	}

	for day := 0; day < 256; day++ {
		calendar = [9]int{
			calendar[1], calendar[2], calendar[3], calendar[4], calendar[5],
			calendar[6], calendar[7] + calendar[0], calendar[8], calendar[0],
		}
	}

	sum := 0
	for _, count := range calendar {
		sum += count
	}

	return sum
}

func main() {
	numbers := shared.ReadLinesFromFile("day06/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
