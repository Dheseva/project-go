package middleware

import "github.com/dgrijalva/jwt-go"

type MyCustomClaims struct {
	IdUser int    `json:"iduser"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

// func init() {

// 	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

// 	if err != nil{
// 		fmt.Println("Error generating ECDSA private key:", err)
// 		return
// 	}

// 	convKey = privateKey
// }

//SigningMethodHS256 for ecdsa
// 	claims := MyCustomClaims{
// 	"bar",
// 	jwt.StandardClaims{
// 		ExpiresAt: 15000,
// 		Issuer:    "test",
// 	},
// }
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// ss, err := token.SignedString(mySigningKey)