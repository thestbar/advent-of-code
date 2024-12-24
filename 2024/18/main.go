package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	gridSize    = 71
	bytesToRead = 1024
	dirs        = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	bytes := []Byte{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		byteCoords := strings.Split(line, ",")

		c, _ := strconv.Atoi(byteCoords[0])
		r, _ := strconv.Atoi(byteCoords[1])

		bytes = append(bytes, Byte{r, c})
	}

	fmt.Println(bytes)

	// Build the grid
	grid := [][]string{}

	for i := 0; i < gridSize; i++ {
		row := []string{}
		for j := 0; j < gridSize; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	// Add the bytes to the grid
	for i, b := range bytes {
		if i == bytesToRead {
			break
		}

		grid[b.row][b.col] = "#"
	}

	PrintGrid(grid)

  part1Ans := Part1(grid)

  fmt.Println("Part 1:", part1Ans)

  part2Ans := Part2(grid, bytes)

  fmt.Println("Part 2:", part2Ans.String())

	data.Close()
}

type Byte struct {
	row int
	col int
}

func PrintGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func Part1(inputGrid [][]string) int {
	// Copy grid to new grid
  // And initialize visited
	grid := [][]string{}
  visited := [][]bool{}
	for i := 0; i < len(inputGrid); i++ {
		row := []string{}
    visitedRow := []bool{}
		for j := 0; j < len(inputGrid[i]); j++ {
			row = append(row, inputGrid[i][j])
      visitedRow = append(visitedRow, false)
		}
		grid = append(grid, row)
    visited = append(visited, visitedRow)
	}

	// BFS from top left corner to bottom right corner
	queue := [][]int{{0, 0, 0}}

	for len(queue) > 0 {
		row, col, steps := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		if row == gridSize-1 && col == gridSize-1 {
      return steps
		}

		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]

			if newRow < 0 || newRow >= gridSize || newCol < 0 || newCol >= gridSize {
				continue
			}

			if grid[newRow][newCol] == "#" {
				continue
			}

      if visited[newRow][newCol] {
        continue
      }

      visited[newRow][newCol] = true
			queue = append(queue, []int{newRow, newCol, steps + 1})
		}
	}

  return -1
}

func Part2(grid [][]string, bytes []Byte) Byte {
  for i := bytesToRead; i < len(bytes); i++ {
    grid[bytes[i].row][bytes[i].col] = "#"

    ans := Part1(grid)

    if ans == -1 {
      return bytes[i]
    }
  }

  return Byte{-1, -1}
}

func (b *Byte) String() string {
  return fmt.Sprintf("%d, %d", b.col, b.row)
}
