package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestCalibration(t *testing.T) {

	t.Run("at edges", func(t *testing.T) {
		// act
		calibration := calibrate("1abc2")

		// assert
		assert(t, calibration, 12)
	})

	t.Run("in the value", func(t *testing.T) {
		// act
		calibration := calibrate("pqr3stu8vwx")

		// assert
		assert(t, calibration, 38)
	})

	t.Run("when there are more than two digits", func(t *testing.T) {
		// act
		calibration := calibrate("a1b2c3d4e5f")

		// assert
		assert(t, calibration, 15)
	})

	t.Run("when there is only one digit", func(t *testing.T) {
		// act
		calibration := calibrate("treb7uchet")
		// assert
		assert(t, calibration, 77)
	})
}

func TestSumOfAllCalibrationValues(t *testing.T) {
	// arrange
	var calibrations []string
	calibrations = append(calibrations, "1abc2")
	calibrations = append(calibrations, "pqr3stu8vwx")

	// act
	got := sumAll(calibrations)

	// assert
	assert(t, got, 50)
}

func TestSumFromLinesInFile(t *testing.T) {

	// arrange
	input := readInputFromFile("input.txt")

	// act
	got := sumAll(input)

	// assert
	assert(t, got, 54697)
}

func assert(t *testing.T, received, expected int) {
	if received != expected {
		t.Fatalf("received (%d) but expected (%d)", received, expected)
	}
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
