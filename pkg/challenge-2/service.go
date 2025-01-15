package challenge2

import (
	"math"
	"strconv"
)

// Function to decode and find the minimum sum sequence
func ProcessKeyboardDecode(encoded string) responseKeyBoardDecoded {
	n := len(encoded) + 1
	minSum := math.MaxInt
	var result []int

	// Start backtracking
	result = backtrack(encoded, n, []int{}, &minSum)

	// Convert result sequence to string and calculate the sum
	decoded := joinSequence(result)
	sum := sumSequence(result)

	return responseKeyBoardDecoded{Decoded: decoded, Sum: sum}
}

// Backtracking function
func backtrack(encoded string, n int, sequence []int, minSum *int) []int {
	if len(sequence) == n {
		// Check if the sequence matches the encoded pattern
		if isValid(sequence, encoded) {
			sum := sumSequence(sequence)
			if sum < *minSum {
				*minSum = sum
				return append([]int{}, sequence...)
			}
		}
		return nil
	}

	var bestResult []int
	// Try every digit from 0 to 9
	for i := 0; i <= 9; i++ {
		if len(sequence) == 0 || isValidPair(sequence[len(sequence)-1], i, encoded[len(sequence)-1]) {
			nextResult := backtrack(encoded, n, append(sequence, i), minSum)
			if nextResult != nil {
				bestResult = nextResult
			}
		}
	}

	return bestResult
}

// Check if the current sequence matches the encoded pattern
func isValid(sequence []int, encoded string) bool {
	for i := 0; i < len(encoded); i++ {
		if !isValidPair(sequence[i], sequence[i+1], encoded[i]) {
			return false
		}
	}
	return true
}

// Check if a pair of numbers matches the given condition
func isValidPair(left, right int, condition byte) bool {
	switch condition {
	case 'L':
		return left > right
	case 'R':
		return left < right
	case '=':
		return left == right
	}
	return false
}

// Calculate the sum of a sequence
func sumSequence(sequence []int) int {
	sum := 0
	for _, num := range sequence {
		sum += num
	}
	return sum
}

// Join a sequence of numbers into a string
func joinSequence(sequence []int) string {
	result := ""
	for _, num := range sequence {
		result += strconv.Itoa(num)
	}
	return result
}
