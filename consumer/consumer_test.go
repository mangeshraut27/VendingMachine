package consumer

import (
	"VendingMachine/constant"
	"VendingMachine/utils"
	"errors"
	"fmt"
	"testing"
	"time"
)

var (
	errTest      = errors.New("Test Error")
	printMessage string
)

var isAnyProductAvailableMock = func() bool {
	return true
}
var isAnyProductAvailableMockFalse = func() bool {
	return false
}

var GetUserInputPressedOneMock = func() (int, error) {
	return 1, nil
}
var GetUserInputPressedTwoMock = func() (int, error) {
	return 2, nil
}
var GetUserInputPressedThreeMock = func() (int, error) {
	return 3, nil
}
var GetUserInputPressedFourMock = func() (int, error) {
	return 4, nil
}
var GetUserInputPressedFiveMock = func() (int, error) {
	return 5, nil
}
var GetUserInputPressedTenMock = func() (int, error) {
	return 10, nil
}
var GetUserInputPressedTwentyFiveMock = func() (int, error) {
	return 25, nil
}
var GetUserInputMockError = func() (int, error) {
	return 0, errTest
}

var fmtPrintfMock = func(format string, msg ...interface{}) (n int, err error) {
	printMessage = fmt.Sprintf(format, msg...)
	return 0, nil
}

var fmtPrintlnMock = func(a ...interface{}) (n int, err error) {
	printMessage = fmt.Sprintln(a...)
	return 0, nil
}

var confirmPurchaseMock = func(product string) bool {
	return false
}

var buyProductMock = func(product string) bool {
	return false
}

var getCoinsMock = func(amount int) bool {
	return false
}

func TestSelectProduct(t *testing.T) {
	oldIsAnyProductAvailable := isAnyProductAvailable
	oldConfirmPurchase := confirmPurchase
	oldGetUserInput := utils.GetUserInput

	defer func() {
		isAnyProductAvailable = oldIsAnyProductAvailable
		confirmPurchase = oldConfirmPurchase
		utils.GetUserInput = oldGetUserInput
	}()
	t.Run("isAnyProductAvailable false", func(t *testing.T) {
		isAnyProductAvailable = isAnyProductAvailableMockFalse
		fmtPrintln = fmtPrintlnMock
		SelectProduct()
		if printMessage != fmt.Sprintln("No products available at this moment.") {
			t.Errorf("Expected %s Print Statement and received : %s", "'No products available at this moment.'", printMessage)
		}
	})
	t.Run("Error", func(t *testing.T) {
		utils.GetUserInput = GetUserInputMockError
		isAnyProductAvailable = isAnyProductAvailableMock
		fmtPrintf = fmtPrintfMock
		SelectProduct()
		if printMessage != "Runtime error occured : Test Error \n" {
			t.Errorf("Expected %s Print Statement and received : %s", "'Runtime error occured : Test Error \n'", printMessage)
		}
	})
	t.Run("Coke unavailable", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedOneMock
		fmtPrintln = fmtPrintlnMock
		constant.ProductAvailabilityMap["Coke"] = 0
		SelectProduct()
		if printMessage != fmt.Sprintln("Coke is not available at this moment. Please try another drink.") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Coke is not available at this moment. Please try another drink.'", printMessage)
		}
	})
	t.Run("Coke available", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedOneMock
		fmtPrintln = fmtPrintlnMock
		confirmPurchase = confirmPurchaseMock
		constant.ProductAvailabilityMap["Coke"] = 10
		SelectProduct()
		if printMessage != fmt.Sprintln("Processing purchase for Coke") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Processing purchase for Coke'", printMessage)
		}
	})
	t.Run("Pepsi unavailable", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTwoMock
		fmtPrintln = fmtPrintlnMock
		constant.ProductAvailabilityMap["Pepsi"] = 0
		SelectProduct()
		if printMessage != fmt.Sprintln("Pepsi is not available at this moment. Please try another drink.") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Pepsi is not available at this moment. Please try another drink.'", printMessage)
		}
	})
	t.Run("Pepsi available", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTwoMock
		fmtPrintln = fmtPrintlnMock
		confirmPurchase = confirmPurchaseMock
		constant.ProductAvailabilityMap["Pepsi"] = 10
		SelectProduct()
		if printMessage != fmt.Sprintln("Processing purchase for Pepsi") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Processing purchase for Pepsi'", printMessage)
		}
	})
	t.Run("Soda unavailable", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedThreeMock
		fmtPrintln = fmtPrintlnMock
		constant.ProductAvailabilityMap["Soda"] = 0
		SelectProduct()
		if printMessage != fmt.Sprintln("Soda is not available at this moment. Please try another drink.") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Soda is not available at this moment. Please try another drink.'", printMessage)
		}
	})
	t.Run("Soda available", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedThreeMock
		fmtPrintln = fmtPrintlnMock
		confirmPurchase = confirmPurchaseMock
		constant.ProductAvailabilityMap["Soda"] = 10
		SelectProduct()
		if printMessage != fmt.Sprintln("Processing purchase for Soda") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Processing purchase for Soda'", printMessage)
		}
	})

	t.Run("default", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedFourMock
		fmtPrintln = fmtPrintlnMock
		SelectProduct()
		if printMessage != fmt.Sprintln("Invalid Input. Please try again") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Invalid Input. Please try again'", printMessage)
		}
	})
}

