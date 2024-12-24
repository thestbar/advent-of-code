package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	grid := [][]string{}

	start := []int{0, 0}
	end := []int{0, 0}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		gridRow := []string{}
		for _, char := range line {
			gridRow = append(gridRow, string(char))

			if char == 'S' {
				start = []int{len(grid), len(gridRow) - 1}
			}

			if char == 'E' {
				end = []int{len(grid), len(gridRow) - 1}
			}
		}

		grid = append(grid, gridRow)
	}

	data.Close()

	PrintGrid(grid)

	fmt.Println("Start:", start)
	fmt.Println("End:", end)

	Part1(grid, start, end)
	Part2(grid, start, end)
}

func Part1(gridInput [][]string, start []int, end []int) {
	// Copy grid
	grid := [][]string{}
	visited := [][]bool{}
	for _, row := range gridInput {
		gridRow := []string{}
		visitedRow := []bool{}
		for _, char := range row {
			gridRow = append(gridRow, char)
			visitedRow = append(visitedRow, false)
		}
		grid = append(grid, gridRow)
		visited = append(visited, visitedRow)
	}

	// First find the exit without cheating
	answers := []int{}
	FindExit(grid, start, end, visited, 0, 0, &answers)

	// Now there will be only one answer which is the minimum time it takes
	// to get to the exit without cheating
	withoutCheating := answers[0]
	fmt.Println("Minimum time without cheating:", withoutCheating)

	// Now find the exit with cheating
	answers = []int{}
	FindExit(grid, start, end, visited, 1, 0, &answers)

	// Count them
	totalSaves := make(map[int]int)
	for _, answer := range answers {
		totalSaves[withoutCheating-answer]++
	}

	fmt.Println("Total saves:", totalSaves)

	// Best cheats are the ones that saves more or equal than 100 ps
	bestCheats := 0
	for k, v := range totalSaves {
		if k >= 100 {
			bestCheats += v
		}
	}

	fmt.Println("Best cheats:", bestCheats)
}

func FindExit(grid [][]string, pos []int, end []int, visited [][]bool, cheatsLeft int, seconds int, answers *[]int) {
	if pos[0] == end[0] && pos[1] == end[1] {
		*answers = append(*answers, seconds)

		// fmt.Println("Found exit in", seconds, "seconds")
		return
	}

	if visited[pos[0]][pos[1]] {
		return
	}

	visited[pos[0]][pos[1]] = true

	for _, dir := range dirs {
		newPos := []int{pos[0] + dir[0], pos[1] + dir[1]}
		// If out of bounds just continue
		if newPos[0] < 0 || newPos[0] >= len(grid) || newPos[1] < 0 || newPos[1] >= len(grid[0]) {
			continue
		}

		// Detect already visited
		if visited[newPos[0]][newPos[1]] {
			continue
		}

		// Detect wall hit
		if grid[newPos[0]][newPos[1]] == "#" {
			// If you have cheats left, use one!
			if cheatsLeft > 0 {
				// First you check the next position
				FindExit(grid, newPos, end, visited, 0, seconds+1, answers)
			}

			continue
		}

		FindExit(grid, newPos, end, visited, cheatsLeft, seconds+1, answers)
	}

	visited[pos[0]][pos[1]] = false
}

func Part2(grid [][]string, start, end []int) {
	rows := len(grid)
	cols := len(grid[0])

	dists := [][]int{}
	for i := 0; i < rows; i++ {
		distsRow := []int{}
		for j := 0; j < cols; j++ {
			distsRow = append(distsRow, -1)
		}
		dists = append(dists, distsRow)
	}

	// BFS
	queue := [][]int{start}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			newPos := []int{pos[0] + dir[0], pos[1] + dir[1]}
			if newPos[0] < 0 || newPos[0] >= rows || newPos[1] < 0 || newPos[1] >= cols {
				continue
			}

			if dists[newPos[0]][newPos[1]] != -1 {
				continue
			}

			if grid[newPos[0]][newPos[1]] == "#" {
				continue
			}

			dists[newPos[0]][newPos[1]] = dists[pos[0]][pos[1]] + 1
			queue = append(queue, newPos)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(dists[i][j], " ")
		}
		fmt.Println()
	}

	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == "#" {
				continue
			}

			for radius := 2; radius < 2; radius++ {
				for dr := 0; dr < radius+1; dr++ {
					dc := radius - dr
					cheatDists := [][]int{{dr, dc}, {dr, -dc}, {-dr, dc}, {-dr, -dc}}
					for _, cheatDist := range cheatDists {
						newPos := []int{i + cheatDist[0], j + cheatDist[1]}
						if newPos[0] < 0 || newPos[0] >= rows || newPos[1] < 0 || newPos[1] >= cols {
							continue
						}

						if grid[newPos[0]][newPos[1]] == "#" {
							continue
						}

						if dists[i][j]-dists[newPos[0]][newPos[1]] >= 100+radius {
							count++
						}
					}
				}
			}
		}
	}

	fmt.Println("Total saves with cheats:", count)
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
}
