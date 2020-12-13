package product

import (
	"VendingMachine/config"
	"VendingMachine/constant"
)

//ConvertProductToMap convert product list to map of product name:price name:availability
var ConvertProductToMap = func() {
	for _, product := range config.Config.Products {
		constant.ProductPriceMap[product.Name] = product.Price
		constant.ProductAvailabilityMap[product.Name] = product.DefaultQuantity
	}
}
