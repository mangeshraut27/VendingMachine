package utils

import "fmt"

//GetUserInput GetUserInput
func GetUserInput() (int, error) {
	var input int
	_, err := fmt.Scanln(&input)
	return input, err
}
