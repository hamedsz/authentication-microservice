package test

import (
	config "auth_micro/config/database"
	"auth_micro/internal/services/database"
	"os"
	"testing"
)

func init()  {

}


func TestMain(m *testing.M) {

	exitVal := m.Run()

	_ = database.GetAdabter(config.GetDefault()).DropAll()
	os.Exit(exitVal)
}