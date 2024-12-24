package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	readDesigns := false

	patterns := []string{}
	designs := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			readDesigns = true

			continue
		}

		if readDesigns {
			designs = append(designs, line)
		} else {
			patternsRow := strings.Split(line, ", ")
			patterns = append(patterns, patternsRow...)
		}
	}

	fmt.Println(patterns)
	fmt.Println(designs)

	part1Answer := Part1(patterns, designs)
	fmt.Println("Part 1 answer:", part1Answer)

  part2Answer := Part2(patterns, designs)
  fmt.Println("Part 2 answer:", part2Answer)

	data.Close()
}

func Part1(patterns []string, designs []string) int {
	validDesigns := 0
	memo := make(map[string]bool)

	for i, design := range designs {
		fmt.Println("Checking", i+1, "out of", len(designs))
		if isValidDesign(patterns, design, 0, memo) {
			validDesigns++
		}
	}
	// isValidDesign(patterns, designs[0], 0)

	return validDesigns
}

func isValidDesign(patterns []string, design string, start int, memo map[string]bool) bool {
	// fmt.Println("Deeper:", design[start:], "at", start)
	if start == len(design) {
		return true
	}

	key := fmt.Sprintf("%s-%d", design, start)
	if val, ok := memo[key]; ok {
		return val
	}

	for _, pattern := range patterns {
		patternLen := len(pattern)

		end := start + patternLen

		if end > len(design) {
			continue
		}

		// fmt.Println("Checking", design[start:end], "against", pattern, design[start:end] == pattern)

		if design[start:end] == pattern {
			if isValidDesign(patterns, design, end, memo) {
				memo[key] = true
				return true
			}
		}
	}

	memo[key] = false
	return false
}

func Part2(patterns []string, designs []string) int {
	validDesigns := 0
	memo := make(map[string]int)

	for i, design := range designs {
		fmt.Println("Checking", i+1, "out of", len(designs))

		validDesigns += numberOfValidDesigns(patterns, design, 0, memo)
	}

	return validDesigns
}

func numberOfValidDesigns(patterns []string, design string, start int, memo map[string]int) int {
	// fmt.Println("Deeper:", design[start:], "at", start)
	if start == len(design) {
		return 1
	}

	key := fmt.Sprintf("%s-%d", design, start)
	if val, ok := memo[key]; ok {
		return val
	}

	validDesigns := 0

	for _, pattern := range patterns {
		patternLen := len(pattern)

		end := start + patternLen

		if end > len(design) {
			continue
		}

		if design[start:end] == pattern {
			validDesigns += numberOfValidDesigns(patterns, design, end, memo)
		}
	}

	memo[key] = validDesigns
	return validDesigns
}
