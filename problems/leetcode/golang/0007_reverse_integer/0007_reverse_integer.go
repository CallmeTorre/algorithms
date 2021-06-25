/*
Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
	Input: x = 123
	Output: 321

Example 2:
	Input: x = -123
	Output: -321

Example 3:
	Input: x = 120
	Output: 21

Example 4:
	Input: x = 0
	Output: 0
*/

package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	if x == 0 && x > math.MaxInt32 {
		return 0
	}
	var result int
	stringNumber := strconv.Itoa(x)
	value := make([]byte, len(stringNumber))

	for i, j := 0, len(stringNumber)-1; i <= j; i, j = i+1, j-1 {
		value[i], value[j] = stringNumber[j], stringNumber[i]
	}

	if x < 0 {
		result, _ = strconv.Atoi(string(value[:len(value)-1]))
		result *= -1
	} else {
		result, _ = strconv.Atoi(string(value[:len(value)]))
	}
	if result >= math.MaxInt32 || result <= math.MinInt32 {
		return 0
	} else {
		return result
	}
}
