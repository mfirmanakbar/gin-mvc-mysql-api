package Services

import (
	"github.com/mfirmanakbar/gin-mvc-mysql-api/domain/products"
	"github.com/mfirmanakbar/gin-mvc-mysql-api/utils/date_utils"
	"github.com/mfirmanakbar/gin-mvc-mysql-api/utils/errors"
)

var (
	ProductService productServiceInterface = &productsService{}
)

type productsService struct{}

type productServiceInterface interface {
	CreateProduct(products.Product) (*products.Product, *errors.RestErr)
}

func (s *productsService) CreateProduct(product products.Product) (*products.Product, *errors.RestErr) {
	if err := product.Validate(); err != nil {
		return nil, err
	}

	product.CreatedAt = date_utils.CurrentDateDbFormat()
	product.ProductBundle = products.ProductBundleDefault
	if err := product.Save(); err != nil {
		return nil, err
	}

	return &product, nil
}
