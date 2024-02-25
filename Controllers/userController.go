package controllers

import (
	initializers "myapp/Initializers"
	models "myapp/Models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// import "net/http"
func SignUp(c *gin.Context) {

	//Get email/pass off req body
	var body struct {
		Id          string
		First_name  string
		Last_name   string
		PhoneNumber string
		Email       string
		Password    string
		Role        string
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	//create the user
	c.Bind(&body)
	user := models.User{Id: body.Id, First_Name: body.First_name, Last_Name: body.Last_name, PhoneNumber: body.PhoneNumber, Email: body.Email, Password: string(hash), Role: body.Role}
	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	//get the email  and pass off request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//compare sent in pass with saved user pass hash

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//generate a jwt toke

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	//send it back
	//cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})

}
