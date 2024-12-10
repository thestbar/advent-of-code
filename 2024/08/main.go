package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File open error", err)
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	grid := [][]string{}
	antennaPositionsByFreq := map[string][][2]int{}

	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		row := []string{}
		for _, char := range line {
			value := string(char)
			row = append(row, value)

			if value != "." {
				antennaPositionsByFreq[value] = append(antennaPositionsByFreq[value], [2]int{len(grid), len(row) - 1})
			}
		}

		grid = append(grid, row)
	}

	printGrid(grid)
	fmt.Println("--------------------")
	fmt.Println(antennaPositionsByFreq)

	// numberOfAntinodes := 0
	numberOfAntinodesWithHarmonics := 0

	// Now iterate over the antennas and find the antinodes
	for _, positions := range antennaPositionsByFreq {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				// numberOfAntinodes += calculateResonance(grid, positions[i][0], positions[i][1], positions[j][0], positions[j][1])
				numberOfAntinodesWithHarmonics += calculateResonanceWithHarmonics(
					grid, positions[i][0], positions[i][1], positions[j][0], positions[j][1])
			}
		}
	}

	// Now count to the result all the nodes that are still not counted.
	for _, row := range grid {
		for _, char := range row {
			if char != "." && char != "#" {
				numberOfAntinodesWithHarmonics++
			}
		}
	}

	fmt.Println("--------------------")
	fmt.Println()
	fmt.Println("Final grid")
	printGrid(grid)
	fmt.Println("--------------------")
	// fmt.Println("Number of antinodes", numberOfAntinodes)
	fmt.Println("Number of antinodes with harmonics", numberOfAntinodesWithHarmonics)
}

func calculateResonance(grid [][]string, antenna1Row, antenna1Col, antenna2Row, antenna2Col int) int {
	// Find resonance positions

	// Here is the formula to find the antinodes
	// 1, 8 r1, c1
	// 2, 5 r2, c2
	//
	// 2 + (2 - 1), 5 + (5 - 8) = 3, 2
	// r2 + (r2 - r1), c2 + (c2 - c1)
	//
	// 1 - (2 - 1), 8 - (5 - 8) = 0, 11
	// r1 - (r2 - r1), c1 - (c2 - c1)

	numberOfAntinodes := 0

	antinode1Row := antenna2Row + (antenna2Row - antenna1Row)
	antinode1Col := antenna2Col + (antenna2Col - antenna1Col)
	if antinode1Row >= 0 && antinode1Row < len(grid) && antinode1Col >= 0 && antinode1Col < len(grid[0]) {
		// Do not count an antinode mulitple times
		if grid[antinode1Row][antinode1Col] != "#" {
			numberOfAntinodes++
		}
		grid[antinode1Row][antinode1Col] = "#"
	}

	antinode2Row := antenna1Row - (antenna2Row - antenna1Row)
	antinode2Col := antenna1Col - (antenna2Col - antenna1Col)
	if antinode2Row >= 0 && antinode2Row < len(grid) && antinode2Col >= 0 && antinode2Col < len(grid[0]) {
		// Do not count an antinode mulitple times
		if grid[antinode2Row][antinode2Col] != "#" {
			numberOfAntinodes++
		}
		grid[antinode2Row][antinode2Col] = "#"
	}

	return numberOfAntinodes
}

func calculateResonanceWithHarmonics(grid [][]string, antenna1Row, antenna1Col, antenna2Row, antenna2Col int) int {
	numberOfAntinodes := 0

	antinode1Row := antenna2Row + (antenna2Row - antenna1Row)
	antinode1Col := antenna2Col + (antenna2Col - antenna1Col)

	for antinode1Row >= 0 && antinode1Row < len(grid) && antinode1Col >= 0 && antinode1Col < len(grid[0]) {
		// Do not count an antinode mulitple times
		if grid[antinode1Row][antinode1Col] != "#" {
			numberOfAntinodes++
		}
		grid[antinode1Row][antinode1Col] = "#"

		antinode1Row += antenna2Row - antenna1Row
		antinode1Col += antenna2Col - antenna1Col
	}

	antinode2Row := antenna1Row - (antenna2Row - antenna1Row)
	antinode2Col := antenna1Col - (antenna2Col - antenna1Col)

	for antinode2Row >= 0 && antinode2Row < len(grid) && antinode2Col >= 0 && antinode2Col < len(grid[0]) {
		// Do not count an antinode mulitple times
		if grid[antinode2Row][antinode2Col] != "#" {
			numberOfAntinodes++
		}
		grid[antinode2Row][antinode2Col] = "#"

		antinode2Row -= antenna2Row - antenna1Row
		antinode2Col -= antenna2Col - antenna1Col
	}

	return numberOfAntinodes
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
}
