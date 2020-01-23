package helper

import (
	"log"
	"newspaper-backend/constants"

	jwt "github.com/dgrijalva/jwt-go"
)

// EncodeJWT : Generate JSON Web Token
func EncodeJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email

	tokenString, e := token.SignedString(constants.JWTKey)
	if e != nil {
		log.Println("Something Went Wrong: ", e.Error())
		return "", e
	}

	return tokenString, nil
}

// DecodeJWT : Decode JWT to Email
func DecodeJWT(token string) (string, error) {
	var decodedToken string

	claims := jwt.MapClaims{}
	_, e := jwt.ParseWithClaims(
		token,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return constants.JWTKey, nil
		})
	if e != nil {
		log.Println("Failed to Decode")
		return "", e
	}

	for _, val := range claims {
		decodedToken, _ = val.(string)
	}

	return decodedToken, nil
}
