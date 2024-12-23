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
	memo               = map[string]int{}
	output             = []int{}
	registerA          = 0
	registerB          = 0
	registerC          = 0
	instructionPointer = 0
	ans                = []int{}
)

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	instructions := []int{}

	calculatingRegisters := true
	counter := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			continue
		}

		values := strings.Split(line, ": ")
		if calculatingRegisters {

			if counter == 0 {
				registerA, _ = strconv.Atoi(values[1])
			} else if counter == 1 {
				registerB, _ = strconv.Atoi(values[1])
			} else if counter == 2 {
				registerC, _ = strconv.Atoi(values[1])

				calculatingRegisters = false
			}

			counter++
		} else {
			instructionsStr := strings.Split(values[1], ",")

			for _, instructionStr := range instructionsStr {
				instruction, _ := strconv.Atoi(instructionStr)
				instructions = append(instructions, instruction)
			}
		}
	}

	data.Close()

	// initialRegisterA := registerA
	initialRegisterB := registerB
	initialRegisterC := registerC
	// currentRegisterA := 0

	/* Part 2 thoughts
	We understand (by converting the register value and observing the output) that the program
	adds a extra value to the output for every 3 bits of the register A. This means that the
	output that we are trying to match has 16 decimal values, which means that the register A
	has 48 bits.
	The range of the numbers we need to check are from
	100000000000000000000000000000000000000000000000 to
	111111111111111111111111111111111111111111111111

	The decimal range is:
	140737488355328 to
	281474976710655

	Register A: 24847151
	Register B: 0
	Register C: 0

	Program: 2,4, 1,5, 7,5, 1,6, 0,3, 4,0, 5,5, 3,0
	The program is a big loop (because of the 3,0 at the end)

	b = a % 8
	b = b ^ 5
	c = a / 2**b
	b = b ^ 6
	a = a / 8
	b = b ^ c
	out(b % 8)
	if a != 0 jump to start

	Here we understand for example that:
	On the last iteration the a needs to be 3 because only if it is 3 then the output will be 5.
	We understood this because we know that it should be 0 after a / 8, therefore there were only
	7 possible values for it (1,2,3,4,5,6,7 - it could not already be 0).

	By checking each one of them we come to the result that the register A should be 3 after dividing
	it by 8.

	So, on the previous iteration it should be from 24 to 31. We can check each one of them and see
	which one is the correct one.

	...
	*/

	find(instructions, len(instructions)-1, 0, initialRegisterB, initialRegisterC)

	fmt.Println("ans:", ans)

	// Find min
	if len(ans) == 0 {
		return
	}

	min := ans[0]
	for _, v := range ans {
		if v < min {
			min = v
		}
	}

	fmt.Println("Min:", min)
}

// Divide A by 2^combo and store the result in A register
func adv(combo int) {
	denominator := int(math.Pow(2, float64(getCombo(combo))))

	registerA = registerA / denominator
}

// Bitwise XOR between register B and literal and store the result in B register
func bxl(literal int) {
	registerB = registerB ^ literal
}

// combo modulo 8 and store the result in B register
func bst(combo int) {
	registerB = getCombo(combo) % 8
}

// If A == 0 just move, else jump to instruction number literal
func jnz(literal int) {
	if registerA != 0 {
		instructionPointer = literal
	} else {
		instructionPointer += 2
	}
}

// Bitwise XOR between values of registers B and C, store in B
func bxc() {
	registerB = registerB ^ registerC
}

// Output combo module 8
func out(combo int) {
	output = append(output, getCombo(combo)%8)

	// fmt.Println("Output:", output)
}

// Divide A by 2^combo and store the result in B register
func bdv(combo int) {
	denominator := int(math.Pow(2, float64(getCombo(combo))))

	registerB = registerA / denominator
}

// Divide A by 2^combo and store the result in C register
func cdv(combo int) {
	denominator := int(math.Pow(2, float64(getCombo(combo))))

	registerC = registerA / denominator
}

func getCombo(combo int) int {
	if combo >= 0 && combo <= 3 {
		return combo
	}

	if combo == 4 {
		return registerA
	}

	if combo == 5 {
		return registerB
	}

	if combo == 6 {
		return registerC
	}

	return -1
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// Prints a decimal number in binary
func printBinary(n int) string {
	return fmt.Sprintf("%b", n)
}

// Converts a binary number given as a string to int
func binaryToInt(binary string) int {
	n, _ := strconv.ParseInt(binary, 2, 64)

	return int(n)
}

func revertBinary(binary string) string {
	runes := []rune(binary)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func addOneToBinary(binary string) string {
	runes := []rune(binary)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '0' {
			runes[i] = '1'

			break
		}

		runes[i] = '0'
	}

	return string(runes)
}

func runProgram(instructions []int, a, b, c int) {
	instructionPointer = 0

	registerA = a
	registerB = b
	registerC = c

	output = []int{}

	for instructionPointer < len(instructions) && instructionPointer+1 < len(instructions) {
		opcode := instructions[instructionPointer]
		operand := instructions[instructionPointer+1]

		switch opcode {
		case 0:
			adv(operand)
		case 1:
			bxl(operand)
		case 2:
			bst(operand)
		case 3:
			jnz(operand)
		case 4:
			bxc()
		case 5:
			out(operand)
		case 6:
			bdv(operand)
		case 7:
			cdv(operand)
		}

		if opcode != 3 {
			instructionPointer += 2
		}
	}
}

func find(instructions []int, i, a, b, c int) {
	for j := 0; j < 8; j++ {
		registerA = a + j
		out := smartRunProgram(registerA, b, c)

		if equal(out, instructions[i:]) {
			if i == 0 {
				fmt.Println("out:", out)
				ans = append(ans, registerA)
			} else {
				find(instructions, i-1, registerA*8, b, c)
			}
		}
	}
}

func smartRunProgram(a, b, c int) []int {
	out := []int{}

	for a != 0 {
		b = a % 8
		b = b ^ 5
		c = a / int(math.Pow(2, float64(b)))
		b = b ^ 6
		a = a / 8
		b = b ^ c
		out = append(out, b%8)
	}

	return out
}
