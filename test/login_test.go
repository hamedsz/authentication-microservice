package test

import (
	config "auth_micro/config/database"
	"auth_micro/internal/controllers/user"
	userModel "auth_micro/internal/models/user"
	"auth_micro/internal/services/auth"
	"auth_micro/internal/services/database"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func TestLoginWithOkBody(t *testing.T){
	createUser(t)

	userController := user.NewUserController()

	w := testController(
		userController.Login ,
		map[string]interface{}{
			"email":    "hamed@mail.com",
			"password": "as12AS!@",
		},
	)

	assert.Equal(t, w.Code , http.StatusOK)

	_ = database.GetAdabter(config.GetDefault()).DropAll()
}

func TestLoginGeneratedTokenIsValid(t *testing.T) {
	createUser(t)

	userController := user.NewUserController()

	w := testController(
		userController.Login ,
		map[string]interface{}{
			"email":    "hamed@mail.com",
			"password": "as12AS!@",
		},
	)

	var result user.LoginResponse

	err := json.Unmarshal([]byte(w.Body.String()), &result)

	if err != nil{
		t.Errorf("sing up response is not valid")
	}

	err = auth.GetAdabter().ValidateToken(result.Token)

	if err != nil {
		t.Errorf("sing up generated token is not valid")
	}
}

func createUser(t *testing.T) userModel.Model {

	hashedPassword , err := bcrypt.GenerateFromPassword([]byte("as12AS!@") , 0)

	if err != nil{
		t.Errorf("error in generate password")
	}

	model := userModel.Model{
		Email:     "hamed@mail.com",
		Password: string(hashedPassword),
		FirstName: "hamed",
		LastName:  "sz",
		Role:      "user",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	err = model.Save()

	if err != nil{
		t.Errorf("error in saving user")
	}

	return model
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