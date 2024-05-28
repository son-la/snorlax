package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/son-la/snorlax/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Save(user *models.User) error {
	return nil
}

func (m *mockUserRepository) FindByEmail(email string) *models.User {
	return &models.User{}
}

func TestRegisterUser(t *testing.T) {

	mockUserRepo := new(mockUserRepository)
	mockUserRepo.On("Save", mock.Anything).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a new instance of the BaseHandler
	handler := NewBaseHandler(mockUserRepo)
	router.POST("/register", handler.RegisterUser)

	t.Run("Normal user input", func(t *testing.T) {

	})

	t.Run("Normal user input", func(t *testing.T) {
		reqBody := []byte(`{"username": "example_user", "password": "example_password", "name": "user", "email": "test_user@gmail.com"}`)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		// Create a new HTTP response recorder
		recorder := httptest.NewRecorder()

		// Serve the HTTP request
		router.ServeHTTP(recorder, req)

		// Assert the response status code
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})

	t.Run("Missing user field", func(t *testing.T) {
		reqBody := []byte(`{"username": "example_user", "password": "example_password", "name": "user"}`)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		// Create a new HTTP response recorder
		recorder := httptest.NewRecorder()

		// Serve the HTTP request
		router.ServeHTTP(recorder, req)

		// Assert the response status code
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

}
