// 10001st prime
// https://projecteuler.net/problem=7
//
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
package main

import (
	"fmt"
)

func isPrimeNumber(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var nthPrime int
	n := 10001
	count := 1
	for i := 3; count < n; i += 2 {
		if isPrimeNumber(i) {
			nthPrime = i
			count += 1
		}
	}
	fmt.Printf("%dst prime number is %d\n", n, nthPrime)
}
