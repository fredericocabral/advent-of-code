package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	GREEN = "green"
	BLUE  = "blue"
	RED   = "red"
)

func main() {

	readInputFromFile("input.txt")

}

func ParseInput(lines []string) [][]map[string]int {
	var games [][]map[string]int

	for _, line := range lines {
		game := strings.Split(line, ":")[1]
		sets := strings.Split(game, ";")

		var setsToAdd []map[string]int
		for _, set := range sets {

			cubesRevealed := strings.Split(set, ",")

			cubeToAdd := make(map[string]int)
			for _, cube := range cubesRevealed {
				cubeSliptted := strings.Split(cube, " ")
				color := cubeSliptted[2]
				quantity, _ := strconv.Atoi(cubeSliptted[1])
				cubeToAdd[color] = quantity
			}

			setsToAdd = append(setsToAdd, cubeToAdd)

		}

		games = append(games, setsToAdd)
	}

	return games
}

func SumNumberOfCuber(game []map[string]int) map[string]int {
	acc := make(map[string]int)

	for i := 0; i < len(game); i++ {
		for key, value := range game[i] {
			acc[key] = acc[key] + value
			//fmt.Printf("%s:%d\n", key, acc[key])
		}
	}

	return acc

}

func IsItPossible(game map[string]int) bool {

	for key, value := range game {
		if key == "red" && value > 12 {
			return false
		}

		if key == "green" && value > 13 {
			return false
		}

		if key == "blue" && value > 14 {
			return false
		}
	}

	return true
}

func SumAllPossibleGames(gamesResult []bool) int {

	sum := 0

	for i, isPossible := range gamesResult {
		if isPossible {
			sum = sum + i + 1
		}
	}

	return sum
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
