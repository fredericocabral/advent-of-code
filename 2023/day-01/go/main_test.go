package main

import "testing"

func TestExtractDigits(t *testing.T) {

	t.Run("at edges", func(t *testing.T) {
		// act
		first, second := digits("1abc2")

		// assert
		assert(t, first, '1')
		assert(t, second, '2')
	})

	t.Run("in the value", func(t *testing.T) {
		// act
		first, second := digits("pqr3stu8vwx")

		// assert
		assert(t, first, '3')
		assert(t, second, '8')
	})

	t.Run("when there are more than two digits", func(t *testing.T) {
		// act
		first, second := digits("a1b2c3d4e5f")

		// assert
		assert(t, first, '1')
		assert(t, second, '5')
	})

	t.Run("when there is only one digit", func(t *testing.T) {
		// act
		first, second := digits("treb7uchet")
		// assert
		assert(t, first, '7')
		assert(t, second, '7')
	})

}

func assert(t *testing.T, received, expected rune) {
	if received != expected {
		t.Fatalf("received (%c) but expected (%c)", received, expected)
	}
}
