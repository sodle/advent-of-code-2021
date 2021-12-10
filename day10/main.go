package main

import (
	"aoc2021/shared"
	"log"
	"sort"
)

func Part1(lines []string) int {
	score := 0

	for _, line := range lines {
		var tokens []string
		for _, charCode := range line {
			token := string(charCode)

			lastIndex := len(tokens) - 1
			lastToken := ""
			if lastIndex >= 0 {
				lastToken = tokens[lastIndex]
			}

			if token == "(" || token == "[" || token == "{" || token == "<" {
				tokens = append(tokens, token)
			} else if token == ")" {
				if lastToken == "(" {
					tokens = tokens[:lastIndex]
				} else {
					score += 3
					break
				}
			} else if token == "]" {
				if lastToken == "[" {
					tokens = tokens[:lastIndex]
				} else {
					score += 57
					break
				}
			} else if token == "}" {
				if lastToken == "{" {
					tokens = tokens[:lastIndex]
				} else {
					score += 1197
					break
				}
			} else if token == ">" {
				if lastToken == "<" {
					tokens = tokens[:lastIndex]
				} else {
					score += 25137
					break
				}
			}
		}
	}

	return score
}

func Part2(lines []string) int {
	var scores []int

	for _, line := range lines {
		var tokens []string
		valid := true

		for _, charCode := range line {
			token := string(charCode)

			lastIndex := len(tokens) - 1
			lastToken := ""
			if lastIndex >= 0 {
				lastToken = tokens[lastIndex]
			}

			if token == "(" || token == "[" || token == "{" || token == "<" {
				tokens = append(tokens, token)
			} else if token == ")" {
				if lastToken == "(" {
					tokens = tokens[:lastIndex]
				} else {
					valid = false
					break
				}
			} else if token == "]" {
				if lastToken == "[" {
					tokens = tokens[:lastIndex]
				} else {
					valid = false
					break
				}
			} else if token == "}" {
				if lastToken == "{" {
					tokens = tokens[:lastIndex]
				} else {
					valid = false
					break
				}
			} else if token == ">" {
				if lastToken == "<" {
					tokens = tokens[:lastIndex]
				} else {
					valid = false
					break
				}
			}
		}

		if valid {
			score := 0
			for i := 0; i < len(tokens); i++ {
				score *= 5

				token := tokens[len(tokens)-1-i]
				switch token {
				case "(":
					score += 1
				case "[":
					score += 2
				case "{":
					score += 3
				case "<":
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func main() {
	numbers := shared.ReadLinesFromFile("day10/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
