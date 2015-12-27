// Sum square difference
// https://projecteuler.net/problem=6
//
// The sum of the squares of the first ten natural numbers is,
//
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
package main

import (
	"fmt"
	"math"
)

func square(num int) int {
	return int(math.Pow(float64(num), 2))
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}

func main() {
	sumOfSquares := 0
	sumOfNums := 0
	for i := 1; i <= 100; i++ {
		sumOfNums = sumOfNums + i
		sumOfSquares = sumOfSquares + square(i)
	}

	fmt.Println(abs(sumOfSquares - square(sumOfNums)))
}
