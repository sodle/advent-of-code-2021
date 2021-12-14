package main

import (
	"aoc2021/shared"
	"log"
	"math"
	"sort"
	"strings"
)

type PolymerInserter struct {
	rules map[string]string
}

func NewPolymerInserter() PolymerInserter {
	return PolymerInserter{
		rules: map[string]string{},
	}
}

func (pi *PolymerInserter) AddRule(rule string) {
	fields := strings.Fields(rule)
	pair := fields[0]
	insert := fields[2]
	pi.rules[pair] = insert
}

func (pi *PolymerInserter) Insert(template string) (out string) {
	for i := 0; i < len(template)-1; i++ {
		p0 := string(template[i])
		p1 := string(template[i+1])

		pair := p0 + p1
		insert := pi.rules[pair]

		if i == 0 {
			out += p0
		}
		out += insert + p1
	}
	return
}

func Part1(lines []string) int {
	pi := NewPolymerInserter()
	template := ""

	for i, line := range lines {
		if i == 0 {
			template = line
		} else {
			pi.AddRule(line)
		}
	}

	for i := 0; i < 10; i++ {
		template = pi.Insert(template)
	}

	frequencies := [26]int{}
	for _, p := range template {
		idx := p - 65
		frequencies[idx]++
	}
	frequenciesSlice := frequencies[:]
	sort.Ints(frequenciesSlice)

	leastCommon := 0
	mostCommon := frequencies[25]
	for _, f := range frequencies {
		if f > 0 {
			leastCommon = f
			break
		}
	}

	return mostCommon - leastCommon
}

type PolymerTree struct {
	pairFrequencies [26][26]int
	atomFrequencies [26]int
	rules           [26][26]int
}

func NewPolymerTree(polymer string) PolymerTree {
	tree := PolymerTree{}
	for i := 0; i < len(polymer)-1; i++ {
		// convert char codes to their A-Z indices
		a0 := polymer[i] - 65
		a1 := polymer[i+1] - 65

		if i == 0 {
			tree.atomFrequencies[a0]++
		}
		tree.atomFrequencies[a1]++
		tree.pairFrequencies[a0][a1]++
	}
	return tree
}

func (pt *PolymerTree) AddRule(rule string) {
	a0 := rule[0] - 65
	a1 := rule[1] - 65
	insert := rule[6] - 65
	pt.rules[a0][a1] = int(insert)
}

func (pt *PolymerTree) Insert() PolymerTree {
	newTree := PolymerTree{rules: pt.rules, atomFrequencies: pt.atomFrequencies}

	for a0, subRules := range pt.rules {
		for a1, insert := range subRules {
			frequency := pt.pairFrequencies[a0][a1]
			newTree.atomFrequencies[insert] += frequency
			newTree.pairFrequencies[a0][insert] += frequency
			newTree.pairFrequencies[insert][a1] += frequency
		}
	}

	return newTree
}

func Part2(lines []string) int {
	template := lines[0]
	pt := NewPolymerTree(template)

	for i, line := range lines {
		if i == 0 {
			continue
		} else {
			pt.AddRule(line)
		}
	}

	for i := 0; i < 40; i++ {
		pt = pt.Insert()
	}

	lowest := math.MaxInt
	highest := 0

	for _, frequency := range pt.atomFrequencies {
		if frequency > 0 {
			if frequency < lowest {
				lowest = frequency
			}
			if frequency > highest {
				highest = frequency
			}
		}
	}

	return highest - lowest
}

func main() {
	numbers := shared.ReadLinesFromFile("day14/input.txt")
	log.Printf("Part 1: %d\n", Part1(numbers))
	log.Printf("Part 2: %d\n", Part2(numbers))
}
