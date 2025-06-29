package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maxabella/appgym/domain"
	"github.com/maxabella/appgym/services"
	"github.com/maxabella/appgym/utils"
)

type Inscription interface {
	GetUserActivities(*gin.Context)
	MakeInscription(*gin.Context)
	DeleteUserInscription(*gin.Context)
}

type InscriptionController struct {
	InscriptionService *services.InscriptionService
}

func (s *InscriptionController) GetUserActivities(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	valid, err := utils.ValidateToken(authHeader)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	useridstr, _ := c.Params.Get("userid")

	userid, err := strconv.Atoi(useridstr)
	if err != nil {
		panic(err.Error())
	}
	response, err := s.InscriptionService.GetUserActivities(uint(userid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func (s *InscriptionController) MakeInscription(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	valid, err := utils.ValidateToken(authHeader)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	// End token validation
	var inscriptionrequest domain.MakeInscription
	if err := c.ShouldBindJSON(&inscriptionrequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}
	err = s.InscriptionService.Makeinscription(inscriptionrequest)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadGateway, err.Error())
	}
}

func (s *InscriptionController) DeleteUserInscription(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	valid, err := utils.ValidateToken(authHeader)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	Iid, _ := c.Params.Get("id")
	intID, err := strconv.Atoi(Iid)
	if err != nil {
		panic(err.Error())
		return
	}
	err = s.InscriptionService.DeleteUserInscription(uint(intID))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(202)
}
