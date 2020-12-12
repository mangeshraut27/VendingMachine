package main

import (
	"VendingMachine/config"
	"VendingMachine/constant"
	"log"
)

func main() {
	err := config.LoadConfig(constant.ConfigFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
}
