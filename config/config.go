package config

import "os"

var (
    JWTSecret = ""
)

func LoadConfig() {
    JWTSecret = os.Getenv("JWT_SECRET")
}
