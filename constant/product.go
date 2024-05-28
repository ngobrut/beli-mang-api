package constant

type ProductCategory string

const (
	Beverage   ProductCategory = "Beverage"
	Food       ProductCategory = "Food"
	Snack      ProductCategory = "Snack"
	Condiments ProductCategory = "Condiments"
	Additions  ProductCategory = "Additions"
)

var ValidProductCategory = map[string]bool{
	string(Beverage):   true,
	string(Food):       true,
	string(Snack):      true,
	string(Condiments): true,
	string(Additions):  true,
}

var ProductCategories = []string{
	string(Beverage),
	string(Food),
	string(Snack),
	string(Condiments),
	string(Additions),
}
