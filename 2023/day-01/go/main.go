package main

import (
	"strconv"
	"unicode"
)

func calibrate(line string) int {

	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	first, last := 0, 0

	for len(line) > 0 {
		found := false
		if unicode.IsDigit(rune(line[0])) {
			if first == 0 {
				first, _ = strconv.Atoi(string(line[0]))
			} else {
				last, _ = strconv.Atoi(string(line[0]))
			}

		} else {

			for i := 0; i < len(numbers); i++ {

				currentNumber := numbers[i]

				if len(line) >= len(currentNumber) {
					if line[0:len(currentNumber)] == currentNumber {
						if first == 0 {
							first = i + 1
						} else {
							last = i + 1
						}

						line = line[len(currentNumber):]
						found = true
						break

					}
				}

			}
		}

		if !found {
			line = line[1:]
		}

	}

	if last == 0 {
		last = first
	}

	return first*10 + last

}

func sumAll(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += calibrate(line)
	}

	return sum
}
