package controllers

import (
	"encoding/json"
	"github.com/gichohi/go-rest.git/models"
	"github.com/gichohi/go-rest.git/services"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	json.NewDecoder(req.Body).Decode(&user)
	var response = services.CreateUser(user)
	json.NewEncoder(w).Encode(response)
}

func TestAPI(w http.ResponseWriter, r *http.Request) {
	var resp models.Response
	resp.Code = 200
	resp.Message = "This is it"
	json.NewEncoder(w).Encode(resp)
}

func Index(w http.ResponseWriter, r *http.Request) {

	var resp models.Response
	resp.Code = 200
	resp.Message = "Give it to me"
	json.NewEncoder(w).Encode(resp)
}

func  LoginUser(w http.ResponseWriter, req *http.Request) {
	var loginreq models.LoginRequest
	json.NewDecoder(req.Body).Decode(&loginreq)
	username := loginreq.UserName
	password := loginreq.Password
	var user = services.Login(username,password)
	json.NewEncoder(w).Encode(user)
}

func  GetUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var user = services.GetUser(params["username"])
	json.NewEncoder(w).Encode(user)
}





