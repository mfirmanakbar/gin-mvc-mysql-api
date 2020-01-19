package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mfirmanakbar/gin-mvc-mysql-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	port := ":8282"
	Routes()
	logger.Info(fmt.Sprintf("Started Application on Port %s", port))
	router.Run(port)
}
