package main

import (
	"aoc2021/shared"
	"log"
	"strings"
)

func Part1(lines []string) int {
	matches := 0

	for _, line := range lines {
		split := strings.Split(line, " | ")
		rhs := split[1]
		digits := strings.Fields(rhs)

		for _, digit := range digits {
			segments := len(digit)
			if segments == 2 || segments == 4 || segments == 3 || segments == 7 {
				matches++
			}
		}
	}

	return matches
}

func inCommon(a string, b string) (count int) {
	for _, m := range a {
		for _, n := range b {
			if m == n {
				count++
			}
		}
	}

	return count
}

func Part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		split := strings.Split(line, " | ")

		lhs := split[0]
		rhs := split[1]

		lhsDigits := strings.Fields(lhs)
		rhsDigits := strings.Fields(rhs)

		one := ""
		four := ""
		seven := ""
		eight := ""
		var len5 []string
		var len6 []string
		for _, digit := range lhsDigits {
			switch len(digit) {
			case 2:
				one = digit
			case 4:
				four = digit
			case 3:
				seven = digit
			case 7:
				eight = digit
			case 5:
				len5 = append(len5, digit)
			case 6:
				len6 = append(len6, digit)
			}
		}

		zero := ""
		six := ""
		nine := ""
		for _, digit := range len6 {
			switch inCommon(digit, four) {
			case 4:
				nine = digit
			case 3:
				if inCommon(digit, seven) == 3 {
					zero = digit
				} else {
					six = digit
				}
			}
		}

		two := ""
		three := ""
		five := ""
		for _, digit := range len5 {
			if inCommon(digit, nine) == 5 {
				if inCommon(digit, seven) == 3 {
					three = digit
				} else {
					five = digit
				}
			} else {
				two = digit
			}
		}

		var answerDigits [4]int
		for i, digit := range rhsDigits {
			if len(digit) == 6 && inCommon(digit, zero) == 6 {
				answerDigits[i] = 0
			} else if len(digit) == 2 && inCommon(digit, one) == 2 {
				answerDigits[i] = 1
			} else if len(digit) == 5 && inCommon(digit, two) == 5 {
				answerDigits[i] = 2
			} else if len(digit) == 5 && inCommon(digit, three) == 5 {
				answerDigits[i] = 3
			} else if len(digit) == 4 && inCommon(digit, four) == 4 {
				answerDigits[i] = 4
			} else if len(digit) == 5 && inCommon(digit, five) == 5 {
				answerDigits[i] = 5
			} else if len(digit) == 6 && inCommon(digit, six) == 6 {
				answerDigits[i] = 6
			} else if len(digit) == 3 && inCommon(digit, seven) == 3 {
				answerDigits[i] = 7
			} else if len(digit) == 7 && inCommon(digit, eight) == 7 {
				answerDigits[i] = 8
			} else {
				answerDigits[i] = 9
			}
		}

		sum += 1000*answerDigits[0] + 100*answerDigits[1] + 10*answerDigits[2] + answerDigits[3]
	}

	return sum
}

func main() {
	numbers := shared.ReadLinesFromFile("day08/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
