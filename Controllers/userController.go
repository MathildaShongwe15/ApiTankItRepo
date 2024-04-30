package controllers

import (
	"log"
	"math/rand"
	initializers "myapp/Initializers"
	models "myapp/Models"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var body struct {
		Id                string
		ServiceProviderId *string
		First_name        string
		Last_name         string
		Email             string
		PhoneNumber       string
		Password          string
		Role              string
	}

	c.BindJSON(&body)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to hash password",
		})
	}

	user := models.User{Id: body.Id, ServiceProviderId: body.ServiceProviderId, First_Name: body.First_name, Last_Name: body.Last_name, Email: body.Email, PhoneNumber: body.PhoneNumber, Password: string(hash), Role: body.Role}
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

	var body struct {
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password1",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password2",
			"err":   err,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token":      tokenString,
		"role":       user.Role,
		"Id":         user.Id,
		"ProviderId": user.ServiceProviderId,
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

func ResetEmail() {

	otp := rand.Intn(999999)
	otpString := strconv.Itoa(otp)

	auth := smtp.PlainAuth(
		"",
		"tankitroadsideassistance@gmail.com",
		"mflqvpvhtjfvbevg",
		"smtp.gmail.com",
	)
	rand.Int()
	msg := "Subject: Reset Password\nYour OTP for reset" + otpString + "If this was not you please resport to tankitroadsideassistance@gmail.com"

	smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"tankitroadsideassistance@gmail.com",
		[]string{"tankitroadsideassistance@gmail.com"},
		[]byte(msg),
	)

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
func GetUserByEmail(c *gin.Context) {

	var user models.User
	id := c.Param(("email"))

	result := initializers.DB.Where("email = ?", id).First(&user)

	if result.Error != nil {
		log.Fatalf("cannot retrieve user: %v\n", result.Error)
	}

	initializers.DB.Find(&user)

	if http.StatusOK == 200 {
		c.JSON(http.StatusOK, gin.H{
			"message": "user does exist",
		})

	}
	if http.StatusOK == 404 {
		c.JSON(http.StatusOK, gin.H{
			"message": "user does NOT exist",
		})

	}

	//ResetEmail()
	//generate random OTP to user
	otp := rand.Intn(999999)
	otpString := strconv.Itoa(otp)

	auth := smtp.PlainAuth(
		"",
		"tankitroadsideassistance@gmail.com",
		"mflqvpvhtjfvbevg",
		"smtp.gmail.com",
	)
	rand.Int()
	msg := "Subject: Reset Password\nYour OTP for reset " + otpString + " If this was not you please report to tankitroadsideassistance@gmail.com"

	smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"tankitroadsideassistance@gmail.com",
		[]string{"tankitroadsideassistance@gmail.com"},
		[]byte(msg),
	)

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
