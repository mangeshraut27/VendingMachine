package main

import (
	"VendingMachine/config"
	"VendingMachine/constant"
	"VendingMachine/operations"
	"log"
)

func main() {
	err := config.LoadConfig(constant.ConfigFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	operations.Start()
}
