package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rajesh4b8/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("Application is about to start...")
	router.Run("127.0.0.1:8080")
}
