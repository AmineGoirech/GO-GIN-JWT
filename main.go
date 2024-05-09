package main

import (
	"log"
	"os"

	"github.com/AmineGoirech/gin-auth/database"
	"github.com/AmineGoirech/gin-auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Cant load the env file .")
	}
	database.ConnectDB()
}

func main() {

	//close the db connection 
	sqlDB,err := database.DBConn.DB()
	if err != nil {
		log.Println("error in getting db conn")
	}
	defer sqlDB.Close()


	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	
	// server.GET("/hello", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"data":  "welcome",
	// 		"amine": "gouerch",
	// 	})
	// })

	routes.SetupRoutes(server)
	port := os.Getenv("PORT")
	server.Run(":"+port)

}
