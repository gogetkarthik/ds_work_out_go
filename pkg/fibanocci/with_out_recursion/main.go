package main

import (
	"fmt"
)

//https://play.golang.org/p/I9H_c_Chm5-
func main() {
	fmt.Println("Hello, playground")

	iStartVal := 0
	iEndVal := 1
	numFib := 15

	fmt.Printf("%d %d", iStartVal, iEndVal)

	for i := 0; i < numFib-1; i++ {
		fmt.Printf(" %d", iStartVal+iEndVal)
		iStartVal, iEndVal = iEndVal, iStartVal+iEndVal
	}

}
