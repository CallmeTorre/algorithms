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

func isPalindrome(x int) bool {
	/*
		When x < 0, x is not a palindrome.
		Also if the last digit of the number is 0, in order to be a palindrome,
		the first digit of the number also needs to be 0.
		Only 0 satisfy this property.
	*/
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	var revertedNumber int = 0
	for x > revertedNumber {
		xLastDigit := x % 10
		revertedNumber = revertedNumber*10 + xLastDigit
		x /= 10
	}
	if x == revertedNumber || x == revertedNumber/10 { // The operation x == revertedNumber/10 is for numbers with odd length
		return true
	}
	return false
}
