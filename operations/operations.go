package operations

import (
	"VendingMachine/config"
	"VendingMachine/constant"
	"VendingMachine/product"
	"fmt"
	"log"
)

//Start : Starts Vending Machine
func Start() {
	product.ConvertProductToMap()
	var input int
	for {
		fmt.Println("Welcome at Vending Machine")
		fmt.Println("Enter 1 for Supplier and 2 for User")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Printf("Runtime error occured : %v \n", err)
			return
		}
		switch input {
		case 1:
			supplier()
		case 2:
			consumer()
		default:
			fmt.Println("Invalid Input. Please try again")
			continue
		}
	}
}

func consumer() {
	var input int
	fmt.Println("Press 1 to Buy, 2 for Return or 3 for Exit")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Printf("Runtime error occured : %v \n", err)
		return
	}
	switch input {
	case 1:
		selectProduct()
	case 2:
		returnProduct()
	case 3:
		return
	default:
		fmt.Println("Invalid Input. Please try again")
	}
	consumer()
}

func selectProduct() {
	var input int
	fmt.Println("Press 1 for Coke, 2 for Pepsi, 3 for Soda or 4 for exit")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Printf("Runtime error occured : %v \n", err)
		return
	}
	switch input {
	case 1:
		confirmPurchase("Coke")
	case 2:
		confirmPurchase("Pepsi")
	case 3:
		confirmPurchase("Soda")
	case 4:
		return
	default:
		fmt.Println("Invalid Input. Please try again")
		selectProduct()
	}
}

func returnProduct() {
	var input int
	fmt.Println("Press 1 for Coke, 2 for Pepsi, 3 for Soda or 4 for exit")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Printf("Runtime error occured : %v \n", err)
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
}

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
		log.Printf("Runtime error occured : %v \n", err)
		return
	}
	switch input {
	case 1:
		purchaseSuccess := BuyProduct(productName)
		if purchaseSuccess {
			fmt.Println("Purchase Successful")
			return
		}
		fmt.Println("Purchase Failed. Please try again")
		confirmPurchase(productName)
	case 2:
		return
	default:
		fmt.Println("Invalid Input. Please try again")
		confirmPurchase(productName)
	}
}

func supplier() {
	var input int
	fmt.Println("Press 1 to reset Vending Machine, 2 for checking available quantity, 3 for Total Amount Collected and 4 for Exit")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Printf("Runtime error occured : %v \n", err)
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
	supplier()
}

func checkAvailanilityQuantity() {
	for _, products := range config.Config.Products {
		fmt.Printf("Product %s has %d quantity remaining \n", products.Name, constant.ProductAvailabilityMap[products.Name])
	}
}

//BuyProduct : steps to perform when purchase is confirmed
func BuyProduct(productName string) bool {
	priceOfProduct := constant.ProductPriceMap[productName]
	success := GetCoins(priceOfProduct)
	constant.ProductAvailabilityMap[productName] = constant.ProductAvailabilityMap[productName] - 1
	return success
}

//GetCoins received given amount
func GetCoins(amount int) bool {
	fmt.Printf("Please enter coins. Total Amount : %d \n", amount)
	totalAmountReceived := 0
	var amountReceived int
	for totalAmountReceived < amount {
		_, err := fmt.Scanln(&amountReceived)
		if err != nil {
			log.Printf("Runtime error occured : %v \n", err)
			return GetCoins(amount)
		}
		if constant.AcceptedCoinsMap[amountReceived] {
			totalAmountReceived += amountReceived
			if totalAmountReceived < amount {
				fmt.Printf("Please add more %d", amount-totalAmountReceived)
			}
			continue
		}
		fmt.Println("Please enter coins of 1,5,10 Or 25")
	}
	if totalAmountReceived > amount {
		fmt.Printf("Please collect change of %d \n", totalAmountReceived-amount)
	}
	constant.TotalAmountCollected += amount
	return true
}
