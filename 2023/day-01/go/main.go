package main

import (
	"strconv"
	"unicode"
)

func calibrate(value string) int {
	var first, second rune

	for _, char := range value {
		if unicode.IsDigit(char) {
			if first == 0 {
				first = char
			}
			second = char
		}
	}

	val := string(first) + string(second)
	num, err := strconv.Atoi(val)
	if err != nil {
		panic("error: " + string(val))
	}

	return num
}

func sumAll(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += calibrate(line)
	}

	return sum
}
