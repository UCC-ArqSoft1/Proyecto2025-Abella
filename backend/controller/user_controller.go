package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxabella/appgym/domain"
	"github.com/maxabella/appgym/services"
	"github.com/maxabella/appgym/utils"
)

type User interface {
	CreateUser(*gin.Context)
	Login(*gin.Context)
	GetCoaches(*gin.Context)
}

type UserController struct {
	Userservice *services.UserService
}

func (s *UserController) Login(c *gin.Context) {
	var requestedUser domain.UserLoginRequest
	if err := c.BindJSON(&requestedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Invalid format error
		return
	}
	UserLogged, err := s.Userservice.Login(requestedUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserLogged)
}

func (s *UserController) CreateUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*") // Be specific about the origin
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type") // Only the headers you expect
	c.Header("Access-Control-Allow-Credentials", "true")
	var UserData domain.UserRegister

	if err := c.BindJSON(&UserData); err != nil {
		fmt.Println("Now allowed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserResponse, err := s.Userservice.CreateUser(UserData)
	if err != nil {
		c.Status(404)
	} else {
		c.JSON(http.StatusOK, UserResponse)
	}
}

func (s *UserController) GetCoaches(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	valid, err := utils.ValidateToken(authHeader)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	Coaches, err := s.Userservice.GetCoaches()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, Coaches)
}
