package main

import (
	"aoc2021/shared"
	"testing"
)

func TestPart1(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 5
	got := Part1(input)

	if want != got {
		t.Errorf("Day 00 Part 1 incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	input := shared.ReadNumberFile("test_input.txt")

	want := 15
	got := Part2(input)

	if want != got {
		t.Errorf("Day 00 Part 2 incorrect: Expected %d, got %d", want, got)
	}
}
