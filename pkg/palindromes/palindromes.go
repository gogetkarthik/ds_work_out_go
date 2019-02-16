package main

import "fmt"

func main() {

	str := "karthik"

	pal := "refer"

	fmt.Println(Palindromes(pal))
	fmt.Println(Palindromes(str))
	fmt.Println(Palindromes(pal))
}

func Palindromes(str string) bool {

	rn := []rune(str)
	for i, j := 0, len(rn)-1; i < j; i, j = i+1, j-1 {
		rn[i], rn[j] = rn[j], rn[i]
	}

	revStr := string(rn)

	return revStr == str

}
