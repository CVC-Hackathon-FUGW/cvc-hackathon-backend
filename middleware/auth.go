package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Username  string
	FirstName string
	LastName  string
	Uid       string
	Role      string
	jwt.StandardClaims
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	SECRET_KEY := os.Getenv("SECRET_KEY")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = "Invalid token"
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "Invalid token"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "Token expired"
		return
	}
	return claims, msg
}
