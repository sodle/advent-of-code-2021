package main

import (
	"aoc2021/shared"
	"testing"
)

func TestRuneToInt(t *testing.T) {
	for want, r := range "0123456789" {
		got := RuneToInt(r)

		if int64(want) != got {
			t.Errorf("RuneToInt incorrect: Expected %d, got %d", want, got)
		}
	}
}

func TestPart1(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 40
	got := Part1(input)

	if want != got {
		t.Errorf("Part 1 incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 315
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
