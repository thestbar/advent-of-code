package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File open error", err)
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	decoded := []string{}
	id := 0
	isFreespace := false

	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		for _, char := range line {
			val, _ := strconv.Atoi(string(char))

			if isFreespace {
				for i := 0; i < val; i++ {
					decoded = append(decoded, ".")
				}
			} else {
				for i := 0; i < val; i++ {
					decoded = append(decoded, strconv.Itoa(id))
				}

				id++
			}

			isFreespace = !isFreespace
		}
	}

	fmt.Println(decoded)

	decodedCopy := []string{}
	for i := 0; i < len(decoded); i++ {
		decodedCopy = append(decodedCopy, decoded[i])
	}

	// Part 1

	leftIndex, rightIndex := 0, len(decoded)-1

	for leftIndex < rightIndex {
		for decoded[leftIndex] != "." {
			leftIndex++
		}
		for decoded[rightIndex] == "." {
			rightIndex--
		}

		decoded[leftIndex] = decoded[rightIndex]
		decoded[rightIndex] = "."

		leftIndex++
		rightIndex--
	}

	fmt.Println(decoded)

	fmt.Println("Sum:", CalculateSum(decoded, true))

	// Part 2

	// [start, length]
	memory := CalculateFreeSpace(decodedCopy)

	length := 0

	for i := len(decodedCopy) - 1; i >= 0; i-- {
		if decodedCopy[i] == "." {
			continue
		}

		val := decodedCopy[i]

		j := i - 1
		for j = i - 1; j >= 0; j-- {
			if decodedCopy[j] != val {
				break
			}
		}

		if i-j >= 1 {
			length = i - j

			newBlockStart := -1
			newBlockLength := -1

			for _, row := range memory {
				key := row[0]
				value := row[1]

				// fmt.Println("Compare", length, "with", value)
				if length <= value {
					newBlockStart = key
					newBlockLength = value
				}
			}

			// fmt.Println("start:", j + 1, "end:", i, "length:", length, "memoryIndexArray:", memoryIndexArray)

			if newBlockLength != -1 && newBlockStart != -1 && newBlockStart < j+1 {
				// fmt.Println("Move from", j+1, "to", newBlockStart, "length", length)
				for start := j + 1; start <= i; start++ {
					decodedCopy[newBlockStart] = decodedCopy[start]
					decodedCopy[start] = "."

					newBlockStart++
				}
				// fmt.Println(decodedCopy)
				memory = CalculateFreeSpace(decodedCopy)
				// fmt.Println(memory)
				// fmt.Println("----------------")
			}
		}

		i = j + 1
	}

	// fmt.Println(memory)
	// fmt.Println(decodedCopy)

	fmt.Println("Sum:", CalculateSum(decodedCopy, false))
}

func RemoveLastIndex(s []int) []int {
	return s[:len(s)-1]
}

func CalculateFreeSpace(decoded []string) [][]int {
	memory := [][]int{}

	length := 0
	for i := len(decoded) - 1; i >= 0; i-- {
		if decoded[i] == "." {
			length++
		} else {
			if length != 0 {
				memory = append(memory, []int{i + 1, length})
			}

			length = 0
		}
	}

	return memory
}

func CalculateSum(decoded []string, part1 bool) int {
	sum := 0

	id := 0

	for i := 0; i < len(decoded); i++ {
		if decoded[i] != "." {
			val, _ := strconv.Atoi(decoded[i])

			sum += id * val

			if part1 {
				id++
			}
		}

		if !part1 {
			id++
		}
	}

	return sum
}
