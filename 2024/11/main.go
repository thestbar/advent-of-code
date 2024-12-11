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

	stones := make([]int, 0)

	for fileScanner.Scan() {
		var line string = fileScanner.Text()

		// Split stones
		stonesString := strings.Split(line, " ")

		// Convert string to int
		for _, stone := range stonesString {
			stoneNumber, _ := strconv.Atoi(stone)
			stones = append(stones, stoneNumber)
		}
	}

	fmt.Println("Stones: ", stones)

	// Smart blink
	res := 0

	memo := make(map[string]int)

	for _, stone := range stones {
		res += SmartBlink(stone, 0, 75, memo)
	}

	fmt.Println("Smart blink result:", res)
}

func Blink(stones []int) []int {
	newStones := make([]int, 0)

	for _, stone := range stones {
		// * If the stone is engraved with the number 0, it is replaced by a stone
		// engraved with the number 1.
		if stone == 0 {
			newStones = append(newStones, 1)
		} else {
			stoneString := strconv.Itoa(stone)
			stoneNumberLength := len(stoneString)

			// * If the stone is engraved with a number that has an even number of
			// digits, it is replaced by two stones. The left half of the digits are
			// engraved on the new left stone, and the right half of the digits are
			// engraved on the new right stone. (The new numbers don't keep extra leading
			// zeroes: 1000 would become stones 10 and 0.)
			if stoneNumberLength%2 == 0 {
				leftNumber := stoneString[:stoneNumberLength/2]
				rightNumber := stoneString[stoneNumberLength/2:]

				leftNumberInt, _ := strconv.Atoi(leftNumber)
				rightNumberInt, _ := strconv.Atoi(rightNumber)

				newStones = append(newStones, leftNumberInt)
				newStones = append(newStones, rightNumberInt)
			} else {
				// * If none of the other rules apply, the stone is replaced by a new stone;
				// the old stone's number multiplied by 2024 is engraved on the new stone.
				newStone := stone * 2024
				newStones = append(newStones, newStone)
			}
		}
	}
	return newStones
}

func SmartBlink(stone, blink, maxBlinks int, memo map[string]int) int {
	if blink == maxBlinks {
		return 1
	}

	key := strconv.Itoa(stone) + "_" + strconv.Itoa(blink)

	if memo[key] != 0 {
		return memo[key]
	}

	if stone == 0 {
		memo[key] = SmartBlink(1, blink+1, maxBlinks, memo)

		return memo[key]
	}

	stoneString := strconv.Itoa(stone)
	stoneNumberLength := len(stoneString)

	if stoneNumberLength%2 == 0 {
		leftNumber := stoneString[:stoneNumberLength/2]
		rightNumber := stoneString[stoneNumberLength/2:]

		leftNumberInt, _ := strconv.Atoi(leftNumber)
		rightNumberInt, _ := strconv.Atoi(rightNumber)

		memo[key] = SmartBlink(leftNumberInt, blink+1, maxBlinks, memo) + SmartBlink(rightNumberInt, blink+1, maxBlinks, memo)
	} else {
		memo[key] = SmartBlink(stone*2024, blink+1, maxBlinks, memo)
	}

	return memo[key]
}
