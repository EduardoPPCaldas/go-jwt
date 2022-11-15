package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/EduardoPPCaldas/go-jwt/initializers"
	"github.com/EduardoPPCaldas/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		fmt.Println("Error parsing jwt")
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("Token expired")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			fmt.Println("User not found")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Set("user", user)
		ctx.Next()
	} else {
		fmt.Println("Error on getting token claims")
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
