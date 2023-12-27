// @Author huzejun 2023/12/27 21:31:00
package jwt_plugin

import "github.com/golang-jwt/jwt/v4"

var key = "abcdefg1234567"

type Data struct {
	Name   string
	Age    int
	Gender int
	jwt.RegisteredClaims
}

func Sign(data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	sign, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return sign, err
}

func Verify(sign string, data jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, data, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return err
}
