package main

import (
	"aoc2021/shared"
	"log"
	"strconv"
	"strings"
)

type BingoCard struct {
	grid  [25]int
	daubs [25]bool
}

func NewBingoCard(lines [5]string) BingoCard {
	grid := [25]int{}
	for i, line := range lines {
		for j, column := range strings.Fields(line) {
			idx := 5*i + j
			num, err := strconv.Atoi(column)
			if err != nil {
				panic(err)
			}
			grid[idx] = num
		}
	}
	return BingoCard{grid: grid}
}

func (c *BingoCard) CheckNumber(num int) bool {
	for i, cardNum := range c.grid {
		if cardNum == num {
			c.daubs[i] = true
			return true
		}
	}
	return false
}

func (c *BingoCard) CheckBingo() bool {
	for i := 0; i < 5; i++ {
		if c.daubs[i*5] && c.daubs[i*5+1] && c.daubs[i*5+2] && c.daubs[i*5+3] && c.daubs[i*5+4] {
			return true
		}
		if c.daubs[i] && c.daubs[i+5] && c.daubs[i+10] && c.daubs[i+15] && c.daubs[i+20] {
			return true
		}
	}
	return false
}

func (c *BingoCard) SumUnmarked() (sum int) {
	for i, daub := range c.daubs {
		if !daub {
			sum += c.grid[i]
		}
	}
	return sum
}

func ReadNumbers(numberLine string) (out []int) {
	numStrings := strings.Split(numberLine, ",")
	for _, num := range numStrings {
		number, err := strconv.Atoi(num)
		if err != nil {
			panic(nil)
		}
		out = append(out, number)
	}
	return out
}

func Part1(lines []string) int {
	numbers := ReadNumbers(lines[0])

	var cards []BingoCard
	for i := 1; i < len(lines); i += 5 {
		card := NewBingoCard([5]string{
			lines[i],
			lines[i+1],
			lines[i+2],
			lines[i+3],
			lines[i+4],
		})
		cards = append(cards, card)
	}

	for _, number := range numbers {
		for i := 0; i < len(cards); i++ {
			if cards[i].CheckNumber(number) {
				if cards[i].CheckBingo() {
					return cards[i].SumUnmarked() * number
				}
			}
		}
	}

	return 0
}

func Part2(lines []string) int {
	numbers := ReadNumbers(lines[0])

	var cards []BingoCard
	for i := 1; i < len(lines); i += 5 {
		card := NewBingoCard([5]string{
			lines[i],
			lines[i+1],
			lines[i+2],
			lines[i+3],
			lines[i+4],
		})
		cards = append(cards, card)
	}

	var scores []int

	for _, number := range numbers {
		for i := 0; i < len(cards); i++ {
			if cards[i].CheckBingo() {
				continue
			}

			if cards[i].CheckNumber(number) {
				if cards[i].CheckBingo() {
					scores = append(scores, cards[i].SumUnmarked()*number)
				}
			}
		}
	}

	return scores[len(scores)-1]
}

func main() {
	numbers := shared.ReadLinesFromFile("day04/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
