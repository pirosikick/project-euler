// Summation of primes
// https://projecteuler.net/problem=10
//
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
package main

import (
	"fmt"
	"math"
)

func sqrt(num int) int {
	return int(math.Sqrt(float64(num)))
}

func isPrimeNumber(num int) bool {
	for i := 3; i <= sqrt(num); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	threshold := 2000000
	sum := 2
	for i := 3; i < threshold; i += 2 {
		if isPrimeNumber(i) {
			sum += i
			fmt.Printf("prime:%d, sub:%d\n", i, sum)
		}
	}
	fmt.Printf("the sum of all the primes below two million is %d", sum)
}
