package main

import (
	"aoc2021/shared"
	"testing"
)

func TestPart1_1(t *testing.T) {
	input := []string{"8A004A801A8002F478"}

	want := 16
	got := Part1(input)

	if want != got {
		t.Errorf("Part 1 incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart1_2(t *testing.T) {
	input := []string{"620080001611562C8802118E34"}

	want := 12
	got := Part1(input)

	if want != got {
		t.Errorf("Part 1 incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart1_3(t *testing.T) {
	input := []string{"C0015000016115A2E0802F182340"}

	want := 23
	got := Part1(input)

	if want != got {
		t.Errorf("Part 1 incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart1_4(t *testing.T) {
	input := []string{"A0016C880162017C3686B18A3D4780"}

	want := 31
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
