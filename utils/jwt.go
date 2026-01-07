package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTKey = []byte("uran tủn")

type Claims struct {
	UserEmail string `json:"user_email"`
	UserRole  string `json:"user_role"`
	jwt.RegisteredClaims
}

func GenerateToken(email string, userRole string) (string, error) {
	claims := Claims{
		UserEmail: email,
		UserRole:  userRole,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return JWTKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
