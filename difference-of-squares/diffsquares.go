// Package diffsquares finds the difference between the square of the sum
// and the sum of the squares of the first N natural numbers.
package diffsquares

const testVersion = 1

// SquareOfSums adds all num and returns its square
func SquareOfSums(num int) int {
	result := 0
	for i := 0; i <= num; i++ {
		result += i
	}
	return result * result
}

// SumOfSquares takes all num to the square and
// returns its sum
func SumOfSquares(num int) int {
	result := 0
	for i := 0; i <= num; i++ {
		result += i * i
	}
	return result
}

// Difference returns the difference of both sums
func Difference(num int) int {
	return SquareOfSums(num) - SumOfSquares(num)
}
