/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.
For example, 121 is palindrome while 123 is not.

Example 1:
	Input: x = 121
	Output: true

Example 2:
	Input: x = -121
	Output: false
	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
	Input: x = 10
	Output: false
	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Example 4:
	Input: x = -101
	Output: false
*/

package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	start, end := 0, len(s)-1
	for start <= end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
