package util

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTData struct {
	jwt.RegisteredClaims
	CustomClaims map[string]string `json:"custom_claims"`
}

type TokenType = string

const (
	ACCESS_TOKEN           TokenType = "access_token"
	REFRESH_TOKEN          TokenType = "refresh_token"
	ACCESS_TOKEN_DURATION            = time.Second * 3600
	REFRESH_TOKEN_DURATION           = time.Hour * 24
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

type PayloadIF interface {
	ToJson() map[string]string
}

func GenerateJWTToken(payload PayloadIF, tokenType TokenType) (string, error) {

	var expAt *jwt.NumericDate
	if tokenType == ACCESS_TOKEN {
		expAt = &jwt.NumericDate{
			Time: time.Now().Add(ACCESS_TOKEN_DURATION),
		}
	} else if tokenType == REFRESH_TOKEN {
		expAt = &jwt.NumericDate{
			Time: time.Now().Add(REFRESH_TOKEN_DURATION),
		}
	} else {
		return "", fmt.Errorf("invalid token type :%s", tokenType)
	}

	claims := JWTData{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expAt,
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			Subject: tokenType,
		},
		CustomClaims: payload.ToJson(),
	}
	fmt.Println("YEETTt")

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenString.SignedString([]byte(SECRET_KEY))
}

func validateJWTToken(token string) (map[string]string, error) {

	claims := JWTData{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	return claims.CustomClaims, nil
}
