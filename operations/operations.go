package operations

import (
	"VendingMachine/consumer"
	"VendingMachine/product"
	"VendingMachine/supplier"
	"VendingMachine/utils"
	"fmt"
	"io"
	"log"
)

//Start : Starts Vending Machine
func Start() {
	product.ConvertProductToMap()
	fmt.Println("Welcome at Vending Machine")
	for {
		breakOperation := selectUser()
		if breakOperation {
			break
		}
	}
}

var selectUser = func() bool {
	var input int
	var err error
	fmt.Println("Enter 1 for Supplier, 2 for User and 3 for Exit")
	input, err = utils.GetUserInput()
	if err != nil {
		if err != io.EOF {
			log.Printf("Runtime error occured : %v \n", err)
		}
		return true
	}
	switch input {
	case 1:
		supplier.AccessMachine()
	case 2:
		consumer.SelectProduct()
	case 3:
		return true
	default:
		fmt.Println("Invalid Input. Please try again")
	}
	return false
}
