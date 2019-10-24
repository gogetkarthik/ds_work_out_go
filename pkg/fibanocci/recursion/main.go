package main

import "fmt"

//https://play.golang.org/p/zga_82B1fOj
func main() {

	numFib := 14

	fmt.Printf("The computed fibonacci is %d ", fibonacci(numFib))
}

func fibonacci(numFib int) int {

	if numFib == 1 {
		return 1
	}

	if numFib == 0 {
		return 0
	}

	//this should be return as
	/*

		if numFib < 2 {
			return numFib
		}
	*/

	return fibonacci(numFib-1) + fibonacci(numFib-2)
}
