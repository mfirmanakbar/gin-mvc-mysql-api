package products

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

// the POJO for internal users only with a fake HEADER as go-access-token
type PrivateProduct struct {
	Id            int64           `json:"id"`
	Name          string          `json:"name"`
	Code          string          `json:"code"`
	SellPrice     decimal.Decimal `json:"sell_price" sql:"type:decimal(50,6)"`
	BuyPrice      decimal.Decimal `json:"buy_price" sql:"type:decimal(50,6)"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
	DeletedAt     string          `json:"deleted_at"`
	ProductBundle int             `json:"product_bundle"`
}

// the POJO for external users only with no HEADER
type PublicProduct struct {
	Name          string          `json:"name"`
	Code          string          `json:"code"`
	Price         decimal.Decimal `json:"price" sql:"type:decimal(50,6)"`
	ProductBundle int             `json:"product_bundle"`
}

func (product *Product) Marshall(isPublic bool) interface{} {
	if isPublic {
		// if there's custom field then using manual mapping
		// e.g i have `sell_price` on DB but i want to show it like `price` on the JSON response
		return PublicProduct{
			Name:          product.Name,
			Code:          product.Code,
			Price:         product.SellPrice,
			ProductBundle: product.ProductBundle,
		}
	} else {
		// if there's no custom field between original POJO and this marshaller
		// we can using json.Marshal and json.Unmarshal
		productJson, _ := json.Marshal(product)
		var privateProduct PrivateProduct
		json.Unmarshal(productJson, &privateProduct)
		return privateProduct
	}
}
