package main

import (
	"aoc2021/shared"
	"testing"
)

func TestNewBingoCard(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	cardLines := [5]string{
		input[1],
		input[2],
		input[3],
		input[4],
		input[5],
	}
	card := NewBingoCard(cardLines)

	expectedGrid := [25]int{
		22, 13, 17, 11, 0,
		8, 2, 23, 4, 24,
		21, 9, 14, 16, 7,
		6, 10, 3, 18, 5,
		1, 12, 20, 15, 19,
	}

	expectedDaubs := [25]bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}

	if expectedGrid != card.grid {
		t.Errorf("Failed to read bingo card, incorrect numbers!")
	}

	if expectedDaubs != card.daubs {
		t.Errorf("Failed to read bingo card, incorrect daubs!")
	}
}

func TestBingoCard_CheckNumber(t *testing.T) {
	card := BingoCard{
		grid: [25]int{
			22, 13, 17, 11, 0,
			8, 2, 23, 4, 24,
			21, 9, 14, 16, 7,
			6, 10, 3, 18, 5,
			1, 12, 20, 15, 19,
		},
	}
	expectedDaubs := [25]bool{
		false, false, false, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}

	if !card.CheckNumber(23) {
		t.Errorf("Error checking number - expected 23 to be a match!")
	}

	if card.CheckNumber(42) {
		t.Errorf("Error checking number - expected 42 to not be a match!")
	}

	if card.daubs != expectedDaubs {
		t.Errorf("Error checking number - daub not placed!")
	}
}

func TestBingoCard_CheckBingo(t *testing.T) {
	horizontalBingo := BingoCard{daubs: [25]bool{
		true, true, true, true, true,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}
	if !horizontalBingo.CheckBingo() {
		t.Errorf("Error checking bingo - expected horizontal win!")
	}

	verticalBingo := BingoCard{daubs: [25]bool{
		true, false, false, false, false,
		true, false, false, false, false,
		true, false, false, false, false,
		true, false, false, false, false,
		true, false, false, false, false,
	}}
	if !verticalBingo.CheckBingo() {
		t.Errorf("Error checking bingo - expected vertical win!")
	}
}

func TestBingoCard_SumUnmarked(t *testing.T) {
	c := BingoCard{
		grid: [25]int{
			22, 13, 17, 11, 0,
			8, 2, 23, 4, 24,
			21, 9, 14, 16, 7,
			6, 10, 3, 18, 5,
			1, 12, 20, 15, 19,
		},
		daubs: [25]bool{
			true, true, true, true, true,
			false, false, false, false, false,
			false, false, false, false, false,
			false, false, false, false, false,
			false, false, false, false, false,
		},
	}

	want := 237
	got := c.SumUnmarked()

	if want != got {
		t.Errorf("Error summing unmarked squares - wanted %d, got %d", want, got)
	}
}

func TestPart1(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 4512
	got := Part1(input)

	if want != got {
		t.Errorf("Part 1 incorrect: Expected %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	input := shared.ReadLinesFromFile("test_input.txt")

	want := 1924
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
