package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxabella/appgym/clients"
	"github.com/maxabella/appgym/controller"
	"github.com/maxabella/appgym/services"
	"github.com/maxabella/appgym/utils"
)

// DEBUG FUNCTION
func printRequestBody(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Print the raw body as bytes
	fmt.Printf("Raw Request Body (Bytes): %v\n", bodyBytes)

	// Print the raw body as a string
	fmt.Printf("Raw Request Body (String): %s\n", string(bodyBytes))

	// Important: Reset the body so that subsequent handlers can read it again
	c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	c.JSON(http.StatusOK, gin.H{"message": "Request body printed"})
}

// END DEBUG

func Start() {
	// Inicializamos los clientes,servicios y controladores
	MySQLClient := clients.Mysql_Client{}
	MySQLClient.Connect()
	MySQLClient.Migrate()
	Userclient := clients.UserClient{}
	Userclient.DbClient = &MySQLClient
	Userservice := services.UserService{}
	Userservice.Userclient = &Userclient
	Usercontroller := controller.UserController{}
	Usercontroller.Userservice = &Userservice
	Activitesclient := clients.ActivityClient{}
	Activitesclient.DbClient = &MySQLClient
	Activitiesservice := services.ActivityService{}
	Activitiesservice.ActivityClient = &Activitesclient
	ActivitiesController := controller.ActivityController{}
	ActivitiesController.ActivitySerivice = &Activitiesservice
	InscriptionClient := clients.InscriptionClient{}
	InscriptionClient.DbClient = &MySQLClient
	InscriptionService := services.InscriptionService{}
	InscriptionService.InscriptionClient = &InscriptionClient
	InscriptionController := controller.InscriptionController{}
	InscriptionController.InscriptionService = &InscriptionService
	// Iniciamos la app
	app := gin.Default()
	app.POST("/register" /*Usercontroller.CreateUser*/, Usercontroller.CreateUser)
	app.POST("login", utils.CORS, Usercontroller.Login)
	app.GET("/actividades", utils.CORS, ActivitiesController.GetActivities)
	app.GET("/user/:userid/activities", utils.CORS, InscriptionController.GetUserActivities)
	app.GET("/actividades/:id", utils.CORS, ActivitiesController.GetActivityById)
	app.POST("/users/inscription", utils.CORS, InscriptionController.MakeInscription) // New inscription for userid
	/*
		{
			Id_usuario
			Id_actividad
			Dia
			Starting_hour
			Finishing_hour
		}
	*/

	app.Run("localhost:8523")
}
