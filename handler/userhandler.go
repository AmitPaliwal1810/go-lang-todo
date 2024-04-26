// we are the package name such that where folder they belongs
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AmitPaliwal1810/go-lang-todo/models"
)

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
