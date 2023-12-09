package main

import (
	"unicode"
)

func calibrate(value string) (rune, rune) {

	var first rune
	var second rune

	for _, char := range value {
		if unicode.IsDigit(char) {
			if first == 0 {
				first = char
			}
			second = char
		}
	}

	return first, second
}
