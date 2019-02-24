package main

import (
	"fmt"
)

func main() {
	fmt.Println(handlerReversInt(-456))
}

func handlerReversInt(sourceInt int) int {

	intToReverse := sourceInt
	revisedInt := 0

	if sourceInt < 0 {
		intToReverse = sourceInt * -1
	}

	intToReverse, revisedInt = reverseInt(intToReverse, 0)

	if sourceInt < 0 {
		revisedInt = revisedInt * -1
	}

	return revisedInt
}

//24
func reverseInt(sourceInt, reversed int) (int, int) {

	if sourceInt < 10 {
		return 0, reversed*10 + sourceInt
	}

	reminder := sourceInt % 10
	reversed = reversed*10 + reminder
	return reverseInt(sourceInt/10, reversed)

}
