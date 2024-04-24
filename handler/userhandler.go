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
