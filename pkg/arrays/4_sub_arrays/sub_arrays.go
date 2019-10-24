package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	sArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(subArray(sArray, 2))
}

func subArray(sArray []int, size int) [][]int {
	sLen := len(sArray)
	numTargetSize := sLen / size
	if sLen%size != 0 {
		numTargetSize += 1
	}

	targetArray := make([][]int, numTargetSize)
	for i := 0; i < numTargetSize; i++ {
		targetArray[i] = make([]int, size)
		lSize := 0
		if (i+1)*size < sLen {
			lSize = (i + 1) * size
		} else {
			lSize = sLen
		}
		targetArray[i] = sArray[(i * size):lSize]
	}
	return targetArray
}
