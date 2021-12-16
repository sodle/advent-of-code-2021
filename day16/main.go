package main

import (
	"aoc2021/shared"
	"log"
)

func HexToBin(nibble rune) []bool {
	switch nibble {
	case '0':
		return []bool{false, false, false, false}
	case '1':
		return []bool{false, false, false, true}
	case '2':
		return []bool{false, false, true, false}
	case '3':
		return []bool{false, false, true, true}
	case '4':
		return []bool{false, true, false, false}
	case '5':
		return []bool{false, true, false, true}
	case '6':
		return []bool{false, true, true, false}
	case '7':
		return []bool{false, true, true, true}
	case '8':
		return []bool{true, false, false, false}
	case '9':
		return []bool{true, false, false, true}
	case 'A':
		return []bool{true, false, true, false}
	case 'B':
		return []bool{true, false, true, true}
	case 'C':
		return []bool{true, true, false, false}
	case 'D':
		return []bool{true, true, false, true}
	case 'E':
		return []bool{true, true, true, false}
	case 'F':
		return []bool{true, true, true, true}
	}
	return []bool{}
}

func Part1(lines []string) (versionSum int) {
	var bits []bool
	for _, nibble := range lines[0] {
		bits = append(bits, HexToBin(nibble)...)
	}

	pc := 0
	for pc < len(bits)-6 {
		// Version field
		if bits[pc] {
			versionSum += 4
		}
		if bits[pc+1] {
			versionSum += 2
		}
		if bits[pc+2] {
			versionSum++
		}
		pc += 3

		// Type ID field
		typeId := 0
		if bits[pc] {
			typeId += 4
		}
		if bits[pc+1] {
			typeId += 2
		}
		if bits[pc+2] {
			typeId++
		}
		pc += 3

		if typeId == 4 {
			// literal value
			for bits[pc] {
				pc += 5
			}
			pc += 5
		} else {
			// operator value
			if bits[pc] {
				// packet count mode
				pc += 12
			} else {
				pc += 16
			}
		}
	}

	return
}

func Part2(lines []string) int {
	return 0
}

func main() {
	numbers := shared.ReadLinesFromFile("day16/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
