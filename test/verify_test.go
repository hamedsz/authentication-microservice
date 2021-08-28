package test

import (
	userController "auth_micro/internal/controllers/user"
	"auth_micro/internal/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVerifyWithOkToken(t *testing.T)  {
	user := createUser(t)
	token := auth.GetAdabter().GenerateToken(user.Email)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	controller := userController.NewUserController()
	r.POST("/", controller.Verify)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
	c.Request.Header.Set("Authorization" , "Bearer "+ token)
	r.ServeHTTP(w, c.Request)

	assert.Equal(t , w.Code , http.StatusOK)
}

func TestVerifyWithNoToken(t *testing.T)  {
	createUser(t)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	controller := userController.NewUserController()
	r.POST("/", controller.Verify)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
	r.ServeHTTP(w, c.Request)

	assert.Equal(t , w.Code , http.StatusUnauthorized)
}

func TestVerifyWithWrongTokenFormatToken(t *testing.T)  {
	createUser(t)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	controller := userController.NewUserController()
	r.POST("/", controller.Verify)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
	c.Request.Header.Set("Authorization" , "hello world")
	r.ServeHTTP(w, c.Request)

	assert.Equal(t , w.Code , http.StatusUnauthorized)
}

func TestVerifyWhenHeaderFormatIsOkButTokenIsWrongToken(t *testing.T)  {
	createUser(t)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	controller := userController.NewUserController()
	r.POST("/", controller.Verify)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
	c.Request.Header.Set("Authorization" , "Bearer gaefavafeafawfea")
	r.ServeHTTP(w, c.Request)

	assert.Equal(t , w.Code , http.StatusUnauthorized)
}