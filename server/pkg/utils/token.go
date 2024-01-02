package utils

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(uid uint32) (string, error) {
	// Retrieve the secret key from environment variable
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = uid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key and get the complete signed token string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractUserID(r *http.Request) (uint32, error) {
	tokenString := r.Header.Get("Authorization")
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1) // Remove Bearer prefix if present

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uidFloat64, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("uid in not valid int32")
		}

		uid := uint32(uidFloat64)
		return uid, nil
	}

	return 0, errors.New("invalid token or claims")
}
