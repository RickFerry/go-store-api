package utils

import (
    "time"
    "store-api/config"

    "github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte(config.JWTSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ValidateToken(tokenString string) bool {
    token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.JWTSecret), nil
    })

    return token.Valid
}
