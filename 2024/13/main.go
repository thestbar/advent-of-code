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

	lineCounter := 0
	a := [2]int{0, 0}
	b := [2]int{0, 0}
	prizes := [2]int{0, 0}

	totalCost := 0

	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		if lineCounter == 0 {
			// Calculate A button
			a = FetchValue(line, a)
		} else if lineCounter == 1 {
			// Calculate B button
			b = FetchValue(line, b)
		} else if lineCounter == 2 {
			// Get prizes
			prizes = FetchValue(line, prizes)
		} else {
			// Find the answer

			// Part 1
			x, y := Part1(a[0], a[1], b[0], b[1], prizes[0]+10000000000000, prizes[1]+10000000000000)
			cost := 3*x + y
			fmt.Println("Cost: ", cost)

			totalCost += cost

			// Move to next machine
			lineCounter = -1

			fmt.Println("Going to next machine")
		}
		lineCounter++
	}

	x, y := Part1(a[0], a[1], b[0], b[1], prizes[0]+10000000000000, prizes[1]+10000000000000)
	cost := 3*x + y
	fmt.Println("Cost: ", cost)

	totalCost += cost
	fmt.Println("Total Cost: ", totalCost)
}

func Part1(a1, a2, b1, b2, c, d int) (int, int) {
	x := (b1*d - c*b2) / (b1*a2 - a1*b2)
	y := (c - a1*x) / b1

	// Validate the answer
	if a1*x+b1*y != c || a2*x+b2*y != d {
		return 0, 0
	}

	return x, y
}

func FetchValue(line string, arr [2]int) [2]int {
	stringArray := strings.Split(line, ":")
	stringArray = strings.Split(stringArray[1], ",")
	xVal := stringArray[0][3:]
	yVal := stringArray[1][3:]
	arr[0], _ = strconv.Atoi(xVal)
	arr[1], _ = strconv.Atoi(yVal)

	return arr
}
