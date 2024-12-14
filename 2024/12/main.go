package main

import (
	"bufio"
	"container/list"
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

	grid := make([][]string, 0)
	visited := make([][]bool, 0)

	for fileScanner.Scan() {
		var line string = fileScanner.Text()
		row := make([]string, 0)

		for _, char := range line {
			row = append(row, string(char))
		}

		grid = append(grid, row)
		visited = append(visited, make([]bool, len(row)))
	}

	PrintGrid(grid)
	fmt.Println()
	PrintGridBoolean(visited)
	fmt.Println()

	cost := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if visited[row][col] {
				continue
			}

			area, perimeter := FindAreaAndPerimeterPart2(grid, visited, row, col)

			fmt.Println("Area:", area, "Perimeter:", perimeter)

			cost += area * perimeter
		}
	}

	fmt.Println("Cost:", cost)
}

// Somehow you need to count corners instead of fences.
func FindAreaAndPerimeterPart2(grid [][]string, visited [][]bool, startRow int, startCol int) (int, int) {
	label := grid[startRow][startCol]

	queue := list.New()
	queue.PushBack([]int{startRow, startCol})
	visited[startRow][startCol] = true

	area, perimeter := 0, 0
	/*
	   0: Right
	   1: Right-Down
	   2: Down
	   3: Down-Left
	   4: Left
	   5: Left-Up
	   6: Up
	   7: Up-Right
	*/
	directions := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)

		row := front.Value.([]int)[0]
		col := front.Value.([]int)[1]

		area++

		moves := make([]bool, 8)

		for index, direction := range directions {
			newRow := row + direction[0]
			newCol := col + direction[1]

			if !IsInBounds(grid, newRow, newCol) || grid[newRow][newCol] != label {
				moves[index] = true

				continue
			}

			if !visited[newRow][newCol] && index%2 == 0 {
				visited[newRow][newCol] = true
				queue.PushBack([]int{newRow, newCol})
			}
		}

		// If right and down are walls, then it is a corner.
		if moves[0] && moves[2] {
			perimeter++
		}
		// If down and left are walls, then it is a corner.
		if moves[2] && moves[4] {
			perimeter++
		}
		// If left and up are walls, then it is a corner.
		if moves[4] && moves[6] {
			perimeter++
		}
		// If up and right are walls, then it is a corner.
		if moves[6] && moves[0] {
			perimeter++
		}
		// If right and up are not walls but up-right is a wall, then it is a corner.
		if !moves[0] && !moves[6] && moves[7] {
			perimeter++
		}
		// If right and down are not walls but right-down is a wall, then it is a corner.
		if !moves[0] && !moves[2] && moves[1] {
			perimeter++
		}
		// If left and down are not walls but down-left is a wall, then it is a corner.
		if !moves[4] && !moves[2] && moves[3] {
			perimeter++
		}
		// If left and up are not walls but left-up is a wall, then it is a corner.
		if !moves[4] && !moves[6] && moves[5] {
			perimeter++
		}
	}

	return area, perimeter
}

func FindAreaAndPerimeter(grid [][]string, visited [][]bool, startRow int, startCol int) (int, int) {
	label := grid[startRow][startCol]

	queue := list.New()
	queue.PushBack([]int{startRow, startCol})
	visited[startRow][startCol] = true

	area, perimeter := 0, 0
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)

		row := front.Value.([]int)[0]
		col := front.Value.([]int)[1]

		area++

		for _, direction := range directions {
			newRow := row + direction[0]
			newCol := col + direction[1]

			if !IsInBounds(grid, newRow, newCol) {
				perimeter++

				continue
			}

			if grid[newRow][newCol] != label {
				perimeter++

				continue
			}

			if !visited[newRow][newCol] {
				visited[newRow][newCol] = true
				queue.PushBack([]int{newRow, newCol})
			}
		}
	}

	return area, perimeter
}

func IsInBounds(grid [][]string, row int, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Print(char, " ")
		}
		fmt.Println()
	}
}

func PrintGridBoolean(grid [][]bool) {
	for _, row := range grid {
		for _, char := range row {
			if char {
				fmt.Print("T ")
			} else {
				fmt.Print("F ")
			}
		}
		fmt.Println()
	}
}
