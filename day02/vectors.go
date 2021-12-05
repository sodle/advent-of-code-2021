package main

import (
	"aoc2021/shared"
	"strconv"
	"strings"
)

type Vector struct {
	X int
	Y int
}

func forward(magnitude int) Vector {
	return Vector{
		magnitude,
		0,
	}
}

func backward(magnitude int) Vector {
	return Vector{
		-magnitude,
		0,
	}
}

func up(magnitude int) Vector {
	return Vector{
		X: 0,
		Y: -magnitude,
	}
}

func down(magnitude int) Vector {
	return Vector{
		X: 0,
		Y: magnitude,
	}
}

func ReadVectorsFromFile(path string) (out []Vector) {
	lines := shared.ReadLinesFromFile(path)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		direction := tokens[0]
		magnitude, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			out = append(out, forward(magnitude))
		case "backward":
			out = append(out, backward(magnitude))
		case "up":
			out = append(out, up(magnitude))
		case "down":
			out = append(out, down(magnitude))
		}
	}
	return out
}
