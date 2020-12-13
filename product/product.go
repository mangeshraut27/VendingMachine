package product

import (
	"VendingMachine/config"
	"VendingMachine/global"
)

//ConvertProductToMap convert product list to map of product name:price name:availability
var ConvertProductToMap = func() {
	for _, product := range config.Config.Products {
		global.ProductPriceMap[product.Name] = product.Price
		global.ProductAvailabilityMap[product.Name] = product.DefaultQuantity
	}
}
