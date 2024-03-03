package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"usersApi/models"
	"usersApi/repositories"
	"usersApi/services"
)

func AddUser(c *gin.Context) {
	var body struct {
		FirstName   string `json:"Firstname"`
		LastName    string `json:"Lastname"`
		Email       string `json:"Email"`
		Password    string `json:"Password"`
		PhoneNumber string `json:"PhoneNumber"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error while hashing password", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}
	user := models.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		Password:    string(password),
		PhoneNumber: body.PhoneNumber,
	}
	_, err = repositories.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with provided email already exists"})
		return
	}

	services.PublishUserCreatedEvent(user)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetUserById(c *gin.Context) {
	idStr := c.Param("userId")
	id, err := strconv.Atoi(idStr)
	user, err := repositories.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
