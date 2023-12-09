package main

import "testing"

func TestExtractDigits(t *testing.T) {

	t.Run("extract digits from value", func(t *testing.T) {
		// act
		first, second := digits("1abc2")

		// assert
		assert(t, first, '1')
		assert(t, second, '2')
	})

	t.Run("extract digits from other value", func(t *testing.T) {
		// act
		first, second := digits("pqr3stu8vwx")

		// assert
		assert(t, first, '3')
		assert(t, second, '8')
	})

}

func assert(t *testing.T, received, expected rune) {
	if received != expected {
		t.Fatalf("received (%c) but expected (%c)", received, expected)
	}
}
