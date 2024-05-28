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

	tests := map[string]struct {
		input  string
		result int
	}{
		"valid input 1": {
			input:  `{"username": "example_user", "password": "example_password", "name": "user", "email": "test_user@gmail.com"}`,
			result: http.StatusCreated,
		},
		"valid input 2": {
			input:  `{"username": "example_user", "password": "verysecurepassword", "name": "user2", "email": "test_user232@gmail.com"}`,
			result: http.StatusCreated,
		},
		"missing field 1": {
			input:  `{"username": "example_user", "password": "example_password", "name": "user"}`,
			result: http.StatusBadRequest,
		},
		"missing field 2": {
			input:  `{"username": "example_user", "password": "example_password", "email": "test_user@gmail.com"}`,
			result: http.StatusBadRequest,
		},
	}

	mockUserRepo := new(mockUserRepository)
	mockUserRepo.On("Save", mock.Anything).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a new instance of the BaseHandler
	handler := NewBaseHandler(mockUserRepo)
	router.POST("/register", handler.RegisterUser)

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			reqBody := []byte(test.input)

			req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)

			// Create a new HTTP response recorder
			recorder := httptest.NewRecorder()

			// Serve the HTTP request
			router.ServeHTTP(recorder, req)

			// Assert the response status code
			assert.Equal(t, test.result, recorder.Code)
		})
	}
}
