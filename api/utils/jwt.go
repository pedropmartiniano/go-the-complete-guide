package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = `83faeac5-2d54-4115-9a67-99d8e9aa808e`

func GenerateJwtToken(email string, userId int64) (string, error) {
	// criando um novo token, definindo o método de algoritmo dele(HS256), e definindo as claims(informações do token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		// dizendo que o token vai expirar em 2 horas(pegando a data atual e adicionando 2 horas), e depois transformando a data no formato "Unix"
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	// separar o "Bearer" do token
	tokens := strings.Split(token, " ")
	// com o valor de parsedToken, já foi verificado se o algoritmo do token é o mesmo criado por mim, e com esse valor será possível verificar se o token possui a mesma "secretKey" do token criado por mim
	// verifica também se o token foi expirado
	parsedToken, err := jwt.Parse(tokens[1], func(token *jwt.Token) (interface{}, error) {
		// verificar se token.Method é do tipo passado em parenteses
		// retorna o dado convertido para esse tipo, e um bool se o valor de token.Method é do tipo passado
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		// se o tipo do token não for do tipo "jwt.SigningMethodHMAC"
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	// verificar se a assinatura do token é válida de acordo com a secretKey
	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	// verificar se as claims do token é do mesmo tipo da claims to token criado por mim, ou seja "jwt.MapClaims"
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// pegar os valores da claim armazenados dentro do array e converte-los para seus tipos
	// email, _ := claims["email"].(string)

	// todo valor númerico dentro da claims do token é convertido para float64
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
