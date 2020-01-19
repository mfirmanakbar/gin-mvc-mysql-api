package products

import (
	"github.com/mfirmanakbar/gin-mvc-mysql-api/datasources/mysql/pos_db"
	"github.com/mfirmanakbar/gin-mvc-mysql-api/logger"
	"github.com/mfirmanakbar/gin-mvc-mysql-api/utils/errors"
)

const (
	queryInsert = "INSERT INTO products(name, code, sell_price, buy_price, created_at, product_bundle) VALUES (?, ?, ?, ?, ?, ?);"
)

func (product *Product) Save() *errors.RestErr {
	stmt, err := pos_db.Client.Prepare(queryInsert)
	if err != nil {
		logger.Info("error when trying to prepare save product statement. " + err.Error())
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(product.Name, product.Code, product.SellPrice, product.BuyPrice, product.CreatedAt, product.ProductBundle)
	if saveErr != nil {
		logger.Error("error when trying to save product" + err.Error())
		return errors.NewInternalServerError("database error")
	}

	productId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get Last Insert Id product after creating a new product" + err.Error())
		return errors.NewInternalServerError("database error")
	}
	product.Id = productId
	return nil
}
