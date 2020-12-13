package product

import (
	"VendingMachine/config"
	"VendingMachine/constant"
	"testing"
)

func TestConvertProductToMap(t *testing.T) {
	sampleProduct := config.ProductList{
		Name:            "test",
		Price:           10,
		DefaultQuantity: 10,
	}
	config.Config.Products = append(config.Config.Products, sampleProduct)
	t.Run("ConvertProductToMap", func(t *testing.T) {
		ConvertProductToMap()
		if len(constant.ProductPriceMap) != len(config.Config.Products) || len(constant.ProductAvailabilityMap) != len(config.Config.Products) {
			t.Errorf("Expecting length of products price map as 1 and got %d. Expecting length of products Availability map as 1 and got %d", len(constant.ProductPriceMap), len(constant.ProductAvailabilityMap))
		}
	})
}
