package operations

import (
	"VendingMachine/consumer"
	"VendingMachine/product"
	"VendingMachine/supplier"
	"VendingMachine/utils"
	"errors"
	"testing"
)

var (
	errTest = errors.New("Test Error")
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

var SelectProductMock = func() {
	return
}
var AccessMachineMock = func() {
	return
}

var GetUserInputMockError = func() (int, error) {
	return 0, errTest
}

var selectUserMock = func() bool {
	return true
}

func TestStart(t *testing.T) {
	oldConvertProductMap := product.ConvertProductToMap
	oldSelectUser := selectUser
	defer func() {
		product.ConvertProductToMap = oldConvertProductMap
		selectUser = oldSelectUser
	}()
	t.Run("Start", func(t *testing.T) {
		product.ConvertProductToMap = ConvertProductToMapMock
		selectUser = selectUserMock
		Start()
	})
}

func TestSelectUser(t *testing.T) {
	oldGetUserInput := utils.GetUserInput
	oldAccessMachine := supplier.AccessMachine
	oldSelectProduct := consumer.SelectProduct

	defer func() {
		utils.GetUserInput = oldGetUserInput
		supplier.AccessMachine = oldAccessMachine
		consumer.SelectProduct = oldSelectProduct
	}()

	t.Run("Supplier", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedOneMock
		supplier.AccessMachine = AccessMachineMock
		isBreak := selectUser()
		if isBreak {
			t.Error("Expecting false and received true")
		}
	})
	t.Run("Consumer", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedTwoMock
		consumer.SelectProduct = SelectProductMock
		isBreak := selectUser()
		if isBreak {
			t.Error("Expecting false and received true")
		}
	})
	t.Run("Exit", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedThreeMock
		isBreak := selectUser()
		if !isBreak {
			t.Error("Expecting true and received false")
		}
	})
	t.Run("Default", func(t *testing.T) {
		utils.GetUserInput = GetUserInputPressedFourMock
		isBreak := selectUser()
		if isBreak {
			t.Error("Expecting false and received true")
		}
	})
	t.Run("Error", func(t *testing.T) {
		utils.GetUserInput = GetUserInputMockError
		isBreak := selectUser()
		if !isBreak {
			t.Error("Expecting true and received false")
		}
	})
}
