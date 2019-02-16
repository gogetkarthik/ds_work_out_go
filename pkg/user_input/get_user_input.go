package user_input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetLine() string {

	//fmt.Println("enter the required line input ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	text := scanner.Text()

	fmt.Println(text)

	return text

}

func GetInt() int {
	str := GetLine()

	iValue, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		fmt.Println("str to int convert failed " + err.Error())
		os.Exit(1)
	}

	return int(iValue)

}

func GetSlice() []string {
	fmt.Println("enter the required slice")

	scanner := bufio.NewScanner(os.Stdin)

	var text []string

	for scanner.Scan() {

		text = append(text, scanner.Text())

		fmt.Println(text)
	}

	return text
}
