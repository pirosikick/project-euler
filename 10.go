// Summation of primes
// https://projecteuler.net/problem=10
//
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
package main

import (
	"fmt"
)

func isPrimeNumber(num int, primes []int) bool {
	for _, prime := range primes {
		if num%prime == 0 {
			return false
		}
	}
	return true
}

func main() {
	threshold := 2000000
	sum := 2
	primes := []int{2}
	for i := 3; i < threshold; i += 2 {
		if isPrimeNumber(i, primes) {
			sum += i
			primes = append(primes, i)
			fmt.Printf("prime:%d, sub:%d\n", i, sum)
		}
	}
	fmt.Printf("the sum of all the primes below two million is %d", sum)
}
