'use strict';
/**
 * Q1.
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5,
 * we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

function isMultipleOf(num, target) {
  return (target % num) === 0;
}

const isMultipleOf3 = isMultipleOf.bind(undefined, 3);
const isMultipleOf5 = isMultipleOf.bind(undefined, 5);
let sum = 0;

for (let i = 1; i < 1000; i++) {
  if (isMultipleOf3(i) || isMultipleOf5(i)) {
    sum += i;
  }
}

console.log(sum);
