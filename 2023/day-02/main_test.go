package main

import "testing"

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

func TestQuantityOfCubes(t *testing.T) {

	t.Run("A possible game", func(t *testing.T) {

		game := []map[string]int{
			{
				"blue": 3,
				"red":  4,
			},
			{
				"blue": 4,
				"red":  2,
			},
		}

		expected := true
		got := IsItPossible(game)

		if got != expected {
			t.Fatalf("got (%t) bue expected (%t)", got, expected)
		}

	})

	t.Run("A inpossible game", func(t *testing.T) {

		//Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red

		game := []map[string]int{
			{
				"gree": 8,
				"blue": 6,
				"red":  20,
			},
			{
				"blue":  5,
				"red":   4,
				"green": 13,
			},
			{
				"red":   1,
				"green": 5,
			},
		}

		expected := false
		got := IsItPossible(game)

		if got != expected {
			t.Fatalf("got (%t) bue expected (%t)", got, expected)
		}

	})

	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// id
	// sets:
	//	"color"
	//	"quantity"

}
