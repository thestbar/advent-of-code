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
	// data, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("File open error", err)
	}

	fileScanner := bufio.NewScanner(data)

	fileScanner.Split(bufio.ScanLines)

	safeCount := 0

	for fileScanner.Scan() {
		var line string = fileScanner.Text()
		result := strings.Split(line, " ")
		intResult := make([]int, len(result))

		for i := 0; i < len(result); i++ {
			intResult[i] = toInt(result[i])
		}

		if isSafe(intResult) {
			safeCount++
		} else {
			for r := 0; r < len(result); r++ {
				intResult := []int{}

				for j := 0; j < len(result); j++ {
					if j != r {
						intResult = append(intResult, toInt(result[j]))
					}
				}

				if isSafe(intResult) {
					safeCount++

					break
				}
			}
		}
	}

	fmt.Println("Number of safe ops: ", safeCount)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func isSafe(array []int) bool {
	fmt.Println(array)
	if array[0] == array[1] {
		return false
	}

	isIncreasing := array[0] < array[1]

	for i := 1; i < len(array); i++ {
		prev := array[i-1]
		curr := array[i]

		if isIncreasing && (curr <= prev || curr-prev > 3) {
			return false
		}

		if !isIncreasing && (curr >= prev || prev-curr > 3) {
			return false
		}
	}

	return true
}
