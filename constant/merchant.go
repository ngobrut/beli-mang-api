package constant

type MerchantCategory string

const (
	SmallRestaurant       MerchantCategory = "SmallRestaurant"
	MediumRestaurant      MerchantCategory = "MediumRestaurant"
	LargeRestaurant       MerchantCategory = "LargeRestaurant"
	MerchandiseRestaurant MerchantCategory = "MerchandiseRestaurant"
	BoothKiosk            MerchantCategory = "BoothKiosk"
	ConvenienceStore      MerchantCategory = "ConvenienceStore"
)

var ValidMerchantCategory = map[string]bool{
	string(SmallRestaurant):       true,
	string(MediumRestaurant):      true,
	string(LargeRestaurant):       true,
	string(MerchandiseRestaurant): true,
	string(BoothKiosk):            true,
	string(ConvenienceStore):      true,
}

var MerchantCategories = []string{
	string(SmallRestaurant),
	string(MediumRestaurant),
	string(LargeRestaurant),
	string(MerchandiseRestaurant),
	string(BoothKiosk),
	string(ConvenienceStore),
}
