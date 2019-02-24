package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reversInt(-567))
}

func reversInt(sourceInt int) int {

	intValueToReverse := sourceInt
	if sourceInt < 0 {
		intValueToReverse = intValueToReverse * -1
	}

	sourveStr := strconv.Itoa(intValueToReverse)

	sourceArr := strings.Split(sourveStr, "")

	var str string

	for _, char := range sourceArr {
		str = char + str
	}

	revInt, err := strconv.Atoi(str)

	if err != nil {
		fmt.Printf("string to int conversion failed with the following error %s ", err.Error())
	}

	if sourceInt < 0 {
		revInt = revInt * -1
	}

	return revInt
}
