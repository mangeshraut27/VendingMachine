package main

import (
	"VendingMachine/config"
	"VendingMachine/constant"
	"VendingMachine/consumer"
	"VendingMachine/operations"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	err := config.LoadConfig(constant.ConfigFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go operations.Start()
	select {
	case sig := <-c:
		fmt.Printf("Got %s signal. Aborting Vending Machine Execution...\n", sig)
		if consumer.TotalAmountReceived > 0 {
			fmt.Printf("Please Collect %d coins \n", consumer.TotalAmountReceived)
			consumer.TotalAmountReceived = 0
		}
	}
}
