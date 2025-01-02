package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)

	secretNumbers := []int{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		secretNumber, _ := strconv.Atoi(line)

		secretNumbers = append(secretNumbers, secretNumber)
	}

	data.Close()

	Part1(secretNumbers)
	Part2(secretNumbers)
}

func Part1(secretNumbers []int) {
	sum := 0
	for _, secret := range secretNumbers {
		for i := 0; i < 2000; i++ {
			secret = CalculateNextSecret(secret)
		}
		sum += secret
	}

	fmt.Println("Sum of all secret numbers after 2000 iterations:", sum)
}

func Part2(secretNumbers []int) {
	pricesBySeller := [][]int{}
	for _, secret := range secretNumbers {
		sellerPrices := []int{}
		for i := 0; i < 2000; i++ {
			// Convert number to string
			secretStr := strconv.Itoa(secret)
			// Get last digit
			lastDigit, _ := strconv.Atoi(string(secretStr[len(secretStr)-1]))
			// Convert it to int
			lastDigitInt := int(lastDigit)
			// Append it to the slice
			sellerPrices = append(sellerPrices, lastDigitInt)
			// Calculate the next secret number
			secret = CalculateNextSecret(secret)
		}
		// Append the slice to the main slice
		pricesBySeller = append(pricesBySeller, sellerPrices)
	}

	// Calculate the associated changes between each sellers prices
	priceChangesBySeller := [][]int{}
	for i := 0; i < len(pricesBySeller); i++ {
		priceChanges := []int{}
		for j := 1; j < len(pricesBySeller[i]); j++ {
			currentPrice := pricesBySeller[i][j]
			previousPrice := pricesBySeller[i][j-1]
			priceChange := currentPrice - previousPrice
			priceChanges = append(priceChanges, priceChange)
		}
		priceChangesBySeller = append(priceChangesBySeller, priceChanges)
	}

	// Iterate each seller and from the 4th element till the last.
	// On each unique pair of 4 consecutive price changes buy from each
	// seller the first price that occurs after the 4 price changes.
	// If a seller does not have these price changes then skip to the next seller.
	// Memoize the sequences of 4 price changes so you do not check for the same.
	maxPrices := 0
	memo := map[string]int{}
	for i := 0; i < len(priceChangesBySeller); i++ {
		for j := 3; j < len(priceChangesBySeller[i]); j++ {
			fmt.Println("Seller:", i, "Index:", j, "Max Prices:", maxPrices, "Memo Length:", len(memo))

			firstPriceChange := priceChangesBySeller[i][j-3]
			secondPriceChange := priceChangesBySeller[i][j-2]
			thirdPriceChange := priceChangesBySeller[i][j-1]
			fourthPriceChange := priceChangesBySeller[i][j]

			buyingPrice := pricesBySeller[i][j+1]
			if buyingPrice*2000 < maxPrices {
				continue
			}

			key := fmt.Sprintf("%d-%d-%d-%d", firstPriceChange, secondPriceChange, thirdPriceChange, fourthPriceChange)
			if _, ok := memo[key]; ok {
				continue
			}
			memo[key] = 1

			// Now start with this sequence of numbers and buy on the price that occurs first
			// from each buyer and return the sum of all the prices.
			prices := BuyOnFirstPriceChange(
				pricesBySeller,
				priceChangesBySeller,
				[4]int{firstPriceChange, secondPriceChange, thirdPriceChange, fourthPriceChange},
			)

			if prices > maxPrices {
				maxPrices = prices
			}
		}
	}

	fmt.Println("Max prices:", maxPrices)
}

func BuyOnFirstPriceChange(prices, priceChanges [][]int, sequence [4]int) int {
	sum := 0

	for i := 0; i < len(prices); i++ {
		for j := 3; j < len(priceChanges[i]); j++ {
			if priceChanges[i][j-3] == sequence[0] &&
				priceChanges[i][j-2] == sequence[1] &&
				priceChanges[i][j-1] == sequence[2] &&
				priceChanges[i][j] == sequence[3] {
				sum += prices[i][j+1]
				break
			}
		}
	}

	return sum
}

// In particular, each buyer's secret number evolves into the next secret number
// in the sequence via the following process:
//
// * Calculate the result of multiplying the secret number by 64. Then, mix this
// result into the secret number. Finally, prune the secret number.
// * Calculate the result of dividing the secret number by 32. Round the result down
// to the nearest integer. Then, mix this result into the secret number. Finally,
// prune the secret number.
// * Calculate the result of multiplying the secret number by 2048. Then, mix this
// result into the secret number. Finally, prune the secret number.
func CalculateNextSecret(secret int) int {
	secret = Mix(secret*64, secret)
	secret = Prune(secret)
	secret = Mix(secret/32, secret)
	secret = Prune(secret)
	secret = Mix(secret*2048, secret)
	secret = Prune(secret)

	return secret
}

// To mix a value into the secret number, calculate the bitwise XOR of the given
// value and the secret number. Then, the secret number becomes the result of that
// operation.
func Mix(val, secret int) int {
	return val ^ secret
}

// To prune the secret number, calculate the value of the secret number modulo
// 16777216.
func Prune(secret int) int {
	return secret % 16777216
}
