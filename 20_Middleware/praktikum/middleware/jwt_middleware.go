package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/user/go_MoehamadRoisulJihad/20_Middleware/praktikum/constants"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}
