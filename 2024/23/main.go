package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	connections := make(map[string]map[string]bool)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		connectedComputers := strings.Split(line, "-")

		if _, ok := connections[connectedComputers[0]]; !ok {
			connections[connectedComputers[0]] = map[string]bool{connectedComputers[1]: true}
		} else {
			connections[connectedComputers[0]][connectedComputers[1]] = true
		}

		if _, ok := connections[connectedComputers[1]]; !ok {
			connections[connectedComputers[1]] = map[string]bool{connectedComputers[0]: true}
		} else {
			connections[connectedComputers[1]][connectedComputers[0]] = true
		}
	}

	data.Close()

	Part1(connections)
	Part2(connections)
}

// Looking for sets of three computers where each computer in the set is connected
// to the other two computers.
func Part1(connections map[string]map[string]bool) {
	memory := make(map[string]bool)
	sum := 0
	for computer1, computer1Connections := range connections {
		memory[computer1] = true
		for computer2 := range computer1Connections {
			for computer3 := range computer1Connections {
				if computer1 == computer2 || computer1 == computer3 || computer2 == computer3 {
					continue
				}
				list := []string{computer1, computer2, computer3}
				sort.Strings(list)
				key := strings.Join(list, " ")
				if _, ok := memory[key]; ok {
					continue
				}
				memory[key] = true

				if _, ok := connections[computer3]; ok {
					computer3Connections := connections[computer3]

					if computer3Connections[computer2] {
						if computer1[0] == 't' || computer2[0] == 't' || computer3[0] == 't' {
							sum++
						}
					}
				}
			}
		}
	}

	fmt.Println("Sets of three inter-connected computers:", sum)
}

// Looking for the largest set of computers that are all connected to each other.
func Part2(connections map[string]map[string]bool) {
	// Get all distinct computers
	bestPassword := ""
	memo := make(map[string]string)
	for computer := range connections {
		password := Traverse(computer, computer, 0, computer, memo, connections)
		if len(password) > len(bestPassword) {
			bestPassword = password
		}
	}

	fmt.Println("The correct password is:", bestPassword)
}

func Traverse(
	initialComputer string,
	currentComputer string,
	numOfConnections int,
	password string,
	memo map[string]string,
	connections map[string]map[string]bool,
) string {
	// You are at the end
	if currentComputer == initialComputer && numOfConnections > 0 {
		return strings.Join(strings.Split(password, ",")[:numOfConnections], ",")
	}

	// Check if already visited
	key := password + currentComputer
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	bestPassword := ""

	// Get connections of current computer
	computerConnections := connections[currentComputer]

	lanComputers := strings.Split(password, ",")

	// Traverse through all connections only if all lan computers
	// are connected with the next computer or if it is the initial computer.
	for nextComputer := range computerConnections {
		areAllConnected := true
		for _, lanComputer := range lanComputers {
			if _, ok := connections[nextComputer][lanComputer]; !ok {
				areAllConnected = false
				break
			}
		}

		if areAllConnected || nextComputer == initialComputer {
			lanComputers := strings.Split(password, ",")
			sort.Strings(lanComputers)
			nextPassword := strings.Join(lanComputers, ",")

			nextPassword = Traverse(
				initialComputer,
				nextComputer,
				numOfConnections+1,
				nextPassword+","+nextComputer,
				memo,
				connections,
			)

			if len(nextPassword) > len(bestPassword) {
				bestPassword = nextPassword
			}
		}
	}

	memo[key] = bestPassword
	return bestPassword
}
