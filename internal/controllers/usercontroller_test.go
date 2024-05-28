package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/son-la/snorlax/internal/models"
	"github.com/stretchr/testify/assert"
)

type mockUserRepository struct {
	save        func(user *models.User) error
	findByEmail func(email string) *models.User
}

func (m *mockUserRepository) Save(user *models.User) error          { return m.save(user) }
func (m *mockUserRepository) FindByEmail(email string) *models.User { return m.findByEmail(email) }

func TestRegisterUser(t *testing.T) {

	mockUserRepo := &mockUserRepository{
		save: func(user *models.User) error {
			return nil
		},
		findByEmail: func(email string) *models.User {
			return &models.User{}
		},
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a new instance of the BaseHandler
	handler := NewBaseHandler(mockUserRepo)

	// Register the route
	router.POST("/register", handler.RegisterUser)

	// Create a new HTTP request
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
}
