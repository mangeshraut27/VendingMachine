package global

var (
	//AcceptedCoinsMap is map of coins that our machine accepts
	AcceptedCoinsMap = map[int]bool{
		1:  true,
		5:  true,
		10: true,
		25: true,
	}

	//ConfigFilePath config file path
	ConfigFilePath = "config/config.json"

	//ProductAvailabilityMap : Map to store product and available quantity
	ProductAvailabilityMap = make(map[string]int)

	//ProductPriceMap : Map to store product and price map
	ProductPriceMap = make(map[string]int)

	//TotalAmountCollected : TotalAmountCollected in Vending Machine
	TotalAmountCollected = 0
)
