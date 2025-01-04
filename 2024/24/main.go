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
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	inputs := make(map[string]int)
	outputs := []Gate{}
	readingInputs := true

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			readingInputs = false
			continue
		}

		if readingInputs {
			input := strings.Split(line, ": ")
			value, _ := strconv.Atoi(input[1])

			inputs[input[0]] = value
		} else {
			output := strings.Split(line, " -> ")
			gate := Gate{output: output[1]}
			output = strings.Split(output[0], " ")
			gate.input1 = output[0]
			gate.operation = output[1]
			gate.input2 = output[2]

			outputs = append(outputs, gate)
		}
	}

	data.Close()

	Part1(inputs, outputs)
	Part2(inputs, outputs)
}

func Part1(originalInputs map[string]int, outputs []Gate) {
	inputs := make(map[string]int)
	for key, value := range originalInputs {
		inputs[key] = value
	}

	queue := []Gate{}
	for _, gate := range outputs {
		queue = append(queue, gate)
	}

	for len(queue) > 0 {
		gate := queue[0]
		queue = queue[1:]

		if _, ok := inputs[gate.input1]; !ok {
			queue = append(queue, gate)
			continue
		}

		if _, ok := inputs[gate.input2]; !ok {
			queue = append(queue, gate)
			continue
		}

		if _, ok := inputs[gate.output]; ok {
			continue
		}

		switch gate.operation {
		case "AND":
			inputs[gate.output] = And(inputs[gate.input1], inputs[gate.input2])
		case "OR":
			inputs[gate.output] = Or(inputs[gate.input1], inputs[gate.input2])
		case "XOR":
			inputs[gate.output] = Xor(inputs[gate.input1], inputs[gate.input2])
		}
	}

	power := 0
	output := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			byte := fmt.Sprintf("z%d%d", i, j)
			if _, ok := inputs[byte]; ok {
				output += inputs[byte] * (1 << power)
				power++
			}
		}
	}

	fmt.Println("Part1:", output)
}

// Here we need to make sure that the given input correctly implement the Full
// adder between two 45 bit number.
//
// # By using the following formula
//
// Zn = (Xn ⊕ Yn) ⊕ Cn-1
//
// Cn = (Xn * Yn) + (Cn-1 * (Xn ⊕ Yn))
//
// with C0 = (Xn * Yn)
//
// We can derive a series of rule.
//
// AND:
//
// # AND gate can only be input to and OR gate
//
// # AND gate cannot take other AND gate as input
//
// XOR:
//
// XOR gate can only be input to and AND/XOR gate
//
// # XOR gate cannot take AND gate as input
//
// OR:
//
// OR gate can only be input of AND/XOR gate
//
// # OR gate can only take AND gate as input
//
// (Xn ⊕ Yn) ⊕ (a + b) should always output a Zxx except for the last carry z45
//
// A gate with Zxx as its output cannot directly use Xn or Yn as inputs.
//
// Look for gates that does not follow those rules.
func Part2(inputs map[string]int, outputs []Gate) {
	badOutputs := []string{}
	for _, gate := range outputs {
		if gate.output[0] == 'z' && gate.operation != "XOR" && gate.output[1:] != "45" {
			badOutputs = append(badOutputs, gate.output)
		}

		arr := []string{"x", "y", "z"}
		if gate.operation == "XOR" && !Contains(arr, string(gate.output[0])) &&
			!Contains(arr, string(gate.input1[0])) && !Contains(arr, string(gate.input2[0])) {
			badOutputs = append(badOutputs, gate.output)
		}

		arr = []string{gate.input1, gate.input2}
		if gate.operation == "AND" && !Contains(arr, "x00") {
			for _, subGate := range outputs {
				if (gate.output == subGate.input1 || gate.output == subGate.input2) && subGate.operation != "OR" {
					badOutputs = append(badOutputs, gate.output)
				}
			}
		}

		if gate.operation == "XOR" {
			for _, subGate := range outputs {
				if (gate.output == subGate.input1 || gate.output == subGate.input2) && subGate.operation == "OR" {
					badOutputs = append(badOutputs, gate.output)
				}
			}
		}
	}

	// Sort
	sort.Strings(badOutputs)

	// Remove duplicates and print bad outputs
	fmt.Print("Part2: ")
	seen := make(map[string]bool)
	for _, output := range badOutputs {
		if seen[output] {
			continue
		}
		seen[output] = true
		fmt.Print(output)
		if len(seen) < 8 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}

type Gate struct {
	input1    string
	operation string
	input2    string
	output    string
}

// ByCustomOrder implements sort.Interface for []Gate based on multiple fields.
type ByCustomOrder []Gate

func (g ByCustomOrder) Len() int      { return len(g) }
func (g ByCustomOrder) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g ByCustomOrder) Less(i, j int) bool {
	// Compare input1
	if g[i].input1 != g[j].input1 {
		return g[i].input1 < g[j].input1
	}
	// Compare operation
	if g[i].operation != g[j].operation {
		return g[i].operation < g[j].operation
	}
	// Compare input2
	if g[i].input2 != g[j].input2 {
		return g[i].input2 < g[j].input2
	}
	// Compare output
	return g[i].output < g[j].output
}

func And(a, b int) int {
	if a == 1 && b == 1 {
		return 1
	}

	return 0
}

func Or(a, b int) int {
	if a == 1 || b == 1 {
		return 1
	}

	return 0
}

func Xor(a, b int) int {
	if a == 1 && b == 0 {
		return 1
	}

	if a == 0 && b == 1 {
		return 1
	}

	return 0
}
