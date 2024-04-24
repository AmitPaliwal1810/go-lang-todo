package helperfunctions

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		fmt.Println("getting error while executing the bcrypt function")
		return string(bytes), err
	}

	return string(bytes), nil
}
