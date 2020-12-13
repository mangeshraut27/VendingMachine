package supplier

import (
	"VendingMachine/config"
	"VendingMachine/constant"
	"VendingMachine/product"
	"fmt"
	"io"
	"log"
)

//AccessMachine supplier flow
func AccessMachine() {
	var input int
	fmt.Println("Press 1 to reset Vending Machine, 2 for checking available quantity, 3 for Total Amount Collected and 4 for Exit")
	_, err := fmt.Scanln(&input)
	if err != nil {
		if err != io.EOF {
			log.Printf("Runtime error occured : %v \n", err)
		}
		return
	}

	switch input {
	case 1:
		product.ConvertProductToMap()
		constant.TotalAmountCollected = 0
	case 2:
		checkAvailanilityQuantity()
	case 3:
		fmt.Printf("Total Amount Collected : %d \n", constant.TotalAmountCollected)
	case 4:
		return
	default:
		fmt.Println("Invalid Input. Please try again")
	}
	AccessMachine()
}

func checkAvailanilityQuantity() {
	for _, products := range config.Config.Products {
		fmt.Printf("Product %s has %d quantity remaining \n", products.Name, constant.ProductAvailabilityMap[products.Name])
	}
}
