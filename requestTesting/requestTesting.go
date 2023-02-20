package requestTesting

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func RequestTest(t *testing.T, router *gin.Engine, method string, uri string, body io.Reader, expectedCode int) string {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, body)
	router.ServeHTTP(w, req)

	assert.Equal(t, expectedCode, w.Code)
	return w.Body.String()
}

func GetTest(t *testing.T, router *gin.Engine, uri string, expectedCode int) string {
	return RequestTest(t, router, "GET", uri, nil, expectedCode)
}

func PostTest(t *testing.T, router *gin.Engine, uri string, body io.Reader, expectedCode int) string {
	return RequestTest(t, router, "POST", uri, body, expectedCode)
}
