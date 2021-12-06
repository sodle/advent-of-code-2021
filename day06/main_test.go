package main

import (
	"aoc2021/shared"
	"testing"
)

func TestPart1(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 5934
	got := Part1(input)

	if want != got {
		t.Errorf("Part 1 incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 26984457539
	got := Part2(input)

	if want != got {
		t.Errorf("Part 2 incorrect: Expected %d, got %d", want, got)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := shared.ReadLinesFromFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := shared.ReadLinesFromFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part2(input)
	}
}
