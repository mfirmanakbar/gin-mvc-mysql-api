package app

import (
	"github.com/mfirmanakbar/gin-mvc-mysql-api/controllers/products"
)

func Routes() {
	router.POST("/products", products.Create)
}
