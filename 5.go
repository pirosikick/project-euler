package main

import (
	"fmt"
	"math"
)

/*
https://projecteuler.net/problem=://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func main() {
	max := 20
	primes := getPrimes(max)
	primeFactors := make(map[int]int)
	for _, prime := range primes {
		primeFactors[prime] = 0
	}

	for i := max; i >= 1; i-- {
		factors := calcPrimeFactors(i, primes)
		for prime, count := range factors {
			if count > primeFactors[prime] {
				primeFactors[prime] = count
			}
		}
	}

	answer := 1
	for prime, count := range primeFactors {
		p := float64(prime)
		c := float64(count)
		answer = answer * int(math.Pow(p, c))
	}

	fmt.Printf("The answer of problem 5 is %d.\n", answer)
}

func calcPrimeFactors(num int, primes []int) map[int]int {
	factors := make(map[int]int)
	for _, prime := range primes {
		for ; num%prime == 0; num = num / prime {
			count, ok := factors[prime]
			if ok {
				factors[prime] = count + 1
			} else {
				factors[prime] = 1
			}
		}
		if num == 1 {
			break
		}
	}
	return factors
}

func getPrimes(threshold int) []int {
	primes := []int{2}
	num := 3
	for num <= threshold {
		if isPrimeNumber(num) {
			primes = append(primes, num)
		}
		num = num + 2
	}
	return primes
}

func isPrimeNumber(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
