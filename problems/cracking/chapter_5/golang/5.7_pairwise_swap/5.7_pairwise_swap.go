/*
Write a program to swap odd and even bits in an integer with as few instructions as possible (e.g., bit 13 and bit 1 are swapped, bit 2 and bit 3 are swapped, and so on).
EXAMPLE 1:
	 010101
	&010101
	=010101

 	 010101
	&101010
	=000000

	<<010101
	= 101010

	>>000000
	= 000000
	| 101010
	= 101010

EXAMPLE 2:
 	101101
	&010101
	=000101

 	101101
	&101010
	=101000

	<<000101
	= 001010

	>>101000
	= 010100
	| 001010
	= 011110
*/

package main

import "fmt"

func pairwise_swap(number int) int {
	odd_mask := 0x55555555
	even_mask := 0xAAAAAAAA

	odd_numbers := number & odd_mask
	even_numbers := number & even_mask

	odd_numbers = odd_numbers << 1
	even_numbers = even_numbers >> 1

	result := odd_numbers | even_numbers

	return result

}

func main() {
	fmt.Printf("%b\n", pairwise_swap(0b010101))
	fmt.Printf("%b\n", pairwise_swap(0b101101))
}
