package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	arrayToBeSorted := []int{5, 3, 9, 6, 1, 8, 2, 4, 7}
	fmt.Printf("source array %d\n", arrayToBeSorted)
	selectionSort(arrayToBeSorted)
	fmt.Printf("sorted array %d", arrayToBeSorted)
}

func selectionSort(sourceArray []int) {

	for i := 0; i < len(sourceArray); i++ {
		minimumElementPotiton := i
		for j := i; j < len(sourceArray); j++ {
			if sourceArray[minimumElementPotiton] > sourceArray[j] {

				minimumElementPotiton = j
			}
		}
		sourceArray[i], sourceArray[minimumElementPotiton] = sourceArray[minimumElementPotiton], sourceArray[i]

	}
}
