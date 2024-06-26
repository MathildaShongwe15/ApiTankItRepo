package middleware

import (
	"fmt"
	initializers "myapp/Initializers"
	models "myapp/Models"
	"net/http"

	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	//get the cookie off req

	tokenString, err := c.Cookie(("Authorization"))

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode and validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(("Unexpected signing method: %v"), token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User

		initializers.DB.First(&user, claims["sub"])

		if user.Id == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	} else {

		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//check the exp

	//find the user with token sub
}
