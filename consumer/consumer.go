package consumer

import (
	"VendingMachine/constant"
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

//Consume consumer flow
/* var Consume = func() {
	available := isAnyProductAvailable()
	if !available {
		fmtPrintln("No products available at this moment.")
		return
	}
	fmt.Println("Press 1 to Buy")
	input, err := utils.GetUserInput()
	if err != nil {
		if err != io.EOF {
			fmtPrintf("Runtime error occured : %v \n", err)
		}
		return
	}
	switch input {
	case 1:
		SelectProduct()
	case 2:
	returnProduct()
	default:
		fmtPrintln("Invalid Input. Please try again")
	}
} */

var isAnyProductAvailable = func() bool {
	for key := range constant.ProductAvailabilityMap {
		if constant.ProductAvailabilityMap[key] > 0 {
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
		if constant.ProductAvailabilityMap["Coke"] == 0 {
			fmtPrintln("Coke is not available at this moment. Please try another drink.")
			return false
		}
		fmtPrintln("Processing purchase for Coke")
		return confirmPurchase("Coke")
	case 2:
		if constant.ProductAvailabilityMap["Pepsi"] == 0 {
			fmtPrintln("Pepsi is not available at this moment. Please try another drink.")
			return false
		}
		fmtPrintln("Processing purchase for Pepsi")
		return confirmPurchase("Pepsi")
	case 3:
		if constant.ProductAvailabilityMap["Soda"] == 0 {
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

/*
func returnProduct() {
	var input int
	fmt.Println("Press 1 for Coke, 2 for Pepsi, 3 for Soda or 4 for exit")
	_, err := fmt.Scanln(&input)
	if err != nil {
		if err != io.EOF {
			log.Printf("Runtime error occured : %v \n", err)
		}
		return
	}
	switch input {
	case 1:
		processReturnProduct("Coke")
	case 2:
		processReturnProduct("Pepsi")
	case 3:
		processReturnProduct("Soda")
	case 4:
		return
	default:
		fmt.Println("Invalid Input. Please try again")
		returnProduct()
	}
} */

/* func processReturnProduct(productName string) {
	constant.ProductAvailabilityMap[productName] = constant.ProductAvailabilityMap[productName] + 1
	constant.TotalAmountCollected = constant.TotalAmountCollected - constant.ProductPriceMap[productName]
	fmt.Printf("Please collect %d amount \n", constant.ProductPriceMap[productName])
} */

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
	priceOfProduct := constant.ProductPriceMap[productName]
	breakLoop := getCoins(priceOfProduct)
	constant.ProductAvailabilityMap[productName] = constant.ProductAvailabilityMap[productName] - 1
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
		if constant.AcceptedCoinsMap[amountReceived] {
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
	constant.TotalAmountCollected += amount
	return false
}
