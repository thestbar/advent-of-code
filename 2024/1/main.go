package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File open error", err)
	}

	fileScanner := bufio.NewScanner(data)

	fileScanner.Split(bufio.ScanLines)

	leftItems := []int64{}
	rightItems := []int64{}

	scoreByNumber := make(map[int64]int64)

	for fileScanner.Scan() {
		var line string = fileScanner.Text()
		result := strings.Split(line, " ")

		left := toInt64(result[0])
		right := toInt64(result[3])

		scoreByNumber[right] += 1

		leftItems = append(leftItems, left)
		rightItems = append(rightItems, right)
	}

	var score int64 = 0

	for i := 0; i < len(leftItems); i++ {
		score += scoreByNumber[leftItems[i]] * leftItems[i]
	}

	fmt.Println(score)

	// Sort arrays
	sort.Slice(leftItems, func(i, j int) bool { return leftItems[i] < leftItems[j] })
	sort.Slice(rightItems, func(i, j int) bool { return rightItems[i] < rightItems[j] })

	var sum int64 = 0

	for i := 0; i < len(leftItems); i++ {
		diff := leftItems[i] - rightItems[i]

		if diff < 0 {
			sum += -diff
		} else if diff > 0 {
			sum += diff
		}
	}

	fmt.Println(sum)
}

func toInt64(value string) int64 {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}
