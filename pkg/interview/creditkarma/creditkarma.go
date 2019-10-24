package main

import (
	"fmt"
	"time"
)

/*
1 6
2 7
3 8
4 9
3 6
2 5

groups

1 6 3 8
2 7 5
4 9
*/

func main() {
	fmt.Println("Hello, playground", time.Now())

	sourceArray := [][]int{{1, 6}, {2, 7}, {3, 8}, {4, 9}, {3, 6}, {2, 5}}

	fmt.Println(sourceArray)
	dt := make(map[int][]int, 0)

	for i := 0; i < len(sourceArray); i++ {
		_, ok := dt[sourceArray[i][0]]

		if !ok {
			dt[sourceArray[i][0]] = []int{sourceArray[i][1]}
		} else {
			dt[sourceArray[i][0]] = append(dt[sourceArray[i][0]], sourceArray[i][1])
		}

		_, ok = dt[sourceArray[i][1]]

		if !ok {
			dt[sourceArray[i][1]] = []int{sourceArray[i][0]}
		} else {
			dt[sourceArray[i][1]] = append(dt[sourceArray[i][1]], sourceArray[i][0])
		}

	}

	fmt.Println(dt)

	tm := make(map[int]int, 0)
	index := 0
	for key, val := range dt {
		fmt.Println("top", tm)
		if _, ok := tm[key]; !ok {
			tm[key] = index
			GetGroups(tm, dt, val, index)
			index++
		}
	}

	fmt.Println("before final", tm)
	final := make([][]int, index)
	for key, val := range tm {
		final[val] = append(final[val], key)
	}

	fmt.Println(final, time.Now())
}

func GetGroups(tm map[int]int, dt map[int][]int, value []int, index int) {
	fmt.Println("bottom", tm)
	for _, val := range value {
		if _, ok := tm[val]; !ok {
			tm[val] = index
			GetGroups(tm, dt, dt[val], index)
		}
	}
}
