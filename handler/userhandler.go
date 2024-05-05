// we are the package name such that where folder they belongs
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AmitPaliwal1810/go-lang-todo/models"
	helperfunctions "github.com/AmitPaliwal1810/go-lang-todo/provider/helperFunctions"
)

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (srv *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	var createUserData models.CreateUser

	err := json.NewDecoder(r.Body).Decode(&createUserData)

	if err != nil {
		fmt.Println("getting error while decoding the craeteUsers")
		return
	}

	err = srv.DbHelper.CreateUser(createUserData)

	if err != nil {
		return
	}

	response := Response{Message: "successfully Inserted"}

	w.Header().Set("Content-type", "application/json")

	_ = json.NewEncoder(w).Encode(response)

}

func (srv *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	data, err := srv.DbHelper.GetAllUsers()

	if err != nil {
		fmt.Println("getting error while executing getAlluser", err)
	}

	_ = json.NewEncoder(w).Encode(data)
}

func (srv *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userData models.DeleteUser
	err := json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		return
	}

	err = srv.DbHelper.DeleteUser(userData)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	response := Response{Message: "succeessfully delete"}

	_ = json.NewEncoder(w).Encode(response)
}

func (srv *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userData models.UpdateUser

	err := json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		return
	}

	err = srv.DbHelper.UpdateUser(userData)

	if err != nil {
		return
	}

	respose := Response{Message: "Successfully updated"}

	_ = json.NewEncoder(w).Encode(respose)
}

func (srv *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userData models.Login
	var data models.UpdateUser
	var isPasswordMatch bool
	var response LoginResponse
	var tokenSend helperfunctions.CreateToken

	err := json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		return
	}

	data, err = srv.DbHelper.LoginUser(userData)

	if err != nil {
		return
	}

	isPasswordMatch = helperfunctions.CheckPasswordHash(userData.Password, data.Password)

	tokenSend = helperfunctions.CreateToken{
		Id:    data.Id,
		Name:  data.Name,
		Email: data.Email,
	}

	token, error := helperfunctions.CreateJWTToken(tokenSend)

	expoToken, err1 := helperfunctions.ExpoJWTTokenData(token)
	if err1 != nil {
		fmt.Println("err", err1)
		return
	}

	fmt.Println("epxotoken", expoToken)

	if error != nil {
		fmt.Println("fat gya 1")
		response = LoginResponse{Message: "Not Authorized", Token: "null"}
	} else if !isPasswordMatch {
		fmt.Println("fat gya 1")
		response = LoginResponse{Message: "Not Authorized", Token: "null"}
	} else {
		response = LoginResponse{Message: "Authorized", Token: token}

	}

	_ = json.NewEncoder(w).Encode(response)

}
