package middleware

import (
	"auth-demo/result"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Auth(c *gin.Context)  {
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("123456"), nil
	})
	if token!=nil&&token.Valid {
		fmt.Println("You look nice today")
		c.Next()
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			c.JSON(http.StatusUnauthorized,result.Error(gin.H{},"401 Unauthorized"))
			c.Abort()
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			c.JSON(http.StatusUnauthorized,result.Error(gin.H{},"401 Unauthorized"))
			c.Abort()
		} else {
			fmt.Println("Couldn't handle this token:", err)
			c.JSON(http.StatusUnauthorized,result.Error(gin.H{},"401 Unauthorized"))
			c.Abort()
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		c.JSON(http.StatusUnauthorized,result.Error(gin.H{},"401 Unauthorized"))
		c.Abort()
	}
}
