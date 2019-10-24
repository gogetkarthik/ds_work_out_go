package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	arrayToBeSorted := []int{5, 3, 9, 6, 1, 8, 2, 4, 7}
	insertionSort(arrayToBeSorted)
	fmt.Printf("sorted array %d\n", arrayToBeSorted)
}

func insertionSort(sourceArray []int) {

	sLen := len(sourceArray)
	for i := 0; i < sLen; i++ {

		for j := i; j > 0; j-- {

			if sourceArray[j] < sourceArray[j-1] {
				sourceArray[j], sourceArray[j-1] = sourceArray[j-1], sourceArray[j]
			} else {
				break
			}
		}
	}

	fmt.Printf("sorted array %d\n", sourceArray)
}
