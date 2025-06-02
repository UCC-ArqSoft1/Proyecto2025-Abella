package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpCORS(c *gin.Context) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("Content-Type") // Use the "Add" method
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("GET")
	corsConfig.AddAllowMethods("POST")
	corsConfig.AddAllowMethods("PUT")
	corsConfig.AddAllowMethods("DELETE")

}

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
}
