package main

import (
	"fmt"
)

func main() {
	fmt.Println(reversInt(-567))
}

func reversInt(sourceInt int) int {

	intToReverse := sourceInt
	revisedInt := 0

	if sourceInt < 0 {
		intToReverse = sourceInt * -1
	}

	for intToReverse > 0 {
		if intToReverse != 0 {
			reminder := intToReverse % 10

			revisedInt = revisedInt*10 + reminder

			intToReverse = intToReverse / 10
		}
	}

	if sourceInt < 0 {
		revisedInt = revisedInt * -1
	}

	return revisedInt
}
