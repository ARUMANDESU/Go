package coffeeco

import "github.com/Rhymond/go-money"

type ProductSize string

const (
	PRODUCT_SIZE_SMALL  ProductSize = "small"
	PRODUCT_SIZE_MEDIUM ProductSize = "medium"
	PRODUCT_SIZE_LARGE  ProductSize = "large"
)

type Product struct {
	ItemName  string
	BasePrice money.Money
	Size      ProductSize
}
