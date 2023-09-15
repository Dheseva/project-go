package models

import (
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	IdUser int    `json:"iduser"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

