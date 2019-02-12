package main

import (
	"data_structure_work_out_go/pkg/user_input"
	"fmt"
)

func main() {

	inputString := user_input.GetLine()
	length := len(inputString)
	bArr := make([]rune, length)
	for i, char := range inputString {
		fmt.Printf("%s\n", string(char))
		bArr[length-1-i] = char
	}

	fmt.Printf("%v\nrune", string(bArr))

	rn := []rune(inputString)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		rn[i], rn[j] = rn[j], rn[i]
	}

	fmt.Println(string(rn))
}
