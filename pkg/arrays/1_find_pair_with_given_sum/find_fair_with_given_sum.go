package main

import (
	"data_structure_work_out_go/pkg/user_input"
	"data_structure_work_out_go/pkg/util"
	"fmt"
	"sort"
	"strconv"
)

func main() {

	fmt.Println("find pair with given sum")

	sum, _ := strconv.Atoi(user_input.GetLine())

	fmt.Printf("given sum : %d \n", sum)

	arrSlice := util.StringSliceToInt(user_input.GetSlice())

	fmt.Printf("given int array %d \n", arrSlice)

	//Order of O(n^2)
	for i := range arrSlice {

		for j := i + 1; j < len(arrSlice); j++ {

			if arrSlice[i]+arrSlice[j] == sum {
				fmt.Printf("found match %d - %d \n", i, j)
				break
			}
		}

	}

	var arrMap = make(map[int]int)
	// O(n) time complexity
	for i, val := range arrSlice {
		if val2, ok := arrMap[sum-arrSlice[i]]; ok {
			fmt.Printf("map found match %d - %d \n", i, val2)
		} else {
			arrMap[val] = i
		}
	}

	sort.Ints(arrSlice)

	fmt.Printf("sorted array %d \n", arrSlice)

	i := 0
	j := len(arrSlice) - 1

	for i < j {

		if arrSlice[i]+arrSlice[j] == sum {
			fmt.Printf("sort pair found %d - %d \n", i, j)
			i++
			j--
		} else if arrSlice[i]+arrSlice[j] > sum {
			j--
		} else if arrSlice[i]+arrSlice[j] < sum {
			i++
		}
	}

}
