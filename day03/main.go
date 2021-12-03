package main

import (
	"aoc2021/shared"
	"log"
	"strconv"
)

func Gamma(bitStrings []string) int {
	gammaBitString := gammaBin(bitStrings)

	gamma, err := strconv.ParseInt(gammaBitString, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(gamma)
}

func gammaBin(bitStrings []string) string {
	// WARNING: ASSUMES ALL BIT STRINGS HAVE SAME LENGTH
	width := len(bitStrings[0])
	gammaBitString := ""

	for i := 0; i < width; i++ {
		zeros := 0
		ones := 0

		for _, bitString := range bitStrings {
			if bitString[i] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		if zeros > ones {
			gammaBitString += "0"
		} else {
			gammaBitString += "1"
		}
	}
	return gammaBitString
}

func Epsilon(bitStrings []string) int {
	epsilonBitString := epsilonBin(bitStrings)

	epsilon, err := strconv.ParseInt(epsilonBitString, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(epsilon)
}

func epsilonBin(bitStrings []string) string {
	// WARNING: ASSUMES ALL BIT STRINGS HAVE SAME LENGTH
	width := len(bitStrings[0])
	epsilonBitString := ""

	for i := 0; i < width; i++ {
		zeros := 0
		ones := 0

		for _, bitString := range bitStrings {
			if bitString[i] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		if zeros > ones {
			epsilonBitString += "1"
		} else {
			epsilonBitString += "0"
		}
	}
	return epsilonBitString
}

func Part1(bitStrings []string) int {
	gamma := Gamma(bitStrings)
	epsilon := Epsilon(bitStrings)

	return gamma * epsilon
}

func OxygenRating(bitStrings []string) int {
	haystack := bitStrings
	var needles []string

	column := 0
	for len(haystack) > 1 {
		gamma := gammaBin(haystack)
		gammaBit := int(gamma[column])

		for _, bitString := range haystack {
			if int(bitString[column]) == gammaBit {
				needles = append(needles, bitString)
			}
		}

		haystack = needles
		needles = []string{}
		column++
	}

	oxygenRating, err := strconv.ParseInt(haystack[0], 2, 64)
	if err != nil {
		panic(err)
	}
	return int(oxygenRating)
}

func CO2ScrubberRating(bitStrings []string) int {
	haystack := bitStrings
	var needles []string

	column := 0
	for len(haystack) > 1 {
		epsilon := epsilonBin(haystack)
		epsilonBit := int(epsilon[column])

		for _, bitString := range haystack {
			if int(bitString[column]) == epsilonBit {
				needles = append(needles, bitString)
			}
		}

		haystack = needles
		needles = []string{}
		column++
	}

	co2ScrubberRating, err := strconv.ParseInt(haystack[0], 2, 64)
	if err != nil {
		panic(err)
	}
	return int(co2ScrubberRating)
}

func Part2(bitStrings []string) int {
	oxygen := OxygenRating(bitStrings)
	scrubber := CO2ScrubberRating(bitStrings)

	return oxygen * scrubber
}

func main() {
	numbers := shared.ReadLinesFromFile("day03/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
