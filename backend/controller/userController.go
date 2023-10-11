package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kiridharan/initializers"
	"github.com/kiridharan/models"
	"golang.org/x/crypto/bcrypt"
)

// The SignUp function is a handler function for the sign-up route in a Gin web framework.
func SignUp(c *gin.Context) {

	// get the request body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	// validate the request body

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}

	// hash the password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		return

	}
	// create user

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashPass),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	return
}
