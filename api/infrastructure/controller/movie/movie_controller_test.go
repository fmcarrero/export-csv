package movie_test

import (
	"github.com/fmcarrero/export-csv/api/infrastructure/config/app"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExport(t *testing.T) {

	router := gin.Default()
	app.MapUrls(router)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/movies", nil)
	router.ServeHTTP(w, req)

	assert.NoError(t, err, "err should be null when calling to movies endpoint")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "text/csv", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Header().Get("Content-Disposition"), ".csv")
	assert.NotNil(t, w.Body.String())
}