func TestIsAnyProductAvailable(t *testing.T) {
	t.Run("Return true", func(t *testing.T) {
		constant.ProductAvailabilityMap["Coke"] = 1
		availability := isAnyProductAvailable()
		if !availability {
			t.Error("Expecting true but returns false")
		}
	})
	t.Run("Return false", func(t *testing.T) {
		constant.ProductAvailabilityMap["Coke"] = 0
		constant.ProductAvailabilityMap["Pepsi"] = 0
		constant.ProductAvailabilityMap["Soda"] = 0
		availability := isAnyProductAvailable()
		if availability {
			t.Error("Expecting false but returns true")
		}
	})
}

func TestConfirmPurchase(t *testing.T) {
	oldGetUserInput := utils.GetUserInput
	oldBuyProduct := buyProduct
	defer func() {
		buyProduct = oldBuyProduct
		utils.GetUserInput = oldGetUserInput
	}()
	t.Run("Error", func(t *testing.T) {
		utils.GetUserInput = GetUserInputMockError
		fmtPrintf = fmtPrintfMock
		confirmPurchase("Coke")
		if printMessage != "Runtime error occured : Test Error \n" {
			t.Errorf("Expected %s Print Statement and received : %s", "'Runtime error occured : Test Error \n'", printMessage)
		}
	})
	t.Run("Confirm", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedOneMock
		fmtPrintln = fmtPrintlnMock
		buyProduct = buyProductMock
		confirmPurchase("Coke")
		if printMessage != fmt.Sprintln("Purchase confirmed") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Purchase confirmed'", printMessage)
		}
	})
	t.Run("Confirm", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTwoMock
		fmtPrintln = fmtPrintlnMock
		confirmPurchase("Coke")
		if printMessage != fmt.Sprintln("Purchase cancelled") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Purchase cancelled'", printMessage)
		}
	})
	t.Run("Default", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedThreeMock
		fmtPrintln = fmtPrintlnMock
		confirmPurchase("Coke")
		if printMessage != fmt.Sprintln("Invalid Input. Please try again") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Invalid Input. Please try again'", printMessage)
		}
	})
}

func TestBuyProduct(t *testing.T) {
	oldgetCoins := getCoins

	defer func() {
		getCoins = oldgetCoins
	}()
	t.Run("Product Purchased", func(t *testing.T) {
		getCoins = getCoinsMock
		constant.ProductAvailabilityMap["Coke"] = 2
		buyProduct("Coke")
		if constant.ProductAvailabilityMap["Coke"] != 1 {
			t.Errorf("Expecting product quantity reduction by 1 but not reduced. Current Quantity : %d \n", constant.ProductAvailabilityMap["Coke"])
		}
	})
}

func TestGetCoins(t *testing.T) {
	oldGetUserInput := utils.GetUserInput
	defer func() {
		utils.GetUserInput = oldGetUserInput
	}()
	t.Run("Error", func(t *testing.T) {
		utils.GetUserInput = GetUserInputMockError
		fmtPrintf = fmtPrintfMock
		TotalAmountReceived = 0
		getCoins(25)
		if printMessage != "Runtime error occured : Test Error. Please try again \n" {
			t.Errorf("Expected %s Print Statement and received : %s", "'Runtime error occured : Test Error. Please try again \n'", printMessage)
		}
	})
	t.Run("Invalid Coin Entered", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTwoMock
		fmtPrintln = fmtPrintlnMock
		TotalAmountReceived = 0
		ticker := time.NewTicker(2 * time.Second)
		go func(ticker *time.Ticker) {
			for {
				select {
				case <-ticker.C:
					utils.GetUserInput = GetUserInputPressedTwentyFiveMock
				}
			}
		}(ticker)
		getCoins(25)
		if printMessage != fmt.Sprintln("Please enter coins of 1,5,10 Or 25") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Please enter coins of 1,5,10 Or 25'", printMessage)
		}
	})
	t.Run("Amount needed 25", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTwentyFiveMock
		fmtPrintf = fmtPrintfMock
		constant.TotalAmountCollected = 0
		getCoins(25)
		if constant.TotalAmountCollected != 25 {
			t.Errorf("Expecting constant.TotalAmountCollected as 25. received : %d \n", constant.TotalAmountCollected)
		}
	})
	t.Run("Amount needed 10", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTenMock
		fmtPrintf = fmtPrintfMock
		getCoins(20)
		if printMessage != "Please add more 10 \n" {
			t.Errorf("Expected %s Print Statement and received : %s", "'Please add more 10 \n'", printMessage)
		}
	})
	t.Run("Return Change", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTenMock
		fmtPrintf = fmtPrintfMock
		getCoins(15)
		if printMessage != "Please collect change of 5 \n" {
			t.Errorf("Expected %s Print Statement and received : %s", "'Please collect change of 5 \n'", printMessage)
		}
	})
}
