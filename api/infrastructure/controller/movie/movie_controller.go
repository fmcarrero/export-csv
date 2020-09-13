package movie

import (
	"github.com/fmcarrero/export-csv/api/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerMovie struct {
}

func (controller *ControllerMovie) ExportMovies(ctx *gin.Context) {

	doc, err := service.ExportMoviesCsv()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	if doc != nil {
		ctx.Header("Content-Disposition", "attachment; filename=export.csv")
		ctx.Header("Transfer-Encoding", "chunked")
		ctx.Data(http.StatusOK, "text/csv", doc.Bytes())
		return
	}
	ctx.Status(http.StatusInternalServerError)

}
