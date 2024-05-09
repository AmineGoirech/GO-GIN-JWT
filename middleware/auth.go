package middleware

import (
	"log"

	"github.com/AmineGoirech/gin-auth/helper"
	"github.com/gin-gonic/gin"
)


func Authenticate(c *gin.Context){
	token := c.GetHeader("token")
	if token==""{
		c.JSON(400,gin.H{"error":"Token not present"})
		c.Abort()
		return
	}
	claims,msg := helper.ValidateToken(token)
	log.Println(claims)
	if msg!=""{
		c.JSON(401,gin.H{"error":msg})
		c.Abort()
		return
	}
	c.Next()

}