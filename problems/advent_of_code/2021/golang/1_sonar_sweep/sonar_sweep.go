/*
--- Day 1: Sonar Sweep ---

You're minding your own business on a ship at sea when the overboard alarm goes off! Y
ou rush to see if you can help. Apparently, one of the Elves tripped and accidentally sent
the sleigh keys flying into the ocean!

Before you know it, you're inside a submarine the Elves keep ready for situations like this.
It's covered in Christmas lights (because of course it is), and it even has an experimental
antenna that should be able to track the keys if you can boost its signal strength high enough;
there's a little meter that indicates the antenna's signal strength by displaying 0-50 stars.

Your instincts tell you that in order to save Christmas, you'll need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent
calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

As the submarine drops below the surface of the ocean, it automatically performs a sonar sweep
of the nearby sea floor. On a small screen, the sonar sweep report (your puzzle input)
appears: each line is a measurement of the sea floor depth as the sweep looks further and further away from the submarine.

For example, suppose you had the following report:
199
200
208
210
200
207
240
269
260
263

This report indicates that, scanning outward from the submarine,
the sonar sweep found depths of 199, 200, 208, 210, and so on.

The first order of business is to figure out how quickly the depth increases, just so you know what you're dealing
with - you never know if the keys will get carried into deeper water by an ocean current or a fish or something.

To do this, count the number of times a depth measurement increases from the previous measurement.
(There is no measurement before the first measurement.) In the example above, the changes are as follows:

199 (N/A - no previous measurement)
200 (increased)
208 (increased)
210 (increased)
200 (decreased)
207 (increased)
240 (increased)
269 (increased)
260 (decreased)
263 (increased)
In this example, there are 7 measurements that are larger than the previous measurement.

How many measurements are larger than the previous measurement?

Your puzzle answer was 1387.

--- Part Two ---

Considering every single measurement isn't as useful as you expected: there's just too much noise in the data.

Instead, consider sums of a three-measurement sliding window. Again considering the above example:

199  A
200  A B
208  A B C
210    B C D
200  E   C D
207  E F   D
240  E F G
269    F G H
260      G H
263        H
Start by comparing the first and second three-measurement windows. The measurements in the first window
are marked A (199, 200, 208); their sum is 199 + 200 + 208 = 607. The second window is marked B (200, 208, 210);
its sum is 618. The sum of measurements in the second window is larger than the sum of the first, so this first comparison increased.

Your goal now is to count the number of times the sum of measurements in this sliding window increases from the previous sum.
So, compare A with B, then compare B with C, then C with D, and so on. Stop when there aren't enough measurements left to create a new three-measurement sum.

In the above example, the sum of each three-measurement window is as follows:

A: 607 (N/A - no previous sum)
B: 618 (increased)
C: 618 (no change)
D: 617 (decreased)
E: 647 (increased)
F: 716 (increased)
G: 769 (increased)
H: 792 (increased)
In this example, there are 5 sums that are larger than the previous sum.

Consider sums of a three-measurement sliding window. How many sums are larger than the previous sum?

Your puzzle answer was 1362.

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(filename string) []int {
	var result []int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't read file")
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Problem converting string to in")
			panic(err)
		}
		result = append(result, value)
	}
	file.Close()
	return result
}

func countDepthIncreases(values []int) int {
	result := 0
	previous_measure := values[0]
	for i := 1; i < len(values); i++ {
		current_measure := values[i]
		if current_measure > previous_measure {
			result += 1
		}
		previous_measure = current_measure
	}
	return result
}

func sum(slice []int) int {
	result := 0
	for i := range slice {
		result += slice[i]
	}
	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func countDepthIncreasesByThree(values []int) int {
	result := 0
	previous_measure := sum(values[0:3])
	for i := 0; i < len(values); i++ {
		current_values := values[i:min(i+3, len(values))]
		if len(current_values) == 3 {
			current_measure := sum(current_values)
			if current_measure > previous_measure {
				result += 1
			}
			previous_measure = current_measure
		} else {
			break
		}
	}
	return result
}

func main() {
	values := readFile("input.txt")
	fmt.Println(countDepthIncreases(values))
	fmt.Println(countDepthIncreasesByThree(values))
}
