package handler

import (
	"auth-demo/result"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
var   hmacSampleSecret  = []byte("123456")
type CustomClaims struct {
	Username string `json:"username"`
	UserId string `json:"userId"`
	jwt.StandardClaims
}
type LoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func getToken(claims jwt.Claims) string{
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(hmacSampleSecret)
	return tokenString
}
func Login(c *gin.Context) {
	var param LoginParam
	if bindErr := c.ShouldBindJSON(&param); bindErr != nil {
		c.JSON(http.StatusBadRequest, result.Error(gin.H{
			"error": bindErr.Error(),
		}, bindErr.Error()))
		return
	}
	if param.Username != "admin" && param.Password != "admin" {
		c.JSON(http.StatusOK, result.Error(gin.H{
			"user": "user-admin",
		}, "用户名密码错误"))
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"user": "user",
		"token":getToken(CustomClaims{
			param.Username,
			param.Username,
			jwt.StandardClaims{
				ExpiresAt:time.Now().Unix()+5,

			},
		}),
	}))
}
func Info(c *gin.Context) {
	c.JSON(http.StatusOK, result.Success(gin.H{
		"user": "user",
	}))
}
