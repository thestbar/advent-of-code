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

	grid := [][]rune{}
	guardRow := 0
	guardCol := 0

	// Characters ASCII values
	// ^ = 94
	// v = 118
	// < = 60
	// > = 62
	// . = 46
	// # = 35
	// X = 88
	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		row := []rune{}

		for _, c := range line {
			row = append(row, c)
			if c == 94 {
				guardRow = len(grid)
				guardCol = len(row) - 1
			}
		}

		grid = append(grid, row)
	}

	guardStartRow := guardRow
	guardStartCol := guardCol

	moveGuard(grid, guardRow, guardCol)

	positions := 0

	// Calculate the number of positions visited
	for _, row := range grid {
		for _, cell := range row {
			if cell == 88 {
				positions++
			}
		}
	}

	fmt.Println("Positions visited:", positions)
	fmt.Println("--------------------")

	// Now the grid has all the positions that the guard has visited.
	// We need to find a clever way to calculate adding obstacles, which
	// could lead to the guard get in a loop.

	// First create a copy of the grid and replace all X with . and put
	// the guard in the starting position.
	gridCopy := [][]rune{}
	for i := 0; i < len(grid); i++ {
		gridCopy = append(gridCopy, []rune{})
		for j := 0; j < len(grid[0]); j++ {
			gridCopy[i] = append(gridCopy[i], grid[i][j])
			if gridCopy[i][j] == 88 {
				gridCopy[i][j] = 46
			}
		}
	}

	gridCopy[guardStartRow][guardStartCol] = 94

	// printGrid(gridCopy)
	// fmt.Println("--------------------")

	// Now try adding a new obstacle on each position that the guard has visited
	// and see if the guard can escape.
	loopCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 88 || (i == guardStartRow && j == guardStartCol) {
				continue
			}

			// Create a new grid from the copy and add the obstacle
			currentGrid := [][]rune{}
			for i := 0; i < len(gridCopy); i++ {
				currentGrid = append(currentGrid, []rune{})
				for j := 0; j < len(gridCopy[0]); j++ {
					currentGrid[i] = append(currentGrid[i], gridCopy[i][j])
				}
			}
			currentGrid[i][j] = 35

			// fmt.Println("Obstacle at", i, j)
			// printGrid(currentGrid)

			// Now try to move the guard
			if !moveGuard(currentGrid, guardStartRow, guardStartCol) {
				loopCount++
			}
		}
	}

	fmt.Println("Loops:", loopCount)
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func getDelta(direction rune) (int, int) {
	switch direction {
	case 94:
		return -1, 0
	case 118:
		return 1, 0
	case 60:
		return 0, -1
	case 62:
		return 0, 1
	default:
		return 0, 0
	}
}

func changeDirection(grid [][]rune, row, col int) {
	if grid[row][col] == 94 {
		grid[row][col] = 62
	} else if grid[row][col] == 62 {
		grid[row][col] = 118
	} else if grid[row][col] == 118 {
		grid[row][col] = 60
	} else if grid[row][col] == 60 {
		grid[row][col] = 94
	}
}

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Press any button to continue")

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) >= 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

// The guard while moving needs to understand if she has been in a loop.
// In that case this function should return false and the guard should stop.
// Otherwise it returns true and the function ends when the guard is out of the grid.
//
// How to detect a loop?
// We can just count the number of steps the guard has taken and if it exceeds the
// number of cells in the grid, then we can assume that the guard is in a loop.
func moveGuard(grid [][]rune, guardRow, guardCol int) bool {
	isAlive := true

	gridSize := len(grid) * len(grid[0])
	steps := 0

	for isAlive {
		// readInput()
		// printGrid(grid)
		// fmt.Println("--------------------")
		// fmt.Println("Guard at", guardRow, guardCol)

		// Move guard
		direction := grid[guardRow][guardCol]

		deltaRow, deltaCol := getDelta(direction)

		// fmt.Println("Delta row:", deltaRow, "Delta col:", deltaCol)

		foundObstacle := false

		for !foundObstacle {
			if steps > gridSize {
				return false
			}

			newRow := guardRow + deltaRow
			newCol := guardCol + deltaCol

			if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
				isAlive = false
				grid[guardRow][guardCol] = 88

				break
			}

			if grid[newRow][newCol] == 35 {
				foundObstacle = true
				changeDirection(grid, guardRow, guardCol)
			} else {
				steps++
				grid[guardRow][guardCol] = 88
				guardRow = newRow
				guardCol = newCol
				grid[guardRow][guardCol] = direction
			}
		}
	}
	// printGrid(grid)

	return true
}
