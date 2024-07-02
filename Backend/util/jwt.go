package util

import (
	"Backend/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey = []byte("Admiraljj")

type MyClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(user models.User) string {
	c := MyClaims{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*2, // 2 days expiration
			Issuer:    "admin",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := t.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func ParseToken(tokenString string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		user := &models.User{
			Username: claims.Username,
			Email:    claims.Email,
			Role:     claims.Role,
		}
		return user, nil
	}
	return nil, err
}