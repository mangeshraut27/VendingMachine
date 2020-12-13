package operations

import (
	"VendingMachine/consumer"
	"VendingMachine/product"
	"VendingMachine/supplier"
	"fmt"
	"io"
	"log"
)

//Start : Starts Vending Machine
func Start() {
	product.ConvertProductToMap()
	fmt.Println("Welcome at Vending Machine")
	var input int

	for {
		fmt.Println("Enter 1 for Supplier, 2 for User and 3 for Exit")
		_, err := fmt.Scanln(&input)
		if err != nil {
			if err != io.EOF {
				log.Printf("Runtime error occured : %v \n", err)
			}
			return
		}
		switch input {
		case 1:
			supplier.AccessMachine()
		case 2:
			consumer.Consume()
		case 3:
			return
		default:
			fmt.Println("Invalid Input. Please try again")
			continue
		}
	}
}
