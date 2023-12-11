package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func first(line string) int {
	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return extract(line, numbers)
}

func last(line string) int {
	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var reversedNumbers [9]string

	for n := 0; n < len(numbers); n++ {
		reversedNumbers[n] = reverse(numbers[n])
	}

	return extract(reverse(line), reversedNumbers)
}

func extract(line string, numbers [9]string) int {

	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			digit, _ := strconv.Atoi(string(line[i]))
			return digit
		}

		for n := 0; n < len(numbers); n++ {
			if len(line[i:]) >= len(numbers[n]) {
				if numbers[n] == line[i:len(numbers[n])+i] {
					return n + 1
				}
			}
		}
	}
	return 0
}

func reverse(value string) string {
	var reversed string

	for i := len(value) - 1; i >= 0; i-- {
		reversed = reversed + string(value[i])
	}

	return reversed
}

func sumAll(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += first(line)*10 + last(line)
	}

	return sum
}

func main() {
	input := readInputFromFile("input.txt")

	calibration := sumAll(input)

	fmt.Printf("Result: (%d) ", calibration)
}

func readInputFromFile(fileName string) []string {
	var input []string

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
