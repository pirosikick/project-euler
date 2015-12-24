'use strict';
/**
 * https://projecteuler.net/problem=4
 *
 * A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers.
 *
 * palindromic: 回文数
 */
const MIN_THREE_DIGIT_NUM = 100;
const MAX_THREE_DIGIT_NUM = 999;

const isPalindromicNum = num => {
  let str = num.toString();
  for (let i = 0; i < Math.floor(str.length / 2); i++) {
    if (str[i] != str[str.length - 1 - i]) {
      return false;
    }
  }
  return true;
};

let max = 0;
for (let i = MAX_THREE_DIGIT_NUM; i >= MIN_THREE_DIGIT_NUM; i--) {
  for (let j = i; j >= MIN_THREE_DIGIT_NUM; j--) {
    let product = i * j;
    if (isPalindromicNum(i * j) && max < product) {
      max = product;
    }
  }
}

console.log(max);
