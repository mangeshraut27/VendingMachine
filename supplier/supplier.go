package supplier

import (
	"VendingMachine/config"
	"VendingMachine/global"
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
var AccessMachine = func() bool {
	fmt.Println("Press 1 to reset Vending Machine, 2 for checking available quantity, 3 for Total Amount Collected")
	input, err := utils.GetUserInput()
	if err != nil {
		if err != io.EOF {
			fmtPrintf("Runtime error occured : %v \n", err)
		}
		return true
	}

	switch input {
	case 1:
		product.ConvertProductToMap()
		global.TotalAmountCollected = 0
	case 2:
		checkAvailanilityQuantity()
	case 3:
		fmtPrintf("Total Amount Collected : %d \n", global.TotalAmountCollected)
	default:
		fmtPrintln("Invalid Input. Please try again")
	}
	return false
}

var checkAvailanilityQuantity = func() {
	for _, products := range config.Config.Products {
		fmtPrintf("Product %s has %d quantity remaining \n", products.Name, global.ProductAvailabilityMap[products.Name])
	}
}
