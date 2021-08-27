package test

import (
	"auth_micro/internal/controllers/user"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginWithNoBody(t *testing.T)  {
	userController := user.NewUserController()

	w := testController(
		userController.Login ,
		map[string]interface{}{},
	)

	assert.Equal(t , http.StatusBadRequest , w.Code)
}

func TestLoginWithWrongCredentials(t *testing.T)  {
	userController := user.NewUserController()

	w := testController(
		userController.Login ,
		map[string]interface{}{
			"email":    "asgag",
			"password": "asgag",
		},
	)

	assert.Equal(t , http.StatusBadRequest , w.Code)
}

func testController(controller gin.HandlerFunc , data map[string]interface{}) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/", controller)

	jsonData, _ := json.Marshal(data)

	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonData))
	r.ServeHTTP(w, c.Request)

	return w
}