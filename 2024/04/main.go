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

	m := 0
	n := 0

	arr := []rune{}

	for fileScanner.Scan() {
		var line string = fileScanner.Text()
		m++

		for _, c := range line {
			if m == 1 {
				n++
			}
			arr = append(arr, c)
		}
	}

	// Letter codes are:
	// X = 88
	// M = 77
	// A = 65
	// S = 83
	solvePart1(arr, m, n)
	solvePart2(arr, m, n)
}

func solvePart1(arr []rune, m int, n int) {
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Check if there is a X
			if getItemFrom2DArray(arr, i, j, m, n) == 88 {

				// Check if there is a M to the right
				if getItemFrom2DArray(arr, i+1, j, m, n) == 77 {
					// Check if there is a A to the right
					if getItemFrom2DArray(arr, i+2, j, m, n) == 65 {
						// Check if there is a S to the right
						if getItemFrom2DArray(arr, i+3, j, m, n) == 83 {
							count++
						}
					}
				}

				// Check if there is a M to the bottom right
				if getItemFrom2DArray(arr, i+1, j+1, m, n) == 77 {
					// Check if there is a A to the bottom right
					if getItemFrom2DArray(arr, i+2, j+2, m, n) == 65 {
						// Check if there is a S to the bottom right
						if getItemFrom2DArray(arr, i+3, j+3, m, n) == 83 {
							count++
						}
					}
				}

				// Check if there is a M to the bottom
				if getItemFrom2DArray(arr, i, j+1, m, n) == 77 {
					// Check if there is a A to the bottom
					if getItemFrom2DArray(arr, i, j+2, m, n) == 65 {
						// Check if there is a S to the bottom
						if getItemFrom2DArray(arr, i, j+3, m, n) == 83 {
							count++
						}
					}
				}

				// Check if there is a M to the bottom left
				if getItemFrom2DArray(arr, i-1, j+1, m, n) == 77 {
					// Check if there is a A to the bottom left
					if getItemFrom2DArray(arr, i-2, j+2, m, n) == 65 {
						// Check if there is a S to the bottom left
						if getItemFrom2DArray(arr, i-3, j+3, m, n) == 83 {
							count++
						}
					}
				}

				// Check if there is a M to the left
				if getItemFrom2DArray(arr, i-1, j, m, n) == 77 {
					// Check if there is a A to the left
					if getItemFrom2DArray(arr, i-2, j, m, n) == 65 {
						// Check if there is a S to the left
						if getItemFrom2DArray(arr, i-3, j, m, n) == 83 {
							count++
						}
					}
				}

				// Check if there is a M to the top left
				if getItemFrom2DArray(arr, i-1, j-1, m, n) == 77 {
					// Check if there is a A to the top left
					if getItemFrom2DArray(arr, i-2, j-2, m, n) == 65 {
						// Check if there is a S to the top left
						if getItemFrom2DArray(arr, i-3, j-3, m, n) == 83 {
							count++
						}
					}
				}

				// Check if there is a M to the top
				if getItemFrom2DArray(arr, i, j-1, m, n) == 77 {
					// Check if there is a A to the top
					if getItemFrom2DArray(arr, i, j-2, m, n) == 65 {
						// Check if there is a S to the top
						if getItemFrom2DArray(arr, i, j-3, m, n) == 83 {
							count++
						}
					}
				}

				// Check if there is a M to the top right
				if getItemFrom2DArray(arr, i+1, j-1, m, n) == 77 {
					// Check if there is a A to the top right
					if getItemFrom2DArray(arr, i+2, j-2, m, n) == 65 {
						// Check if there is a S to the top right
						if getItemFrom2DArray(arr, i+3, j-3, m, n) == 83 {
							count++
						}
					}
				}
			}
		}
	}

	fmt.Println("Part 1 answer:", count)
}

func solvePart2(arr []rune, m int, n int) {
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Check if there is an A
			if getItemFrom2DArray(arr, i, j, m, n) == 65 {

				// Check if top left is M and bottom right is S
				if getItemFrom2DArray(arr, i-1, j-1, m, n) == 77 &&
					getItemFrom2DArray(arr, i+1, j+1, m, n) == 83 {

					// Check if top right is M and bottom left is S
					if getItemFrom2DArray(arr, i-1, j+1, m, n) == 77 &&
						getItemFrom2DArray(arr, i+1, j-1, m, n) == 83 {
						count++
					}

					// Check if top right is S and bottom left is M
					if getItemFrom2DArray(arr, i-1, j+1, m, n) == 83 &&
						getItemFrom2DArray(arr, i+1, j-1, m, n) == 77 {
						count++
					}
				}

				// Check if top left is S and bottom right is M
				if getItemFrom2DArray(arr, i-1, j-1, m, n) == 83 &&
					getItemFrom2DArray(arr, i+1, j+1, m, n) == 77 {

					// Check if top right is M and bottom left is S
					if getItemFrom2DArray(arr, i-1, j+1, m, n) == 77 &&
						getItemFrom2DArray(arr, i+1, j-1, m, n) == 83 {
						count++
					}

					// Check if top right is S and bottom left is M
					if getItemFrom2DArray(arr, i-1, j+1, m, n) == 83 &&
						getItemFrom2DArray(arr, i+1, j-1, m, n) == 77 {
						count++
					}
				}
			}
		}
	}

	fmt.Println("Part 2 answer:", count)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting to int", err)
	}

	return i
}

func getItemFrom2DArray(arr []rune, i int, j int, m int, n int) rune {
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}

	return arr[i*n+j]
}
