package helper

import (
	"log"
	"time"

	"github.com/AmineGoirech/gin-auth/model"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Email  string
	UserID uint

	jwt.RegisteredClaims
}

var secret string = "GATHEK"

func GenerateToken(user model.User) (string, error) {
	claims := CustomClaims{
		user.Email,
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * 5)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Println("Error in Token signing")
		return "", err
	}

	log.Println("TOKEN SET WITH SUCCESS")

	return t, nil
}

func ValidateToken(clienttoken string) (claims *CustomClaims, msg string) {
	token, err := jwt.ParseWithClaims(clienttoken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}

	claims,ok := token.Claims.(*CustomClaims)
	if !ok {
		msg = err.Error()
		return
	}

	return claims,msg
}
