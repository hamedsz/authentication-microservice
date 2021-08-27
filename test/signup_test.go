package test

import (
	config "auth_micro/config/database"
	"auth_micro/internal/controllers/user"
	userModel "auth_micro/internal/models/user"
	"auth_micro/internal/services/auth"
	"auth_micro/internal/services/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignupWithOkBody(t *testing.T)  {
	w := signUpRequestWithOkBody()
	assert.Equal(t , http.StatusOK , w.Code)

	_ = database.GetAdabter(config.GetDefault()).DropAll()
}

func TestSignupWithOkBodyIsInserted(t *testing.T)  {
	signUpRequestWithOkBody()

	count , err := database.
		GetAdabter(config.GetDefault()).
		Count(
			"users" ,
			gin.H{},
		)

	if err != nil {
		t.Errorf("get count of user err %f" , err)
		return
	}

	assert.Equal(t , int64(1) , count)

	_ = database.GetAdabter(config.GetDefault()).DropAll()
}

func TestSignupIsHashedPasswordOk(t *testing.T)  {
	signUpRequestWithOkBody()

	result , _ := userModel.Find(gin.H{
		"email": "admin@mail.com",
	})

	err := bcrypt.CompareHashAndPassword([]byte(result.Password) , []byte("as12AS!@"))

	if err != nil {
		t.Error( "hashed password is not compare to selected password")
		return
	}

	_ = database.GetAdabter(config.GetDefault()).DropAll()
}

func TestSignupAllFieldsRequired(t *testing.T)  {
	data := map[string]interface{}{
		"first_name":"hamed",
		"last_name":"sz",
		"email":    "admin@mail.com",
		"password": "as12AS!@",
	}

	userController := user.NewUserController()

	for key , _ := range data {
		reqData := data
		reqData[key] = nil

		w := testController(
			userController.Signup ,
			reqData,
		)

		assert.Equal(t , w.Code , http.StatusBadRequest)
	}
}

func TestSignupWithWrongEmailFormat(t *testing.T)  {
	userController := user.NewUserController()
	w :=  testController(
		userController.Signup ,
		map[string]interface{}{
			"first_name":"hamed",
			"last_name":"sz",
			"email":    "hello",
			"password": "as12AS!@",
		},
	)

	assert.Equal(t , w.Code , http.StatusBadRequest)

	_ = database.GetAdabter(config.GetDefault()).DropAll()
}

func TestSignupResponseIsValid(t *testing.T) {
	w := signUpRequestWithOkBody()

	var result user.SignupResponse

	err := json.Unmarshal([]byte(w.Body.String()), &result)

	if err != nil{
		t.Errorf("sing up response is not valid")
	}

	_ = database.GetAdabter(config.GetDefault()).DropAll()
}

func TestSignupGeneratedTokenIsValid(t *testing.T) {
	w := signUpRequestWithOkBody()

	var result user.SignupResponse

	err := json.Unmarshal([]byte(w.Body.String()), &result)

	if err != nil{
		t.Errorf("sing up response is not valid")
	}

	err = auth.GetAdabter().ValidateToken(result.Token)

	if err != nil {
		t.Errorf("sing up generated token is not valid")
	}
}

func signUpRequestWithOkBody() *httptest.ResponseRecorder{
	userController := user.NewUserController()
	return testController(
		userController.Signup ,
		map[string]interface{}{
			"first_name":"hamed",
			"last_name":"sz",
			"email":    "admin@mail.com",
			"password": "as12AS!@",
		},
	)
}