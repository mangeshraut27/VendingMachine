package config

import (
	"errors"
	"testing"
)

var ReadFileMock = func(filename string) ([]byte, error) {
	return []byte(`{
		"products": [
        {
            "name": "Coke",
            "price": 25,
            "defaultQuantity": 0
        },
        {
            "name": "Pepsi",
            "price": 35,
            "defaultQuantity": 10
        },
        {
            "name": "Soda",
            "price": 45,
            "defaultQuantity": 10
        }
    ]
	}`), nil
}

var ReadFileMockError = func(filename string) ([]byte, error) {
	return []byte{}, errors.New("Test Error")
}

func TestLoadConfig(t *testing.T) {
	oldReadFile := ioutilReadFile

	defer func() {
		ioutilReadFile = oldReadFile
	}()
	t.Run("Success", func(t *testing.T) {
		ioutilReadFile = ReadFileMock
		err := LoadConfig("testfilepath")
		if err != nil {
			t.Fatalf("Expected nil, received error. Error : %v", err)
		}
	})
	t.Run("Error", func(t *testing.T) {
		ioutilReadFile = ReadFileMockError
		err := LoadConfig("testfilepath")
		if err == nil {
			t.Fatalf("Expected error, received nil.")
		}
	})
}
