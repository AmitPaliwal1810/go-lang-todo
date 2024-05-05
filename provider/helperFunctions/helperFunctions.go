package helperfunctions

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CreateToken struct {
	Id    string `json:"id"`
	Name  string `json:"userName"`
	Email string `json:"email"`
}

//========================= hash password func========================

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		fmt.Println("getting error while executing the bcrypt function")
		return string(bytes), err
	}

	return string(bytes), nil
}

//============================= check hash password ====================

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

//========================= generate JWT token =============================

func CreateJWTToken(userDetail CreateToken) (string, error) {

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}

	fmt.Println("key", key)

	// short code

	// token, err := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
	// 	"username": username,
	// 	"exp":      time.Now().Add(time.Hour * 24).Unix(),
	// }).SignedString(secretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": userDetail.Name,
		"id":       userDetail.Id,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(key)

	fmt.Println("tokenString", tokenString, err)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExpoJWTTokenData(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			fmt.Println("facing issue")
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return &ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     big.NewInt(0),
			Y:     big.NewInt(0),
		}, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token claims")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	id, ok := claims["id"].(string)

	if !ok {
		return "", fmt.Errorf("id not found")
	}

	return id, nil
}
