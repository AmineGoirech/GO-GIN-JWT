package controllers

import (
	"fmt"
	"log"

	"github.com/AmineGoirech/gin-auth/database"
	"github.com/AmineGoirech/gin-auth/helper"
	"github.com/AmineGoirech/gin-auth/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type FormData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Login handler
func Login(c *gin.Context) {
	returnObject := gin.H{
		"status":  "ok",
		"message": "login router",
	}

	var FormData FormData
	if err := c.ShouldBind(&FormData); err != nil {
		log.Println("Error in Json Binding")
	}

	var User model.User
	database.DBConn.First(&User, "email=?", FormData.Email)

	if User.ID == 0 {
		returnObject["message"] = "USER NOT FOUND "
		c.JSON(400, returnObject)
		log.Println("USER NOT FOUND")
		return
	}

	//validate password
	err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(FormData.Password))
	if err != nil {
		log.Println("INVALID PASSWORD")
		returnObject["message"] = "INVALID PASSWORD"
		c.JSON(401, returnObject)
		return
	}

	//CREATE TOKEN :
	token, err := helper.GenerateToken(User)
	if err != nil {
		returnObject["message"] = "ERROR IN TOKEN CREATION"
		returnObject["status"] = "NO"
		c.JSON(401, returnObject)

	}

	returnObject["Token"] = token

	returnObject["message"] = "User authenticated"
	returnObject["status"] = "ok"
	log.Println("USER AUTHENTICATED")
	c.JSON(200, returnObject)
}

//Logout handler
func Logout(c *gin.Context) {
	returnObject := gin.H{
		"status":  "ok",
		"message": "logout route",
	}

	c.JSON(200, returnObject)
}

// register handler
func Register(c *gin.Context) {
	returnObject := gin.H{
		"status":  "ok",
		"message": "register route",
	}

	//collect from data
	var FormData FormData
	if err := c.ShouldBind(&FormData); err != nil {
		log.Println("Error in Json Binding")
	}

	// Add form data to model
	var user model.User
	user.Email = FormData.Email
	user.Password = helper.HashPassword(FormData.Password)

	result := database.DBConn.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		returnObject["message"] = "Used already exists bro"
		c.JSON(400, returnObject)
	}

	returnObject["message"] = "Used added bro"
	fmt.Println("User added succesf")
	c.JSON(200, returnObject)
}

func RefreshToken(c *gin.Context) {
	returnObject := gin.H{
		"status":  "ok",
		"message": "refresh token route",
	}
	c.JSON(200, returnObject)

}
