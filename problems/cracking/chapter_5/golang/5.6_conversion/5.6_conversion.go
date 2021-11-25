/*
Write a function to determine the number of bits you would need to flip to convert integer A to integer B.
EXAMPLE:
	Input: 29 (or: 11101), 15 (or: 01111)
	Output: 2
*/

package main

import "fmt"

func conversion(x, y int) int {
	fmt.Printf("x: %08b\n", x)
	fmt.Printf("y: %08b\n", y)

	number := x ^ y

	fmt.Printf("number: %08b\n", number)

	count := 0
	for number != 0 {
		number = number & (number - 1)
		fmt.Printf("number: %08b\n", number)
		count++
	}
	return count
}

func main() {
	fmt.Printf("Result: %d\n", conversion(29, 15))
	fmt.Printf("Result: %d\n", conversion(39, 11))

}
