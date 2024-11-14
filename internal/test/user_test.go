package test

import (
	"TODO-List/internal/model/request"
	"TODO-List/internal/repository/user"
	"TODO-List/internal/service"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	connStr := "host=localhost port=5432 user=postgres dbname=todo_list_test sslmode=disable password=1234"

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(db)
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

	userRepository := user.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	context, w := CreateTestContext("POST", "/register", request.CreateUserRequest{
		Username:  "john",
		Password:  "123",
		FirstName: "John",
		LastName:  "Doe",
	})

	userService.CreateUser(context)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	token, ok := response["token"]
	assert.True(t, ok)
	assert.NotEmpty(t, token)
}

func CreateTestContext(method, url string, body interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	var jsonBody []byte
	if body != nil {
		jsonBody, _ = json.Marshal(body)
	}
	req := httptest.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Создаем recorder для получения ответа
	w := httptest.NewRecorder()

	// Создаем gin context с recorder и request
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	return ctx, w
}
