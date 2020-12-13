package consumer

import (
	"VendingMachine/constant"
	"VendingMachine/utils"
	"fmt"
	"io"
	"log"
)

var (
	TotalAmountReceived = 0
)

//Consume consumer flow
func Consume() {
	available := isAnyProductAvailable()
	if !available {
		fmt.Println("No products available at this moment.")
		return
	}
	fmt.Println("Press 1 to Buy or 2 for Exit")
	input, err := utils.GetUserInput()
	if err != nil {
		if err != io.EOF {
			log.Printf("Runtime error occured : %v \n", err)
		}
		return
	}
	switch input {
	case 1:
		selectProduct()
	/* case 2:
	returnProduct() */
	case 2:
		return
	default:
		fmt.Println("Invalid Input. Please try again")
	}
	Consume()
}

func isAnyProductAvailable() bool {
	for key := range constant.ProductAvailabilityMap {
		if constant.ProductAvailabilityMap[key] > 0 {
			return true
		}
	}
	return false
}

func selectProduct() {
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
		if constant.ProductAvailabilityMap["Coke"] == 0 {
			fmt.Println("Coke is not available at this moment. Please try another drink.")
			selectProduct()
		}
		confirmPurchase("Coke")
	case 2:
		if constant.ProductAvailabilityMap["Pepsi"] == 0 {
			fmt.Println("Pepsi is not available at this moment. Please try another drink.")
			selectProduct()
		}
		confirmPurchase("Pepsi")
	case 3:
		if constant.ProductAvailabilityMap["Soda"] == 0 {
			fmt.Println("Soda is not available at this moment. Please try another drink.")
			selectProduct()
		}
		confirmPurchase("Soda")
	case 4:
		return
	default:
		fmt.Println("Invalid Input. Please try again")
		selectProduct()
	}
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

func processReturnProduct(productName string) {
	constant.ProductAvailabilityMap[productName] = constant.ProductAvailabilityMap[productName] + 1
	constant.TotalAmountCollected = constant.TotalAmountCollected - constant.ProductPriceMap[productName]
	fmt.Printf("Please collect %d amount \n", constant.ProductPriceMap[productName])
}

func confirmPurchase(productName string) {
	var input int
	fmt.Println("Press 1 to Confirm or 2 for Exit")
	_, err := fmt.Scanln(&input)
	if err != nil {
		if err != io.EOF {
			log.Printf("Runtime error occured : %v \n", err)
		}
		return
	}
	switch input {
	case 1:
		buyProduct(productName)
	case 2:
		return
	default:
		fmt.Println("Invalid Input. Please try again")
		confirmPurchase(productName)
	}
}

//buyProduct : steps to perform when purchase is confirmed
func buyProduct(productName string) {
	priceOfProduct := constant.ProductPriceMap[productName]
	getCoins(priceOfProduct)
	constant.ProductAvailabilityMap[productName] = constant.ProductAvailabilityMap[productName] - 1
}

//getCoins received given amount
func getCoins(amount int) {
	fmt.Printf("Please enter coins. Total Amount : %d \n", amount)
	var amountReceived int
	for TotalAmountReceived < amount {
		_, err := fmt.Scanln(&amountReceived)
		if err != nil {
			if err != io.EOF {
				log.Printf("Runtime error occured : %v. Please try again \n", err)
				fmt.Printf("Please Collect %d coins", TotalAmountReceived)
				TotalAmountReceived = 0
				getCoins(amount)
			}
			return
		}
		if constant.AcceptedCoinsMap[amountReceived] {
			TotalAmountReceived += amountReceived
			if TotalAmountReceived < amount {
				fmt.Printf("Please add more %d \n", amount-TotalAmountReceived)
			}
			continue
		}
		fmt.Println("Please enter coins of 1,5,10 Or 25")
	}
	fmt.Println("Purchase Successful")
	if TotalAmountReceived > amount {
		fmt.Printf("Please collect change of %d \n", TotalAmountReceived-amount)
	}
	TotalAmountReceived = 0
	constant.TotalAmountCollected += amount
}
