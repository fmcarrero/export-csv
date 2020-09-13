package app

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"os"
)

func StartApp() {

	router := gin.Default()
	MapUrls(router)
	port := os.Getenv("PORT")

	if port == "" {
		port = ":" + "8080"
	}

	if err := router.Run(port); err != nil {
		logger.Infof("err running server on port %s", port)
	}
}
