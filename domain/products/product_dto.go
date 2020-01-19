package products

import (
	"github.com/mfirmanakbar/gin-mvc-mysql-api/utils/errors"
	"github.com/shopspring/decimal"
	"strings"
)

const (
	ProductBundleDefault = 0
)

// the POJO
type Product struct {
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

// the Interface
type Products []Product

// the POJO Validation
func (product *Product) Validate() *errors.RestErr {
	product.Name = strings.TrimSpace(product.Name)
	product.Code = strings.TrimSpace(product.Code)

	if product.Name == "" {
		return errors.NewBadRequestError("product name must be filled")
	}

	if product.Code == "" {
		return errors.NewBadRequestError("product code must be filled")
	}

	if product.SellPrice.IsZero() {
		return errors.NewBadRequestError("Sell Price cannot be 0")
	}

	if product.BuyPrice.IsZero() {
		return errors.NewBadRequestError("Buy Price cannot be 0")
	}

	return nil
}
