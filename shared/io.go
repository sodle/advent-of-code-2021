package shared

import (
	"os"
	"strconv"
	"strings"
)

func ReadLinesFromFile(path string) (out []string) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileText := string(fileData)
	fileLines := strings.Split(fileText, "\n")

	var filtered []string
	for _, line := range fileLines {
		if len(strings.TrimSpace(line)) > 0 {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func ReadNumberFile(path string) (out []int) {
	fileLines := ReadLinesFromFile(path)

	for _, line := range fileLines {
		if len(line) > 0 {
			number, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			out = append(out, number)
		}
	}

	return out
}
