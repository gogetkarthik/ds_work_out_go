package main

import "fmt"

func main() {

	inputNum := 30
	isFizBuzz := false
	for i := 1; i <= inputNum; i++ {
		if i%3 == 0 {
			isFizBuzz = true
			fmt.Printf("Fizz")
		}

		if i%5 == 0 {
			isFizBuzz = true
			fmt.Printf("Buzz")
		}

		if !isFizBuzz {
			fmt.Printf("%d", i)
		}
		isFizBuzz = false
		fmt.Println()
	}
}
