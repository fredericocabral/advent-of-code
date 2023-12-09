package main

import "testing"

func TestExtractDigits(t *testing.T) {

	first, second := digits("1abc2")

	firstExpected := '1'
	if first != firstExpected {
		t.Fatalf("expected (%c) got (%c)", first, firstExpected)
	}

	secondExpected := '2'
	if second != secondExpected {
		t.Fatalf("expected (%c), got (%c)", second, secondExpected)
	}

}
