package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	input := [][][]string{}
	current := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			currentRow := []string{}
			for _, char := range line {
				currentRow = append(currentRow, string(char))
			}
			current = append(current, currentRow)
		} else {
			input = append(input, current)
			current = [][]string{}
		}
	}
	input = append(input, current)

	// Print(input)

	data.Close()

	Part1(input)
}

func Part1(input [][][]string) {
	keys := [][]int{}
	locks := [][]int{}

	for i := 0; i < len(input); i++ {
		if input[i][0][0] == "#" {
			AddItem(&locks, input[i], "#")
		} else {
			AddItem(&keys, input[i], ".")
		}
	}

	// fmt.Println("Keys:", keys)
	// fmt.Println("Locks:", locks)

	FindUniqueKeyLockPairs(keys, locks)
}

// For pair of key and lock, we need the add of each item to be less than 6.
func FindUniqueKeyLockPairs(keys [][]int, locks [][]int) {
	pairs := 0
	for _, key := range keys {
		for _, lock := range locks {
			if IsValid(key, lock) {
				pairs++
			}
		}
	}

	fmt.Println("Valid pairs:", pairs)
}

func IsValid(key []int, lock []int) bool {
	for i := 0; i < len(key); i++ {
		if key[i]+lock[i] >= 6 {
			return false
		}
	}

	return true
}

func AddItem(items *[][]int, item [][]string, specialChar string) {
	nums := []int{}
	for j := 0; j < len(item[0]); j++ {
		nums = append(nums, 0)
	}

	if specialChar == "#" {
		for i := 1; i < len(item); i++ {
			for j := 0; j < len(item[i]); j++ {
				if item[i][j] == "#" {
					nums[j]++
				}
			}
		}
	} else {
		for i := len(item) - 2; i > 0; i-- {
			for j := 0; j < len(item[i]); j++ {
				if item[i][j] == "#" {
					nums[j]++
				}
			}
		}
	}

	*items = append(*items, nums)
}

func Print(input [][][]string) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			for k := 0; k < len(input[i][j]); k++ {
				print(input[i][j][k])
			}
			println()
		}
		println()
	}
}
