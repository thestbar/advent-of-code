package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = map[string][]int{
	"^": {-1, 0},
	"v": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	creatingMap := true
	grid := [][]string{}
	grid2 := [][]string{}
	robot := Robot{0, 0, []string{}, &grid}
	robot2 := Robot{0, 0, []string{}, &grid2}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			creatingMap = false
			continue
		}

		if creatingMap {
			row := []string{}
			row2 := []string{}
			for _, char := range line {
				// Create map for part1
				row = append(row, string(char))
				if string(char) == "@" {
					robot.row = len(grid)
					robot.col = len(row) - 1
				}
				// Create map for part2
				if string(char) == "@" {
					row2 = append(row2, "@")
					robot2.row = len(grid2)
					robot2.col = len(row2) - 1
					row2 = append(row2, ".")
				} else if string(char) == "#" {
					row2 = append(row2, "#")
					row2 = append(row2, "#")
				} else if string(char) == "O" {
					row2 = append(row2, "[")
					row2 = append(row2, "]")
				} else {
					row2 = append(row2, ".")
					row2 = append(row2, ".")
				}
			}
			grid = append(grid, row)
			grid2 = append(grid2, row2)
		} else {
			for _, char := range line {
				robot.moves = append(robot.moves, string(char))
				robot2.moves = append(robot2.moves, string(char))
			}
		}
	}

	for len(robot.moves) > 0 {
		robot.Move()
	}

	for len(robot2.moves) > 0 {
		robot2.Move()
	}

	fmt.Println("GPS coordinate sum:", CalculateGPSCoordinateSum(grid))
	fmt.Println("GPS coordinate sum:", CalculateGPSCoordinateSum(grid2))

	data.Close()
}

type Robot struct {
	row   int
	col   int
	moves []string
	grid  *[][]string
}

func (r *Robot) Move() {
	// Check if there are no more moves
	if len(r.moves) == 0 {
		fmt.Println("Robot has no more moves")

		return
	}

	// Get direction
	direction := directions[r.moves[0]]

	// Remove the move from the list
	r.moves = r.moves[1:]

	// Get new position
	newRow := r.row + direction[0]
	newCol := r.col + direction[1]

	// Check if new position is valid
	if newRow < 0 || newRow >= len(*r.grid) || newCol < 0 || newCol >= len((*r.grid)[0]) {
		fmt.Println("Robot tried to get out of the grid")

		return
	}

	// Check if new position is a wall
	if (*r.grid)[newRow][newCol] == "#" {
		fmt.Println("Robot tried to hit a wall")

		return
	}

	// Check if new position is on a box
	if (*r.grid)[newRow][newCol] == "O" {
		fmt.Println("Robot tried to move a box")

		if !MoveBox(*r.grid, newRow, newCol, direction) {
			return
		}
	}

	// Check if a new position is 2 cells box
	if (*r.grid)[newRow][newCol] == "[" || (*r.grid)[newRow][newCol] == "]" {
		fmt.Println("Robot tried to move a 2 cells box")

		// Copy the grid
		gridCopy := CopyGrid(*r.grid)

		if !MoveBox2(*r.grid, newRow, newCol, direction, map[string]int{}) {
			// We have to revert the moves of the boxes
			*r.grid = gridCopy

			return
		}
	}

	// Move robot
	(*r.grid)[r.row][r.col] = "."

	r.row = newRow
	r.col = newCol

	(*r.grid)[r.row][r.col] = "@"
}

func (r *Robot) Print() {
	// Print position
	fmt.Println("Robot position:", r.row, r.col)
	fmt.Println("-----------------")

	// Print grid
	for _, row := range *r.grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println("-----------------")

	// Print moves
	fmt.Println("Robot moves:", r.moves)
	fmt.Println("-----------------")
}

func MoveBox(grid [][]string, row, col int, direction []int) bool {
	newRow := row + direction[0]
	newCol := col + direction[1]

	// Check if new position is valid
	if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
		fmt.Println("Robot tried to move the box out of the grid")

		return false
	}

	// Check if new position is a wall
	if grid[newRow][newCol] == "#" {
		fmt.Println("Robot tried to move the box into a wall")

		return false
	}

	// Check if new position is a box
	if grid[newRow][newCol] == "O" || grid[newRow][newCol] == "[" || grid[newRow][newCol] == "]" {
		fmt.Println("Robot tried to move the box into another box")

		if !MoveBox(grid, newRow, newCol, direction) {
			return false
		}
	}

	// Move box
	grid[row][col] = "."
	grid[newRow][newCol] = "O"

	return true
}

func CalculateGPSCoordinateSum(grid [][]string) int {
	sum := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == "O" || cell == "[" {
				sum += 100*i + j
			}
		}
	}

	return sum
}

func MoveBox2(grid [][]string, row, col int, direction []int, notCheck map[string]int) bool {
	boxCoords := [][]int{}

	// From left to right
	if grid[row][col] == "[" {
		boxCoords = append(boxCoords, []int{row, col})
		boxCoords = append(boxCoords, []int{row, col + 1})
	} else {
		boxCoords = append(boxCoords, []int{row, col - 1})
		boxCoords = append(boxCoords, []int{row, col})
	}

	for _, boxCoord := range boxCoords {
		coordKey := fmt.Sprintf("%d-%d", boxCoord[0], boxCoord[1])
		notCheck[coordKey]++
	}

	newBoxCoords := [][]int{}
	for _, boxCoord := range boxCoords {
		newRow := boxCoord[0] + direction[0]
		newCol := boxCoord[1] + direction[1]

		newBoxCoords = append(newBoxCoords, []int{newRow, newCol})
	}

	// If any of the box coordinates hit a wall, return false
	for _, newBoxCoord := range newBoxCoords {
		if grid[newBoxCoord[0]][newBoxCoord[1]] == "#" {
			fmt.Println("Robot tried to move the box into a wall")

			return false
		}
	}

	// The special case is when we are moving horizontally we don't
	// need to check both cells, we can just check the last one (in the direction)
	// Otherwise we will fall into a loop!
	// We can just exclude cells that do not matter, passing them as input to the function!

	// If any of the box coordinates hit another box, try to move that box
	fmt.Println("New box coords:", newBoxCoords)
	for _, newBoxCoord := range newBoxCoords {
		coordKey := fmt.Sprintf("%d-%d", newBoxCoord[0], newBoxCoord[1])
		// If moving horizontally, we only need to check the last cell
		if notCheck[coordKey] > 0 {
			continue
		}
		if grid[newBoxCoord[0]][newBoxCoord[1]] == "[" || grid[newBoxCoord[0]][newBoxCoord[1]] == "]" {
			fmt.Println("Robot tried to move the box into another box 2")

			if !MoveBox2(grid, newBoxCoord[0], newBoxCoord[1], direction, notCheck) {
				return false
			}
		}
	}

	// Move box
	grid[boxCoords[0][0]][boxCoords[0][1]] = "."
	grid[boxCoords[1][0]][boxCoords[1][1]] = "."

	grid[newBoxCoords[0][0]][newBoxCoords[0][1]] = "["
	grid[newBoxCoords[1][0]][newBoxCoords[1][1]] = "]"

	return true
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func CopyGrid(grid [][]string) [][]string {
	newGrid := [][]string{}
	for _, row := range grid {
		newRow := []string{}
		for _, cell := range row {
			newRow = append(newRow, cell)
		}
		newGrid = append(newGrid, newRow)
	}

	return newGrid
}
