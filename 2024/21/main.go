package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var memo = make(map[string]int)

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	codes := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		codes = append(codes, line)
	}

	data.Close()

	fmt.Println(codes)

	Part1(codes)
	Part2(codes)
}

func Part1(codes []string) {
	numpad := map[rune][]int{
		'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
		'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
		'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
		'0': {3, 1}, 'A': {3, 2},
	}

	dirpad := map[rune][]int{
		'^': {0, 1}, 'A': {0, 2},
		'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
	}

	numpadGraph := CreateGraph(numpad, []int{3, 0})
	dirpadGraph := CreateGraph(dirpad, []int{0, 0})

	totalComplexity := 0
	for _, code := range codes {
		conversion := Convert(code, numpadGraph)
		// 2 for part 1
		conversion = Convert(conversion, dirpadGraph)
		conversion = Convert(conversion, dirpadGraph)

		totalComplexity += CodeToInt(code) * len(conversion)

		fmt.Println(CodeToInt(code), "*", len(conversion))
	}

	fmt.Println(totalComplexity)
}

func CreateGraph(keypad map[rune][]int, invalidCoords []int) map[string]string {
	graph := make(map[string]string)

	for key, value := range keypad {
		x1 := value[0]
		y1 := value[1]

		for key2, value2 := range keypad {
			x2 := value2[0]
			y2 := value2[1]
			path := strings.Repeat("<", int(math.Max(0, float64(y1-y2)))) +
				strings.Repeat("v", int(math.Max(0, float64(x2-x1)))) +
				strings.Repeat("^", int(math.Max(0, float64(x1-x2)))) +
				strings.Repeat(">", int(math.Max(0, float64(y2-y1))))

			if (invalidCoords[0] == x1 && invalidCoords[1] == y2) || (invalidCoords[0] == x2 && invalidCoords[1] == y1) {
				path = ReverseString(path)
			}

			graph[string(key)+string(key2)] = path + "A"
		}
	}

	return graph
}

func Convert(code string, graph map[string]string) string {
	conversion := ""
	prev := 'A'

	for _, c := range code {
		conversion += graph[string(prev)+string(c)]
		prev = c
	}

	return conversion
}

func CodeToInt(code string) int {
	code = code[:len(code)-1]
	i, _ := strconv.Atoi(code)

	return i
}

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Part2(codes []string) {
	numpad := map[rune][]int{
		'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
		'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
		'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
		'0': {3, 1}, 'A': {3, 2},
	}

	dirpad := map[rune][]int{
		'^': {0, 1}, 'A': {0, 2},
		'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
	}

	numpadGraph := CreateGraph(numpad, []int{3, 0})
	dirpadGraph := CreateGraph(dirpad, []int{0, 0})

	totalComplexity := 0
	for _, code := range codes {
		length := GetLength(code, 26, true, numpadGraph, dirpadGraph)

		totalComplexity += CodeToInt(code) * length

		fmt.Println(CodeToInt(code), "*", length)
	}

	fmt.Println(totalComplexity)
}

func GetLength(
	sequence string,
	iterations int,
	firstIteration bool,
	numGraph map[string]string,
	dirGraph map[string]string,
) int {
	if iterations == 0 {
		return len(sequence)
	}

	key := sequence + strconv.Itoa(iterations) + strconv.FormatBool(firstIteration)
	if val, ok := memo[key]; ok {
		return val
	}

	prev := 'A'
	totalLength := 0
	var graph map[string]string

	if firstIteration {
		graph = numGraph
	} else {
		graph = dirGraph
	}

	for _, c := range sequence {
		totalLength += GetLength(graph[string(prev)+string(c)], iterations-1, false, numGraph, dirGraph)
		prev = c
	}

	memo[key] = totalLength
	return totalLength
}
