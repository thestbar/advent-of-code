package main

import (
	"bufio"
	"fmt"
	"os"
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

	isInput := false

	rules := make(map[int][]int)
	validLines := []string{}
	invalidLines := []string{}

	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		if len(line) == 0 {
			isInput = true

			continue
		}

		if isInput {
			fmt.Println("--------------------")
			fmt.Println("Started checking for:", line)

			numbers := []int{}
			isValid := true

			rowNumbers := strings.Split(line, ",")
			for _, c := range rowNumbers {
				n, _ := strconv.Atoi(string(c))
				invalid := rules[n]

				fmt.Println("Comparing - Invalid:", invalid, "and Numbers:", numbers)
				for _, in := range invalid {
					for _, nn := range numbers {
						if nn == in {
							fmt.Println("Invalid", nn, in)
							isValid = false
							break
						}
					}
					if !isValid {
						break
					}
				}
				numbers = append(numbers, n)

				if !isValid {
					break
				}
			}

			if isValid {
				fmt.Println("VALID", line)
				validLines = append(validLines, line)
			} else {
				invalidLines = append(invalidLines, line)
			}
		} else {
			values := strings.Split(line, "|")

			left, _ := strconv.Atoi(values[0])
			right, _ := strconv.Atoi(values[1])

			rules[left] = append(rules[left], right)
		}
	}

	fmt.Println("--------------------")
	fmt.Println("Rules:", rules)
	fmt.Println("--------------------")
	fmt.Println("Valid Lines:", validLines)
	fmt.Println("--------------------")

	sum := 0
	for _, v := range validLines {
		arr := strings.Split(v, ",")

		index := len(arr) / 2

		n, _ := strconv.Atoi(arr[index])

		fmt.Println("Adding:", n, "from:", arr, "at index:", index)

		sum += n
	}

	fmt.Println("Sum:", sum)
	fmt.Println("--------------------")
	fmt.Println("Invalid Lines:", invalidLines)

	sum = 0
	for _, v := range invalidLines {
		arr := strings.Split(v, ",")
		intArr := []int{}
		for _, a := range arr {
			n, _ := strconv.Atoi(a)
			intArr = append(intArr, n)
		}

		for i := 0; i < len(intArr); i++ {
			rule := rules[intArr[i]]

			for j := 0; j < i; j++ {
				for _, r := range rule {
					if intArr[j] == r {
						tmp := intArr[j]
						intArr[j] = intArr[i]
						intArr[i] = tmp
						i = 0
						break
					}
				}
			}
		}

		fmt.Println("Corrected:", arr, "to:", intArr)

		index := len(arr) / 2

		sum += intArr[index]
	}

	fmt.Println("Sum:", sum)
}
