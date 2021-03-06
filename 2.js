'use strict';
/**
 * https://projecteuler.net/problem=2
 *
 * Each new term in the Fibonacci sequence is generated by adding the previous
 * two terms. By starting with 1 and 2, the first 10 terms will be:
 *
 * 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
 *
 * By considering the terms in the Fibonacci sequence whose values do not exceed
 * four million, find the sum of the even-valued terms.
 */

const fourMillion = 1000 * 1000 * 4;
const isEven = num => num % 2 === 0;

function fibonacciList(threshold) {
  let nums = [1, 2];

  while (1) {
    let next = nums[nums.length - 2] + nums[nums.length - 1];
    if (next > threshold) {
      break;
    }
    nums.push(next);
  }

  return nums;
}

let evenTotal = fibonacciList(fourMillion)
  .filter(isEven)
  .reduce((total, current) => total + current, 0);

console.log(evenTotal);
