package supplier

import (
	"VendingMachine/config"
	"VendingMachine/constant"
	"VendingMachine/product"
	"VendingMachine/utils"
	"fmt"
	"io"
)

var (
	fmtPrintf  = fmt.Printf
	fmtPrintln = fmt.Println
)

//AccessMachine supplier flow
var AccessMachine = func() {
	fmt.Println("Press 1 to reset Vending Machine, 2 for checking available quantity, 3 for Total Amount Collected")
	input, err := utils.GetUserInput()
	if err != nil {
		if err != io.EOF {
			fmtPrintf("Runtime error occured : %v \n", err)
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
		fmtPrintf("Total Amount Collected : %d \n", constant.TotalAmountCollected)
	default:
		fmtPrintln("Invalid Input. Please try again")
	}
}

var checkAvailanilityQuantity = func() {
	for _, products := range config.Config.Products {
		fmtPrintf("Product %s has %d quantity remaining \n", products.Name, constant.ProductAvailabilityMap[products.Name])
	}
}
