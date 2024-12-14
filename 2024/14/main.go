package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	rows int = 103
	cols int = 101
)

type Robot struct {
	x        int
	y        int
	velocity [2]int
}

func (r *Robot) move() {
	r.x += r.velocity[0]
	r.y += r.velocity[1]

	if r.x < 0 {
		r.x = cols + r.x
	}
	if r.y < 0 {
		r.y = rows + r.y
	}
	if r.x >= cols {
		r.x = r.x - cols
	}
	if r.y >= rows {
		r.y = r.y - rows
	}
}

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File open error", err)
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	robots := make([]Robot, 0)
	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		// Generate robots
		splited := strings.Split(line, " ")
		position := splited[0]
		positionArr := strings.Split(position, "=")
		positionArr = strings.Split(positionArr[1], ",")
		x, _ := strconv.Atoi(positionArr[0])
		y, _ := strconv.Atoi(positionArr[1])

		velocity := splited[1]
		velocityArr := strings.Split(velocity, "=")
		velocityArr = strings.Split(velocityArr[1], ",")
		vx, _ := strconv.Atoi(velocityArr[0])
		vy, _ := strconv.Atoi(velocityArr[1])

		robot := Robot{}
		robot.x = x
		robot.y = y
		robot.velocity[0] = vx
		robot.velocity[1] = vy

		robots = append(robots, robot)
	}

	// Generate the grid
	grid := make([][]int, rows)

	for i := 0; i < rows; i++ {
		grid[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = 0
		}
	}

	// Add robots to the grid
	for _, robot := range robots {
		grid[robot.y][robot.x]++
	}

	fmt.Println("Initial state")
	PrintGrid(grid)

	// For part 1 loop 100 seconds
	for i := 0; i >= 0; i++ {
		if rmsDistance(&robots) < 40.0 {
			PrintGrid(grid)
			fmt.Println("Symmetric after", i, "seconds")
			break
		}

		// Move the robots
		for j := 0; j < len(robots); j++ {
			robots[j].move()
		}

		// Clear the grid
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				grid[i][j] = 0
			}
		}

		// Add robots to the grid
		for _, robot := range robots {
			grid[robot.y][robot.x]++
		}
	}

	// fmt.Println("After 100 seconds")
	// PrintGrid(grid)
	//
	// a, b, c, d := SafetyFactors(grid)
	// safetyFactor := a * b * c * d
	//
	// fmt.Println("Safety factor:", safetyFactor)
}

func SafetyFactors(grid [][]int) (int, int, int, int) {
	// Top left
	a := 0
	for i := 0; i < rows/2; i++ {
		for j := 0; j < cols/2; j++ {
			a += grid[i][j]
		}
	}

	// Top right
	b := 0
	for i := 0; i < rows/2; i++ {
		for j := cols/2 + 1; j < cols; j++ {
			b += grid[i][j]
		}
	}

	// Bottom left
	c := 0
	for i := rows/2 + 1; i < rows; i++ {
		for j := 0; j < cols/2; j++ {
			c += grid[i][j]
		}
	}

	// Bottom right
	d := 0
	for i := rows/2 + 1; i < rows; i++ {
		for j := cols/2 + 1; j < cols; j++ {
			d += grid[i][j]
		}
	}

	return a, b, c, d
}

func rmsDistance(robots *[]Robot) float64 {
	distance := 0.0
	for i := 0; i < len(*robots); i++ {
		for j := 0; j < len(*robots); j++ {
			if i == j {
				continue
			}
			robot1 := (*robots)[i]
			robot2 := (*robots)[j]
			distance += float64(robot1.x-robot2.x)*float64(robot1.x-robot2.x) +
				float64(robot1.y-robot2.y)*float64(robot1.y-robot2.y)
		}
	}

	return math.Sqrt(distance / float64(len(*robots)*len(*robots)))
}

func PrintGrid(grid [][]int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(grid[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Pause() {
	fmt.Println("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
