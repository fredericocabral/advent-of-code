package main

import (
	"testing"
)

func TestCalibrationNew(t *testing.T) {

	t.Run("at edges", func(t *testing.T) {
		// act
		first := first("1abc2")
		last := last("1abc2")

		// assert
		assert(t, first, 1)
		assert(t, last, 2)
	})

	t.Run("at edges 2", func(t *testing.T) {
		// act
		first := first("vsskdclbtmjmvrseven6")
		last := last("vsskdclbtmjmvrseven6")

		// assert
		assert(t, first, 7)
		assert(t, last, 6)
	})

	t.Run("at edges 3", func(t *testing.T) {
		// act
		first := first("zlmlk1")
		last := last("zlmlk1")

		// assert
		assert(t, first, 1)
		assert(t, last, 1)
	})

	t.Run("in the value", func(t *testing.T) {
		// act
		first := first("pqr3stu8vwx")
		last := last("pqr3stu8vwx")

		// assert
		assert(t, first, 3)
		assert(t, last, 8)
	})

	t.Run("when there are more than two digits", func(t *testing.T) {
		// act
		first := first("a1b2c3d4e5f")
		last := last("a1b2c3d4e5f")

		// assert
		assert(t, first, 1)
		assert(t, last, 5)
	})

	t.Run("when there is only one digit", func(t *testing.T) {
		// act
		first := first("treb7uchet")
		last := last("treb7uchet")

		// assert
		assert(t, first, 7)
		assert(t, last, 7)
	})

	t.Run("two and nine spelled", func(t *testing.T) {
		// act
		first := first("two1nine")
		last := last("two1nine")

		// assert
		assert(t, first, 2)
		assert(t, last, 9)
	})

	t.Run("eight and three", func(t *testing.T) {
		// act
		first := first("eightwothree")
		last := last("eightwothree")

		// assert
		assert(t, first, 8)
		assert(t, last, 3)
	})

	t.Run("one and three", func(t *testing.T) {
		// act
		first := first("abcone2threexyz")
		last := last("abcone2threexyz")

		// assert
		assert(t, first, 1)
		assert(t, last, 3)
	})

	t.Run("two and four", func(t *testing.T) {
		// act
		first := first("xtwone3four")
		last := last("xtwone3four")

		// assert
		assert(t, first, 2)
		assert(t, last, 4)
	})

	t.Run("starting and finishing with real numbers", func(t *testing.T) {
		// act
		first := first("4nineeightseven2")
		last := last("4nineeightseven2")

		// assert
		assert(t, first, 4)
		assert(t, last, 2)
	})

	t.Run("spelled with real number", func(t *testing.T) {
		// act
		first := first("zoneight234")
		last := last("zoneight234")

		// assert
		assert(t, first, 1)
		assert(t, last, 4)
	})

	t.Run("spelled with real number 2", func(t *testing.T) {
		// act
		first := first("7pqrstsixteen")
		last := last("7pqrstsixteen")

		// assert
		assert(t, first, 7)
		assert(t, last, 6)
	})

}

func TestReverse(t *testing.T) {

	got := reverse("one")
	what := "eno"

	if got != what {
		t.Fatalf("received (%s) but expected (%s)", got, what)
	}

}

func TestSumOfAllCalibrationValues(t *testing.T) {
	// arrange
	var calibrations []string

	calibrations = append(calibrations, "two1nine")
	calibrations = append(calibrations, "eightwothree")
	calibrations = append(calibrations, "abcone2threexyz")
	calibrations = append(calibrations, "xtwone3four")
	calibrations = append(calibrations, "4nineeightseven2")
	calibrations = append(calibrations, "zoneight234")
	calibrations = append(calibrations, "7pqrstsixteen")

	// act
	got := sumAll(calibrations)

	// assert
	assert(t, got, 281)
}

func assert(t *testing.T, received, expected int) {
	if received != expected {
		t.Fatalf("received (%d) but expected (%d)", received, expected)
	}
}
