//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Hello, playground")
//
//	sourceArray := []int{9, 2, 6, 1, 5, 3, 7, 4, 8}
//
//	quickSort(sourceArray, 0, len(sourceArray)-1)
//
//	fmt.Printf("%v", sourceArray)
//
//}
//
//func quickSort(sourceArray []int, startIndex, endIndex int) {
//
//	fmt.Println(sourceArray)
//
//	if startIndex >= endIndex {
//		return
//	}
//
//	pIndex := partician(sourceArray, startIndex, endIndex, sourceArray[endIndex/2])
//
//	fmt.Println(pIndex)
//
//	quickSort(sourceArray, pIndex+1, endIndex)
//	quickSort(sourceArray, startIndex, pIndex)
//
//}
//
//func partician(sourceArray []int, startIndex, endIndex, pivot int) int {
//
//	fmt.Println(sourceArray, startIndex, endIndex, pivot)
//	indexStart := startIndex
//	indexEnd := endIndex
//
//	for startIndex < endIndex {
//		for sourceArray[indexStart] < pivot {
//
//			indexStart++
//			fmt.Printf("index start %d\n", indexStart)
//		}
//
//		for sourceArray[indexEnd] > pivot {
//			indexEnd--
//			fmt.Printf("index end %d\n", indexEnd)
//		}
//
//		if indexStart >= indexEnd {
//			return indexStart
//		}
//
//		sourceArray[indexStart], sourceArray[indexEnd] = sourceArray[indexEnd], sourceArray[indexStart]
//		indexStart++
//		indexEnd--
//	}
//
//	return indexStart
//}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	//sourceArray := []int{9, 2, 6, 1, 5, 3, 7, 4, 8}

	//sourceArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sourceArray := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	quickSort(sourceArray, 0, len(sourceArray)-1)

	fmt.Printf("%v", sourceArray)

}

func quickSort(sourceArray []int, startIndex, endIndex int) {

	fmt.Println(sourceArray)

	if startIndex >= endIndex {
		return
	}

	pIndex := partician(sourceArray, startIndex, endIndex, sourceArray[endIndex/2])

	fmt.Println(pIndex)

	quickSort(sourceArray, pIndex+1, endIndex)
	quickSort(sourceArray, startIndex, pIndex)

}

func partician(sourceArray []int, startIndex, endIndex, pivot int) int {

	fmt.Println(sourceArray, startIndex, endIndex, pivot)
	indexStart := startIndex
	indexEnd := endIndex

	for indexStart < indexEnd {
		fmt.Println("KKK")
		for sourceArray[indexStart] < pivot {
			indexStart++
			fmt.Printf("index start %d\n", indexStart)
		}

		for sourceArray[indexEnd] > pivot {
			indexEnd--
			fmt.Printf("index end %d\n", indexEnd)
		}

		if indexStart < indexEnd {
			sourceArray[indexStart], sourceArray[indexEnd] = sourceArray[indexEnd], sourceArray[indexStart]
			indexStart++
			indexEnd--
		}

	}

	return indexStart
}
