package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxabella/appgym/domain"
	"github.com/maxabella/appgym/services"
	"github.com/maxabella/appgym/utils"
)

type Activity interface {
	GetActivities(*gin.Context) (domain.Activities, error)
	GetActivityById(int)
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
