package main

import (
	"testing"
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

		if IsItPossible(sumOfCubes) {
			t.Fatalf("this game should not be possible")
		}

	})

	t.Run("No, if green above limit", func(t *testing.T) {
		sumOfCubes := map[string]int{GREEN: 14, BLUE: 5, RED: 10}

		if IsItPossible(sumOfCubes) {
			t.Fatalf("this game should not be possible")
		}

	})

	t.Run("No, if blue above limit", func(t *testing.T) {
		sumOfCubes := map[string]int{GREEN: 11, BLUE: 15, RED: 10}

		if IsItPossible(sumOfCubes) {
			t.Fatalf("this game should not be possible")
		}

	})

	t.Run("Yes, if all below limit", func(t *testing.T) {
		sumOfCubes := map[string]int{GREEN: 13, BLUE: 14, RED: 12}

		if !IsItPossible(sumOfCubes) {
			t.Fatalf("this game should be possible")
		}

	})

}

func TestSumAllPossibleGames(t *testing.T) {

	games := []bool{true, true, false, false, true}

	got := SumAllPossibleGames(games)
	expected := 8

	if expected != got {
		t.Fatalf("expected %d, but got %d", expected, got)
	}

}

func TestParseInput(t *testing.T) {

	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red",
	}

	received := ParseInput(lines)

	expected := [][]map[string]int{
		{
			{BLUE: 3, RED: 4},
			{GREEN: 2, BLUE: 6, RED: 1},
			{GREEN: 2},
		},
		{
			{BLUE: 1, GREEN: 2},
			{GREEN: 3, BLUE: 4, RED: 1},
		},
	}

	if len(received) != len(expected) {
		t.Fatalf("Sizes don't match")
	}

	for i := 0; i < len(received); i++ {

		if len(received[i]) != len(expected[i]) {
			t.Fatalf("erro \n received: %d \n expected: %d", len(received[i]), len(expected[i]))
		}

		for j := 0; j < len(received[i]); j++ {

			if !mapsAreEqual(received[i][j], expected[i][j]) {
				t.Fatalf("not equal \n Received:%v \n Expected:%v", received[i][j], expected[i][j])
			}

		}

	}
}

func mapsAreEqual(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		value2, ok := map2[key]
		if !ok || value1 != value2 {
			return false
		}
	}

	return true
}
