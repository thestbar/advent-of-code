package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File open error", err)
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	part2Sum := 0

	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		splitedLine := strings.Split(line, ": ")
		answer, _ := strconv.Atoi(splitedLine[0])
		equation := splitedLine[1]

		terms := []int{}

		for _, term := range strings.Split(equation, " ") {
			intTerm, _ := strconv.Atoi(term)
			terms = append(terms, intTerm)
		}
		if solvable(terms, 1, terms[0], answer) {
			sum += answer
		}

		if solvableWithConcat(terms, 1, terms[0], answer) {
			part2Sum += answer
		}
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Part 2 Sum:", part2Sum)
}

func solvable(terms []int, index, currentSum, answer int) bool {
	if index == len(terms) && currentSum == answer {
		return true
	}

	if index == len(terms) {
		return false
	}

	if solvable(terms, index+1, currentSum+terms[index], answer) {
		return true
	}

	if solvable(terms, index+1, currentSum*terms[index], answer) {
		return true
	}

	return false
}

func solvableWithConcat(terms []int, index, currentSum, answer int) bool {
	if index == len(terms) && currentSum == answer {
		return true
	}

	if index == len(terms) {
		return false
	}

	if solvableWithConcat(terms, index+1, currentSum+terms[index], answer) {
		return true
	}

	if solvableWithConcat(terms, index+1, currentSum*terms[index], answer) {
		return true
	}

	concatSumString := strconv.Itoa(currentSum) + strconv.Itoa(terms[index])
	concatSum, _ := strconv.Atoi(concatSumString)

	if solvableWithConcat(terms, index+1, concatSum, answer) {
		return true
	}

	return false
}
