package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var memo = map[string]int{}

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	grid := [][]string{}

	startRow := 0
	startCol := 0

	endRow := 0
	endCol := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		row := []string{}
		for _, char := range line {
			row = append(row, string(char))

			if char == 'S' {
				startRow = len(grid)
				startCol = len(row) - 1
			} else if char == 'E' {
				endRow = len(grid)
				endCol = len(row) - 1
			}
		}

		grid = append(grid, row)
	}

	answer := SolvePart1(grid, startRow, startCol, endRow, endCol, 0, 0)

	fmt.Println("Min value is", answer)

	// PrintGrid(grid)
	fmt.Println("Start row", startRow, "Start col", startCol, "End row", endRow, "End col", endCol)

	data.Close()
}

// right, down, left, up
var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func SolvePart1(grid [][]string, row, col, endRow, endCol, dir int, steps int) int {
	if row == endRow && col == endCol {
		return 0
	}

	key := fmt.Sprintf("%d-%d-%d-%d", row, col, dir, steps)
	if val, ok := memo[key]; ok {
		return val
	}

	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == "#" {
		return math.MaxInt32 // Use a proper "infinity"
	}

	// Mark as visited
	prevVal := grid[row][col]
	grid[row][col] = "#"

	straight := 1 + SolvePart1(grid, row+dirs[dir][0], col+dirs[dir][1], endRow, endCol, dir, steps+1)

	// Turn left
	leftDir := (dir + 3) % 4
	left := 1001 + SolvePart1(grid, row+dirs[leftDir][0], col+dirs[leftDir][1], endRow, endCol, leftDir, steps+1)

	// Turn right
	rightDir := (dir + 1) % 4
	right := 1001 + SolvePart1(grid, row+dirs[rightDir][0], col+dirs[rightDir][1], endRow, endCol, rightDir, steps+1)

	// Restore grid
	grid[row][col] = prevVal

	// Return the minimum cost
	memo[key] = min(straight, min(left, right))
	return memo[key]
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
}
