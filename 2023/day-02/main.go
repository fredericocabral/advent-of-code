package main

func main() {

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

func IsItPossible(game []map[string]int) bool {

	acc := make(map[string]int)

	for i := 0; i < len(game); i++ {
		for key, value := range game[i] {

			acc[key] = acc[key] + value

		}
	}

	if acc["red"] > 10 {
		return false
	}

	return true
}
