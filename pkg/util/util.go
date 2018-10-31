package util

import "strconv"

func StringSliceToInt(s []string) []int {

	var iSlice []int
	for _, str := range s {
		i, _ := strconv.Atoi(str)
		iSlice = append(iSlice, i)
	}

	return iSlice
}
