package services

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gichohi/go-rest.git/models"
	"github.com/gichohi/go-rest.git/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	DB_HOST 	= "127.0.0.1"
	DB_PORT 	= "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "kim0nd@"
	DB_NAME     = "datasim"
)

var db = utils.GetDB()
var err error
var jwtKey = []byte("@Qangari2013#")

func CreateUser(user models.User) models.Response {
	var resp = models.Response{}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		resp.Code = 500
		resp.Message = err.Error()
		return resp
	}
	user.Password = string(pass)
	db.Create(&user)


	resp.Code = 200
	resp.Message = "Success"
	return resp

}

func GetUser(username string) models.User {
	var user models.User
	db.Where("email = ?", username).Find(&user)
	return user
}

func DeleteUser(username string) models.Response {
	var user models.User
	db.Where("email = ?", username).Find(&user)
	db.Delete(&user)
	var resp = models.Response{}
	resp.Code = 200
	resp.Message = "Success"
	return resp
}

func Login(email, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var response = map[string]interface{}{"code": 500, "message": "Email address not found"}
		return response
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"code": 500, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	expiresAt := time.Now().Add(time.Minute * 14400).Unix()

	token := &models.Token{
		UserID: user.ID,
		Name:   user.FirstName + " " + user.LastName,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	useresponse := &models.LoginResponse{
		FirstName:   user.FirstName,
		LastName:   user.LastName,
		Email:  user.Email,
	}

	accesstoken := jwt.NewWithClaims(jwt.GetSigningMethod("HS512"), token)
	tokenString, err := accesstoken.SignedString(jwtKey)
	if err != nil {
		var response = map[string]interface{}{"code": 500, "message": "Email address not found"}
		return response
	}
	var resp = map[string]interface{}{"code": 200, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = useresponse
	return resp
}

