package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"usersApi/managers"
	"usersApi/repositories"
)

func GenerateToken(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	parsingError := c.BindJSON(&body)
	if parsingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var dbError, userFromDb = repositories.GetByEmail(body.Email)
	if dbError != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email or password"})
		return
	}

	var password = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(body.Password))
	if password != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email or password"})
		return
	}

	var err, token = managers.GenerateToken(userFromDb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
