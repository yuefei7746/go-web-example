package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	router := gin.New()
	router.Use(gin.BasicAuth(AuthorizationAccount))
	router.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, c.MustGet(gin.AuthUserKey).(string))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.Header.Set("Authorization", "Basic eWYuemhhb0B1cGx0di5jb206MTIzMTIz")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "yf.zhao@upltv.com", w.Body.String())
}

func TestAuthorizationHeader(t *testing.T) {
	assert.Equal(t, "Basic eWYuemhhb0B1cGx0di5jb206MTIzMTIz", authorizationHeader("yf.zhao@upltv.com", "123123"))
}
