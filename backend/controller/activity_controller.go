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

type Activity interface {
	GetActivities(*gin.Context) (domain.Activities, error)
	GetActivityById(*gin.Context)
}

type ActivityController struct {
	ActivitySerivice *services.ActivityService
}

func (c *ActivityController) GetActivities(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	valid, err := utils.ValidateToken(authHeader)
	if err != nil {
		fmt.Println(err.Error())
	}
	if valid {
		ctx.Writer.Header().Set("IsTokenValid", "true")
	}
	if !valid {
		ctx.Writer.Header().Set("IsTokenValid", "false")
	}
	if value, ok := ctx.GetQuery("keyword"); !ok {
		Activities, err := c.ActivitySerivice.GetActivities()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, Activities)
	} else {
		Activities, err := c.ActivitySerivice.GetActivityByKeyword(value)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, Activities)
	}
}

func (s *ActivityController) GetActivityById(c *gin.Context) {
	id, found := c.Params.Get("id")
	fmt.Println("string id: ", id)
	if !found {
		c.JSON(http.StatusBadRequest, fmt.Errorf("could not get id from parameters"))
		return
	}
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter type ", err.Error()))
	}
	fmt.Println("int id: ", idint)

	Activity, err := s.ActivitySerivice.GetActivityById(idint)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("Could not process request", err.Error()))
	}
	c.JSON(202, Activity)
}
