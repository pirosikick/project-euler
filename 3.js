'use strict';
const assert = require("assert");

/**
 * https://projecteuler.net/problem=3
 *
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 *
 * prime factor: 素因数
 */
const theNum = 600851475143;

const isPrimeNumber = num => {
  for (let i = 2; i < num; i++) {
    if (num % i === 0) return false;
  }
  return true;
}

const lengthOfPrimes = 10000;
const primes = (function (threshold) {
  let primes = [2];
  let num = 3;
  while (num <= threshold) {
    if (isPrimeNumber(num)) {
      primes.push(num);
    }
    num += 2;
  }

  return primes;
})(lengthOfPrimes);

let _theNum = theNum;
let maxPrime;
const mustbe1 = primes.reduce((num, prime) => {
  while (1) {
    if (num % prime !== 0) {
      return num;
    }

    console.log(`a prime factor of ${theNum} was found:`, prime);

    maxPrime = prime;
    num /= prime;
  }
}, theNum);

assert(mustbe1 === 1, "too short primes.length");

console.log(maxPrime);
