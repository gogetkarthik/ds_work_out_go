package user_input

import (
	"bufio"
	"fmt"
	"os"
)

func GetLine() string {

	fmt.Println("enter the required line input ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	text := scanner.Text()

	fmt.Println(text)

	return text

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
