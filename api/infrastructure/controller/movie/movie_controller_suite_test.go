package movie_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fmcarrero/export-csv/api/infrastructure/config/app"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInsuranceExportExportMovies(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Export movies in csv format Suite")
}

var _ = Describe("Export movies in csv format", func() {

	Context("Export data correctly", func() {
		var router *gin.Engine
		BeforeEach(func() {
			_ = os.Setenv("PORT", ":4578")
			router = gin.Default()
			app.MapUrls(router)
		})
		AfterEach(func() {
			os.Clearenv()
		})
		When("The user is authenticate", func() {
			It("Should return a bytes with the movies data", func() {

				w := httptest.NewRecorder()
				req, err := http.NewRequest(http.MethodGet, "/v1/movies", nil)
				router.ServeHTTP(w, req)

				Expect(err).To(BeNil(), "err should be null when calling to movies endpoint")
				Expect(http.StatusOK).To(Equal(w.Code))
				Expect("text/csv").To(Equal(w.Header().Get("Content-Type")))
				Expect(w.Body.String()).To(Not(BeNil()))
			})
		})
	})
})
