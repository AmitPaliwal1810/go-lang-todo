package newDBHelper

import (
	"fmt"

	"github.com/AmitPaliwal1810/go-lang-todo/models"
	helperfunctions "github.com/AmitPaliwal1810/go-lang-todo/provider/helperFunctions"
)

func (dh DbHelper) CreateUser(user models.CreateUser) error {

	insertQuery := `INSERT INTO users (email, password , name) VALUES ($1,$2,$3)`

	hashPassword, error := helperfunctions.HashPassword(user.Password)

	fmt.Println("hashedPassword", hashPassword)

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
