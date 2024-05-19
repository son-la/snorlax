package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/son-la/snorlax/internal/models"
	"github.com/son-la/snorlax/internal/utils"
	"github.com/google/uuid"

	"log"
)

// Function for logging in
func Login(c *gin.Context) {
	var loginInput map[string]interface{}

	// Check user credentials and generate a JWT token
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	username, ok := loginInput["username"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username field is missing or not a string"})
		return
	}

	// TODO: Lookup for username first

	password, ok := loginInput["password"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password field is missing or not a string"})
		return
	}

	// TODO: Password matching with user database
	if username == "user" && password == "password" {
		// Generate a JWT token
		token, err := utils.GenerateToken(uuid.New())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}


func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		log.Println(err)
		return
	}

	// TODO: Password hashing
	// TODO: Store user in database

	user.ID = uuid.New()
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}