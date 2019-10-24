package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	fmt.Println(repeatedString("ababa", 100000000000000000))
}

func repeatedString(s string, n int64) int64 {

	sLn := int64(len(s))

	return ((n / sLn) * getCount(s)) + getCount(s[0:int(n%sLn)])
}

func getCount(s string) int64 {

	var numAs int64
	for idx, ch := range s {
		fmt.Printf("%d, %c\n", idx, ch)
		if 'a' == ch {
			numAs++
		}

	}
	return numAs
}
