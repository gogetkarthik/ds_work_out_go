package main

import "fmt"

func main() {

	initialArray := []int{4, -6, 3, -1, 4, 2, 7}
	//initialArray := []int{3,4,-7,3,1,3,1,-4,-2,-2}

	fmt.Println(initialArray)

	for i := range initialArray {

		sum := initialArray[i]
		for j := i + 1; j < len(initialArray); j++ {

			sum += initialArray[j]
			if sum == 0 {

				printSubArray(initialArray[i : j+1])
			}

		}

	}
}

func printSubArray(sa []int) {

	fmt.Println(sa)
}
