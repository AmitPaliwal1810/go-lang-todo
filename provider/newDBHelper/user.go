package newDBHelper

import (
	"fmt"
	helperfunctions "github.com/AmitPaliwal1810/go-lang-todo/provider/helperFunctions"

	"github.com/AmitPaliwal1810/go-lang-todo/models"
)

//================================= Create User / sign up ====================================

func (dh DbHelper) CreateUser(user models.CreateUser) error {

	insertQuery := `INSERT INTO users (email, password , name) VALUES ($1,$2,$3)`
	hashPassword, error := helperfunctions.HashPassword(user.Password)

	if error != nil {
		return error
	}

	_, err := dh.DB.Exec(insertQuery, user.Email, hashPassword, user.Name)

	if err != nil {
		fmt.Println("go error while accessing database for users")
		return err
	}

	return nil
}

// ===================================== Get all user =====================================

func (dh DbHelper) GetAllUsers() ([]models.User, error) {

	getUserQuery := `SELECT id, name, email, created_at FROM users WHERE archieved_at IS NULL`

	var allUsers []models.User

	err := dh.DB.Select(&allUsers, getUserQuery)

	if err != nil {
		return nil, err
	}

	fmt.Println("todo-data", allUsers)
	if len(allUsers) == 0 {
		fmt.Println("there no data")
	}

	return allUsers, nil
}

//========================================== Delete User ======================================

func (dh DbHelper) DeleteUser(user models.DeleteUser) error {

	archievedUser := `UPDATE users SET archieved_at=now() WHERE id=$1`

	_, err := dh.DB.Exec(archievedUser, user.Id)

	if err != nil {
		fmt.Println("getting error while executing archieved query")
		return err
	}

	return err
}

//=========================================== Update User ======================================

func (dh DbHelper) UpdateUser(user models.UpdateUser) error {
	updateUser := `UPDATE users SET email=$1, name=$2, password=$3 WHERE id=$4`

	hashPassword, error := helperfunctions.HashPassword(user.Password)

	if error != nil {
		return error
	}

	_, err := dh.DB.Exec(updateUser, user.Email, user.Name, hashPassword, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (dh DbHelper) LoginUser(user models.Login) (models.UpdateUser, error) {

	loginQuery := `SELECT id, password, email, name FROM users where email=$1`

	var userDetails models.UpdateUser

	err := dh.DB.Get(&userDetails, loginQuery, user.Email)

	if err != nil {
		return userDetails, err
	}

	return userDetails, nil

}
