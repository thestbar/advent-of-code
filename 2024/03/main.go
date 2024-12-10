package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	// This is the regex pattern for part 1 of the challenge
	// pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	// This is the regex pattern for part 2 of the challenge
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	sum := 0

	isActive := true

	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			fmt.Println(match)

			if match == "do()" {
				isActive = true
			} else if match == "don't()" {
				isActive = false
			}

			if isActive && match != "do()" && match != "don't()" {
				values := strings.Split(match, ",")
				values[0] = values[0][4:]
				values[1] = values[1][:len(values[1])-1]

				sum += toInt(values[0]) * toInt(values[1])
			}
		}
	}

	fmt.Println("Sum:", sum)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting to int", err)
	}

	return i
}
