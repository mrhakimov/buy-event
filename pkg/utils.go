package pkg

import "fmt"

// Customer's info
type Customer struct {
	ID    int
	phone string
	email string
}

// Customer's single item from purchase
type Item struct {
	ID      int
	product string
	price   float64
}

// SamplePurchase is a sample list of items
var SamplePurchase = &[]Item{
	{1, "pillow", 13.59},
	{2, "brush", 3.1},
	{3, "lamp", 34.5},
	{4, "soap", 1.2},
	{5, "stove", 499.99},
	{6, "sofa", 1532},
	{7, "chair", 899.99},
}

// ShowPurchase converts list of items to string
func ShowPurchase(purchases []Item) (result string) {
	var total float64 = 0
	result += "\n"

	for _, p := range purchases {
		result += fmt.Sprintf("%v\t-\t%.2f\n", p.product, p.price)
		total += p.price
	}

	result += "============\n"
	result += fmt.Sprintf("Total: %.2f", total)

	return result
}
