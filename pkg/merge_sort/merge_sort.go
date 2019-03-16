package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	arrayToBeSorted := []int{5, 3, 9, 6, 1, 8, 2, 4, 7}
	arraySorted := mergeSort(arrayToBeSorted)

	fmt.Printf("sorted array %d\n", arraySorted)
}

func mergeSort(sourceArray []int) []int {

	if len(sourceArray) <= 1 {
		return sourceArray
	}

	sort1 := mergeSort(sourceArray[:len(sourceArray)/2])
	sort2 := mergeSort(sourceArray[(len(sourceArray) / 2):])

	fmt.Printf("source array 1 %d\n", sort1)
	fmt.Printf("source array 2 %d\n", sort2)

	s1 := 0
	s2 := 0

	targetArray := []int{}

	for {

		if sort1[s1] < sort2[s2] {

			targetArray = append(targetArray, sort1[s1])
			s1++
		} else {
			targetArray = append(targetArray, sort2[s2])
			s2++
		}

		if s1 >= len(sort1) {
			break
		}

		if s2 >= len(sort2) {
			break
		}

	}

	targetArray = append(targetArray, sort1[s1:len(sort1)]...)
	targetArray = append(targetArray, sort2[s2:len(sort2)]...)

	fmt.Printf("target array %d \n", targetArray)

	return targetArray
}
