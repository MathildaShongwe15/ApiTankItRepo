package controllers

import (
	"log"
	initializers "myapp/Initializers"
	models "myapp/Models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// import "net/http"
func SignUp(c *gin.Context) {

	//Get email/pass off req body
	var body struct {
		Id          string
		First_name  string
		Last_name   string
		Email       string
		PhoneNumber string
		Password    string
		Role        string
	}

	c.BindJSON(&body)
	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to hash password",
		})
	}

	//create the user
	user := models.User{Id: body.Id, First_Name: body.First_name, Last_Name: body.Last_name, Email: body.Email, PhoneNumber: body.PhoneNumber, Password: string(hash), Role: body.Role}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func Login(c *gin.Context) {
	//get the email  and pass off request body
	var body struct {
		Email    string
		Password string
		Role     string
	}

	if c.BindJSON(&body) != nil {
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
			"error": "Invalid email or password1",
		})
		return
	}
	//compare sent in pass with saved user pass hash

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password2",
			"err":   err,
		})
		return
	}
	//generate a jwt token

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

	if user.Role != body.Role {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Role selected",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"role":  user.Role,
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
func ResetPassword(c *gin.Context) {
	var user models.User
	email := c.Param(("email"))

	var body struct {
		Password string
	}

	c.Bind(&body)

	result := initializers.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		log.Fatalf("cannot retrieve user: %v\n", result.Error)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to hash password",
		})
	}

	initializers.DB.Model(&user).Updates(models.User{
		Password: string(hash),
	})

	c.JSON(200, gin.H{
		"result": " Password Updated successsfully!",
	})
}

func GetUserById(c *gin.Context) {

	var user models.User
	id := c.Param(("id"))

	result := initializers.DB.Where("Id = ?", id).First(&user)

	if result.Error != nil {
		log.Fatalf("cannot retrieve provider: %v\n", result.Error)
	}

	initializers.DB.Find(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	var user models.User
	id := c.Param(("id"))

	var body struct {
		First_name  string
		Last_name   string
		Email       string
		PhoneNumber string
	}

	c.Bind(&body)

	result := initializers.DB.Where("Id = ?", id).First(&user)

	if result.Error != nil {
		log.Fatalf("cannot retrieve request: %v\n", result.Error)
	}

	initializers.DB.Model(&user).Updates(models.User{
		First_Name:  body.First_name,
		Last_Name:   body.Last_name,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
	})

	c.JSON(200, gin.H{
		"result": "user Updated successsfully!",
	})
}
