package utils

import "fmt"

var (
	fmtScanln = fmt.Scanln
)

//GetUserInput GetUserInput
var GetUserInput = func() (int, error) {
	var input int
	_, err := fmtScanln(&input)
	return input, err
}
