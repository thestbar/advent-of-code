package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("File open error", err)
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	grid := [][]int{}

	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		row := []int{}

		for _, char := range line {
			val, _ := strconv.Atoi(string(char))

			row = append(row, val)
		}

		grid = append(grid, row)
	}

	PrintGrid(grid)
	fmt.Println("--------------------")

	sum := 0
	ratingSum := 0

	memo := GenerateMemo(len(grid), len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				CalculateTrailheads(grid, i, j, &sum)
				ratingSum += CalculateTrailheadsRating(grid, memo, i, j)
			}
		}
	}

	fmt.Println("Sum of trailheads:", sum)
	fmt.Println("Rating sum of trailheads:", ratingSum)
}

// BFS to calculate trailheads
func CalculateTrailheads(grid [][]int, row, col int, sum *int) {
	queue := make([][]int, 0)
	visited := make([][]int, 0)

	for i := 0; i < len(grid); i++ {
		row := []int{}
		for j := 0; j < len(grid[0]); j++ {
			row = append(row, 0)
		}
		visited = append(visited, row)
	}

	queue = append(queue, []int{row, col})
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		row := queue[0][0]
		col := queue[0][1]
		queue = queue[1:]

		if grid[row][col] == 9 {
			if visited[row][col] == 0 {
				*sum++
			}

			visited[row][col]++

			continue
		}

		for _, dir := range dirs {
			newRow := row + dir[0]
			newCol := col + dir[1]

			if newRow >= 0 && newRow < len(grid) &&
				newCol >= 0 && newCol < len(grid[0]) &&
				grid[newRow][newCol]-grid[row][col] == 1 {
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}
}

// DFS to calculate trailheads rating
func CalculateTrailheadsRating(grid, memo [][]int, row, col int) int {
	if grid[row][col] == 9 {
		return 1
	}

	if memo[row][col] != -1 {
		return memo[row][col]
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	value := 0

	for _, dir := range dirs {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newCol >= 0 && newRow < len(grid) && newCol < len(grid[0]) &&
			grid[newRow][newCol]-grid[row][col] == 1 {
			value += CalculateTrailheadsRating(grid, memo, newRow, newCol)
		}
	}

	memo[row][col] = value

	return value
}

func GenerateMemo(m, n int) [][]int {
	memo := [][]int{}

	for i := 0; i < m; i++ {
		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, -1)
		}
		memo = append(memo, row)
	}

	return memo
}

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}
