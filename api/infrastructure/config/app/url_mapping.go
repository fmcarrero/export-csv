package app

import (
	"github.com/fmcarrero/export-csv/api/infrastructure/config/container/infrastructure_container"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	router.GET("/v1/movies", infrastructure_container.GetMovieController().ExportMovies)
}
