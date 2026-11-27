package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	auth := gin.BasicAuthForRealm(gin.Accounts{"demo": "belajar"}, "latihan")
	router.GET("/profil", auth, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": c.GetString(gin.AuthUserKey)})
	})
	log.Fatal(router.Run("127.0.0.1:8082"))
}
