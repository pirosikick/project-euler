// Special Pythagorean triplet
// https://projecteuler.net/problem=9
//
// Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
// a^2 + b^2 = c^2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
package main

import (
	"fmt"
)

func valid(a, b, c int) bool {
	return (a*a)+(b*b) == (c * c)
}

func main() {
	for a := 1; a <= 332; a++ {
		for b := a + 1; b <= ((1000 - a) / 2); b++ {
			c := 1000 - (a + b)
			if valid(a, b, c) {
				fmt.Printf("a=%d, b=%d, c=%d, abc=%d", a, b, c, a*b*c)
				return
			}
		}
	}
}
