package utils

import (
	"errors"
	"testing"
)

var fmtScanlnMock = func(a ...interface{}) (n int, err error) {
	return 0, nil
}

var fmtScanlnMockError = func(a ...interface{}) (n int, err error) {
	return 0, errors.New("Test Error")
}

func TestGetUserInput(t *testing.T) {
	oldScanln := fmtScanln
	defer func() {
		fmtScanln = oldScanln
	}()
	t.Run("Success", func(t *testing.T) {
		fmtScanln = fmtScanlnMock
		_, err := GetUserInput()
		if err != nil {
			t.Errorf("Expecting nil got error. Error : %v", err)
		}
	})

	t.Run("Error", func(t *testing.T) {
		fmtScanln = fmtScanlnMockError
		_, err := GetUserInput()
		if err == nil {
			t.Errorf("Expecting error got nil")
		}
	})
}
