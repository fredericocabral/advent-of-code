package main

import (
	"testing"
)

const (
	GREEN = "green"
	BLUE  = "blue"
	RED   = "red"
)

func TestSumNumberOfCubes(t *testing.T) {

	game := []map[string]int{
		{
			GREEN: 8, BLUE: 6, RED: 20,
		},
		{
			BLUE: 5, RED: 4, GREEN: 13,
		},
		{
			RED: 1, GREEN: 5,
		},
	}

	expected := map[string]int{GREEN: 26, BLUE: 11, RED: 25}

	got := SumNumberOfCuber(game)

	for key, valueExpected := range expected {
		valueGot, ok := got[key]
		if !ok || valueExpected != valueGot {
			t.Fatalf("got %s:%d but expected %s:%d", key, valueGot, key, valueExpected)
		}
	}

}

func TestIsPossible(t *testing.T) {

	t.Run("No, if red above limit", func(t *testing.T) {

		sumOfCubes := map[string]int{GREEN: 26, BLUE: 11, RED: 25}

		got := IsItPossible(sumOfCubes)
		expected := false

		if got != expected {
			t.Fatalf("got (%t) bue expected (%t)", got, expected)
		}

	})

	t.Run("No, if green above limit", func(t *testing.T) {

		sumOfCubes := map[string]int{GREEN: 14, BLUE: 5, RED: 10}

		got := IsItPossible(sumOfCubes)
		expected := false

		if got != expected {
			t.Fatalf("got (%t) bue expected (%t)", got, expected)
		}

	})

	t.Run("No, if blue above limit", func(t *testing.T) {

		sumOfCubes := map[string]int{GREEN: 11, BLUE: 14, RED: 10}

		got := IsItPossible(sumOfCubes)
		expected := false

		if got != expected {
			t.Fatalf("got (%t) bue expected (%t)", got, expected)
		}

	})

}
