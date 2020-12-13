package supplier

import (
	"VendingMachine/config"
	"VendingMachine/constant"
	"VendingMachine/product"
	"VendingMachine/utils"
	"errors"
	"fmt"
	"testing"
)

var (
	errTest      = errors.New("Test Error")
	printMessage string
)

var ConvertProductToMapMock = func() {
	return
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

func TestAccessMachine(t *testing.T) {
	oldGetUserInput := utils.GetUserInput
	oldConvertProductMap := product.ConvertProductToMap
	oldFmtPrintf := fmtPrintf
	oldFmtPrintln := fmtPrintln

	defer func() {
		utils.GetUserInput = oldGetUserInput
		product.ConvertProductToMap = oldConvertProductMap
		fmtPrintln = oldFmtPrintln
		fmtPrintf = oldFmtPrintf
	}()
	t.Run("Reset Vending Machine", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedOneMock
		constant.ProductAvailabilityMap["Coke"] = 9
		sampleProduct := config.ProductList{
			Name:            "Coke",
			Price:           10,
			DefaultQuantity: 10,
		}
		config.Config.Products = append(config.Config.Products, sampleProduct)
		AccessMachine()
		if constant.ProductAvailabilityMap["Coke"] != 10 {
			t.Errorf("Expecting availability 10. received %d", constant.ProductAvailabilityMap["Coke"])
		}
	})
	t.Run("Check Product Availability", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTwoMock
		fmtPrintf = fmtPrintfMock
		config.Config.Products = []config.ProductList{}
		sampleProduct := config.ProductList{
			Name:            "Coke",
			Price:           10,
			DefaultQuantity: 10,
		}
		config.Config.Products = append(config.Config.Products, sampleProduct)
		constant.ProductAvailabilityMap["Coke"] = 10
		AccessMachine()
		if printMessage != "Product Coke has 10 quantity remaining \n" {
			t.Errorf("Expected %s Print Statement and received : %s", "'Product Coke has 10 quantity remaining'", printMessage)
		}
	})
	t.Run("Print total amount collected", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedThreeMock
		fmtPrintf = fmtPrintfMock
		constant.TotalAmountCollected = 45
		AccessMachine()
		if printMessage != "Total Amount Collected : 45 \n" {
			t.Errorf("Expected %s Print Statement and received : %s", "'Total Amount Collected : 45 \n'", printMessage)
		}
	})
	t.Run("default", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedFourMock
		fmtPrintln = fmtPrintlnMock
		AccessMachine()
		if printMessage != fmt.Sprintln("Invalid Input. Please try again") {
			t.Errorf("Expected %s Print Statement and received : %s", "'Invalid Input. Please try again'", printMessage)
		}
	})
	t.Run("Error", func(t *testing.T) {
		utils.GetUserInput = GetUserInputMockError
		fmtPrintf = fmtPrintfMock
		AccessMachine()
		if printMessage != "Runtime error occured : Test Error \n" {
			t.Errorf("Expected %s Print Statement and received : %s", "'Runtime error occured : Test Error \n'", printMessage)
		}
	})
}
