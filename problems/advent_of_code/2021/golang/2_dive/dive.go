/*
--- Day 2: Dive! ---

Now, you need to figure out how to pilot this thing.

It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

* forward X increases the horizontal position by X units.
* down X increases the depth by X units.
* up X decreases the depth by X units.

Note: that since you're on a submarine, down and up affect your depth, and so they have the
opposite result of what you might expect.

The submarine seems to already have a planned course (your puzzle input).
You should probably figure out where it's going. For example:

forward 5
down 5
forward 8
up 3
down 8
forward 2

Your horizontal position and depth both start at 0. The steps above would then modify them as follows:

* forward 5 adds 5 to your horizontal position, a total of 5.
* down 5 adds 5 to your depth, resulting in a value of 5.
* forward 8 adds 8 to your horizontal position, a total of 13.
* up 3 decreases your depth by 3, resulting in a value of 2.
* down 8 adds 8 to your depth, resulting in a value of 10.
* forward 2 adds 2 to your horizontal position, a total of 15.

After following these instructions, you would have a horizontal position of 15 and a depth of 10.
(Multiplying these together produces 150.)

Calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?

Your puzzle answer was 1868935.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) [][]string {
	var result [][]string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't read file")
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		instruction := strings.Fields(string(scanner.Text()))
		result = append(result, instruction)
	}
	file.Close()
	return result
}

func getDestination(instructions [][]string) int {
	horizontal := 0
	depth := 0

	for _, instruction := range instructions {
		command := instruction[0]
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatalf("Can't convert to int")
			panic(err)
		}
		if command == "forward" {
			horizontal += value
		} else if command == "up" {
			depth -= value
		} else if command == "down" {
			depth += value
		}
	}
	return horizontal * depth
}

func main() {
	values := readFile("input.txt")
	fmt.Println(getDestination(values))
}
