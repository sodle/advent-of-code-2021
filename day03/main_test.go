package main

import (
	"aoc2021/shared"
	"testing"
)

func TestGamma(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 22
	got := Gamma(input)

	if want != got {
		t.Errorf("Gamma value incorrect: Expected %d, got %d", want, got)
	}
}

func TestEpsilon(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 9
	got := Epsilon(input)

	if want != got {
		t.Errorf("Epsilon value incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart1(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 198
	got := Part1(input)

	if want != got {
		t.Errorf("Day 03 Part 1 incorrect: Expected %d, got %d", want, got)
	}
}

func TestOxygenRating(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 23
	got := OxygenRating(input)

	if want != got {
		t.Errorf("Oxygen Rating incorrect: Expected %d, got %d", want, got)
	}
}

func TestCO2ScrubberRating(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 10
	got := CO2ScrubberRating(input)

	if want != got {
		t.Errorf("CO2 Scrubber Rating incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 230
	got := Part2(input)

	if want != got {
		t.Errorf("Day 02 Part 2 incorrect: Expected %d, got %d", want, got)
	}
}

func BenchmarkGamma(b *testing.B) {
	input := shared.ReadLinesFromFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Gamma(input)
	}
}

func BenchmarkEpsilon(b *testing.B) {
	input := shared.ReadLinesFromFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Epsilon(input)
	}
}

func BenchmarkOxygenRating(b *testing.B) {
	input := shared.ReadLinesFromFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OxygenRating(input)
	}
}

func BenchmarkCO2ScrubberRating(b *testing.B) {
	input := shared.ReadLinesFromFile("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CO2ScrubberRating(input)
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
