package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = `83faeac5-2d54-4115-9a67-99d8e9aa808e`

func GenerateJwtToken(email string, userId int64) (string, error) {
	// criando um novo token, definindo o método de algoritmo dele(HS256), e definindo as claims(informações do token)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		// dizendo que o token vai expirar em 2 horas(pegando a data atual e adicionando 2 horas), e depois transformando a data no formato "Unix"
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(secretKey)
}
