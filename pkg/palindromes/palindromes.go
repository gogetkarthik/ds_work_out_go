package main

import (
	"fmt"
	"strings"
)

func main() {

	//str := "karthik"

	pal := "refer"

	//fmt.Println(Palindromes(pal))
	//fmt.Println(Palindromes(str))
	//fmt.Println(Palindromes(pal))
	fmt.Println(Palindrome1(pal))
}

func Palindromes(str string) bool {

	rn := []rune(str)
	for i, j := 0, len(rn)-1; i < j; i, j = i+1, j-1 {
		rn[i], rn[j] = rn[j], rn[i]
	}

	revStr := string(rn)

	return revStr == str

}

func Palindrome1(str string) bool {

	sourceArr := strings.Split(str, "")

	var targetArr string

	for i, char := range sourceArr {
		targetArr = char + targetArr
		fmt.Println(i, targetArr)

	}

	return true
}
