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
	signalChannel := make(chan struct{})
	signal.Notify(c, os.Interrupt)
	go operations.Start(signalChannel)
	select {
	case sig := <-c:
		if consumer.TotalAmountReceived > 0 {
			fmt.Println("Operation interuppted")
			fmt.Printf("Please Collect %d coins \n", consumer.TotalAmountReceived)
			consumer.TotalAmountReceived = 0
		}
		fmt.Printf("Got %s signal. Aborting Vending Machine Execution...\n", sig)
	case <-signalChannel:
		fmt.Println("Exiting")
	}
}
