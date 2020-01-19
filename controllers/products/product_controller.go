package products

import (
	"github.com/gin-gonic/gin"
	"github.com/mfirmanakbar/gin-mvc-mysql-api/Services"
	"github.com/mfirmanakbar/gin-mvc-mysql-api/domain/products"
	"github.com/mfirmanakbar/gin-mvc-mysql-api/utils/errors"
	"net/http"
)

func Create(c *gin.Context) {
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := Services.ProductService.CreateProduct(product)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	// let say if Header go-access-token empty so it should be true
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("go-access-token") == ""))
}
