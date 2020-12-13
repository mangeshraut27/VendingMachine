package consumer

import (
	"VendingMachine/global"
	"VendingMachine/utils"
	"fmt"
	"io"
)

var (
	//TotalAmountReceived in vending machine
	TotalAmountReceived = 0
	fmtPrintf           = fmt.Printf
	fmtPrintln          = fmt.Println
)

var isAnyProductAvailable = func() bool {
	for key := range global.ProductAvailabilityMap {
		if global.ProductAvailabilityMap[key] > 0 {
			return true
		}
	}
	return false
}

//SelectProduct Select product
var SelectProduct = func() bool {
	var input int
	available := isAnyProductAvailable()
	if !available {
		fmtPrintln("No products available at this moment.")
		return false
	}
	fmt.Println("Press 1 for Coke, 2 for Pepsi, 3 for Soda")
	input, err := utils.GetUserInput()
	if err != nil {
		if err != io.EOF {
			fmtPrintf("Runtime error occured : %v \n", err)
		}
		return true
	}
	switch input {
	case 1:
		if global.ProductAvailabilityMap["Coke"] == 0 {
			fmtPrintln("Coke is not available at this moment. Please try another drink.")
			return false
		}
		fmtPrintln("Processing purchase for Coke")
		return confirmPurchase("Coke")
	case 2:
		if global.ProductAvailabilityMap["Pepsi"] == 0 {
			fmtPrintln("Pepsi is not available at this moment. Please try another drink.")
			return false
		}
		fmtPrintln("Processing purchase for Pepsi")
		return confirmPurchase("Pepsi")
	case 3:
		if global.ProductAvailabilityMap["Soda"] == 0 {
			fmtPrintln("Soda is not available at this moment. Please try another drink.")
			return false
		}
		fmtPrintln("Processing purchase for Soda")
		return confirmPurchase("Soda")
	default:
		fmtPrintln("Invalid Input. Please try again")
	}
	return false
}

var confirmPurchase = func(productName string) bool {
	var input int
	fmt.Println("Press 1 to Confirm or 2 for Cancel")
	input, err := utils.GetUserInput()
	if err != nil {
		if err != io.EOF {
			fmtPrintf("Runtime error occured : %v \n", err)
		}
		return true
	}
	switch input {
	case 1:
		fmtPrintln("Purchase confirmed")
		return buyProduct(productName)
	case 2:
		fmtPrintln("Purchase cancelled")
	default:
		fmtPrintln("Invalid Input. Please try again")
	}
	return false
}

//buyProduct : steps to perform when purchase is confirmed
var buyProduct = func(productName string) bool {
	priceOfProduct := global.ProductPriceMap[productName]
	breakLoop := getCoins(priceOfProduct)
	if !breakLoop {
		fmt.Printf("Please collect %s \n", productName)
		global.ProductAvailabilityMap[productName] = global.ProductAvailabilityMap[productName] - 1
	}
	return breakLoop
}

//getCoins received given amount
var getCoins = func(amount int) bool {
	fmt.Printf("Please enter coins. Total Amount : %d \n", amount)
	var amountReceived int
	var err error
	for TotalAmountReceived < amount {
		amountReceived, err = utils.GetUserInput()
		if err != nil {
			if err != io.EOF {
				fmtPrintf("Runtime error occured : %v. Please try again \n", err)
				fmt.Printf("Please Collect %d coins", TotalAmountReceived)
				TotalAmountReceived = 0
			}
			return true
		}
		if global.AcceptedCoinsMap[amountReceived] {
			TotalAmountReceived += amountReceived
			if TotalAmountReceived < amount {
				fmtPrintf("Please add more %d \n", amount-TotalAmountReceived)
			}
			continue
		}
		fmtPrintln("Please enter coins of 1,5,10 Or 25")
	}
	fmt.Println("Purchase Successful")
	if TotalAmountReceived > amount {
		fmtPrintf("Please collect change of %d \n", TotalAmountReceived-amount)
	}
	TotalAmountReceived = 0
	global.TotalAmountCollected += amount
	return false
}
